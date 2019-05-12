package daybookr

import (
	"fmt"
	"os"
	"strings"
	"text/template"
	"time"

	"gopkg.in/yaml.v2"
)

// PostHeader is the header of a post
type PostHeader struct {
	Tags []string `yaml:",flow"`
	Date time.Time
}

// Post stores data that gets substituted into the post
// template.
type Post struct {
	Header PostHeader
	Body   string
}

func CreatePost(templateFile string, postFile string) (string, error) {
	templateTxt, err := LoadText(templateFile)
	if err != nil {
		return "", fmt.Errorf("could not load template: %v", err)
	}

	post, err := loadPost(postFile)
	if err != nil {
		return "", fmt.Errorf("error while loading: %v", err)
	}

	tmpl, err := template.New("post").Parse(templateTxt)
	if err != nil {
		return "", err
	}

	err = tmpl.Execute(os.Stdout, post)
	if err != nil {
		return "", err
	}

	return "", nil
}

// LoadPost is used to load a markdown file to a Post struct.
func loadPost(filename string) (Post, error) {
	newPost := Post{}

	// load the text from the file
	postMarkdown, err := LoadText(filename)
	if err != nil {
		return Post{}, err
	}

	// split the header and the body
	header, body, err := getPostHeaderAndBody(postMarkdown)
	if err != nil {
		return Post{}, err
	}

	// convert the body markdown into HTML
	newPost.Body = htmlFromMarkdown(body)

	// load and validate the header
	parsedHeader, err := loadHeader(header)
	if err != nil {
		return Post{}, err
	}
	newPost.Header = parsedHeader

	return newPost, nil
}

// split a markdown post into header (YAML front matter) and body (markdown)
// returns header (as string) and body (as string), respectively
func getPostHeaderAndBody(post string) (string, string, error) {
	// we want to split by the separator character, ignore empty elements,
	// and trim the whitespace either side
	var splitNotEmpty []string
	for _, s := range strings.Split(post, "---") {
		if s != "" {
			splitNotEmpty = append(splitNotEmpty, strings.TrimSpace(s))
		}
	}

	// if it's empty, the post was empty
	if len(splitNotEmpty) == 0 {
		return "", "", fmt.Errorf("post was empty")
	}

	// if it's equal to 1, there was either no header or no body
	if len(splitNotEmpty) == 1 {
		return "", "", fmt.Errorf("post needs a header AND a body")
	}

	return splitNotEmpty[0], splitNotEmpty[1], nil
}

// load the header YAML front matter into a PostHeader struct
func loadHeader(headerYaml string) (PostHeader, error) {
	header := PostHeader{}

	// unmarshal the YAML into an instance of the struct
	yamlBytes := []byte(headerYaml)
	err := yaml.UnmarshalStrict(yamlBytes, &header)
	if err != nil {
		return header, fmt.Errorf("could not load post header YAML: %v", err)
	}

	// check to see if the unmarshaled data was valid
	valid, reason := validateHeader(header)
	if !valid {
		return PostHeader{}, fmt.Errorf("post header invalid: %s", reason)
	}

	return header, nil
}

// validate a PostHeader struct to make sure all required fields are present
func validateHeader(header PostHeader) (bool, string) {
	// there must be at least one tag
	if len(header.Tags) == 0 {
		return false, "header didn't have any tags"
	}

	// there must be a date
	if header.Date.IsZero() {
		return false, "header didn't have date"
	}

	return true, ""
}
