package daybookr

import (
	"bufio"
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/smallfish/simpleyaml"
)

const layoutFieldName = "layout"

type Page struct {
	Layout     string
	Metadata   *simpleyaml.Yaml
	RawContent string
	Content    string
	Name       string
	Title      string
	Site       *Site
}

func (page Page) GetHTMLPreview(chars int) string {
	// find first instance of text that isn't a heading
	// append until chars or next heading
	textAlphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var builder strings.Builder
	stringScanner := bufio.NewScanner(strings.NewReader(page.RawContent))
	foundText := false
	for stringScanner.Scan() {
		// if we're over the number of chars, truncate with ellipsis
		if builder.Len() >= chars {
			builder.WriteString("...")
			return htmlFromMarkdown(strings.TrimSpace(builder.String()))
		}

		line := stringScanner.Text()
		if len(line) == 0 {
			builder.WriteString("\n\n")
			continue
		}

		// if the char is a heading and we haven't found text yet, skip it
		if !foundText && line[0] == '#' {
			continue
		} else if !foundText && strings.ContainsRune(textAlphabet, rune(line[0])) {
			// otherwise if the char is an alpha then we have found text
			foundText = true
		} else if foundText && line[0] == '#' {
			break
		}

		for _, c := range line {
			if builder.Len()+1 > chars {
				builder.WriteString("...")
				return htmlFromMarkdown(strings.TrimSpace(builder.String()))
			}
			builder.WriteRune(c)
		}
		builder.WriteRune('\n')
	}
	return htmlFromMarkdown(strings.TrimSpace(builder.String()))
}

func loadAllPages(pagesDir string, site *Site) ([]Page, error) {
	var loadedPages []Page
	pages, err := getFilesInDir(pagesDir, "*.md")
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		loadedPage, err := loadPage(page, site)
		if err != nil {
			return nil, fmt.Errorf("could not load page '%s': %v", page, err)
		}
		loadedPages = append(loadedPages, loadedPage)
	}
	return loadedPages, nil
}

func loadPage(pagePath string, site *Site) (Page, error) {
	// load the page's text
	pageText, err := LoadText(pagePath)
	if err != nil {
		return Page{}, err
	}

	// split into header and body
	header, body, err := getPageHeaderAndBody(pageText)
	if err != nil {
		return Page{}, err
	}

	// load the header
	pageLayout, metadata, err := loadPageHeader(header)
	if err != nil {
		return Page{}, err
	}

	// convert the page body into HTML
	pageBody := htmlFromMarkdown(body)

	// the page name is the filename without the extension
	pageName := path.Base(pagePath)
	pageName = strings.TrimSuffix(pageName, filepath.Ext(pageName))

	pageTitle := strings.Title(pageName)

	return Page{
		Layout:     pageLayout,
		Metadata:   metadata,
		RawContent: body,
		Content:    pageBody,
		Name:       pageName,
		Site:       site,
		Title:      pageTitle,
	}, nil
}

// split a markdown page into header (YAML front matter) and body (markdown)
// returns header (as string) and body (as string), respecti vely
func getPageHeaderAndBody(page string) (string, string, error) {
	// we want to split by the separator character, ignore empty elements,
	// and trim the whitespace either side
	var splitNotEmpty []string
	for _, s := range strings.SplitN(page, "---", 3) {
		if s != "" {
			splitNotEmpty = append(splitNotEmpty, strings.TrimSpace(s))
		}
	}

	// if it's empty, the post was empty
	if len(splitNotEmpty) == 0 {
		return "", "", fmt.Errorf("page was empty")
	}

	// if it's equal to 1, there was either no header or no body
	if len(splitNotEmpty) == 1 {
		return "", "", fmt.Errorf("page needs a header AND a body")
	}

	return splitNotEmpty[0], splitNotEmpty[1], nil
}

func loadPageHeader(pageHeader string) (string, *simpleyaml.Yaml, error) {
	headerBytes := []byte(pageHeader)
	yaml, err := simpleyaml.NewYaml(headerBytes)
	if err != nil {
		return "", nil, err
	}

	// see if the YAML can be converted into a map or not... we need it to be
	yamlAsMap, err := yaml.Map()
	if err != nil {
		return "", nil, fmt.Errorf("malformed header: %v", err)
	}

	// the YAML must have a layout field
	if _, ok := yamlAsMap[layoutFieldName]; !ok {
		return "", nil, fmt.Errorf("header needs %s value", layoutFieldName)
	}

	// try and get the layout field
	layout, err := yaml.Get(layoutFieldName).String()
	if err != nil {
		return "", nil, fmt.Errorf("header %s field must be string: %v", layoutFieldName, err)
	}

	return layout, yaml, nil
}
