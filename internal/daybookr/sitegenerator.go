package daybookr

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

const templatesDir = "templates"
const postsDir = "posts"

func Generate(baseURL string, inputFolder string, outputFolder string, configPath string) error {
	// check to see if the input folder exists
	inputFolderExists, err := exists(inputFolder)
	if err != nil || !inputFolderExists {
		return fmt.Errorf("input folder '%s' did not exist", inputFolder)
	}

	// check to see if the config file exists
	configPathExists, err := exists(configPath)
	if err != nil || !configPathExists {
		return fmt.Errorf("config file '%s' did not exist", configPath)
	}

	// check to see if the output folder exists, and create it if not
	outputFolderExists, err := exists(outputFolder)
	if err != nil {
		return err
	}
	if !outputFolderExists {
		os.MkdirAll(outputFolder, os.ModePerm)
	}

	// load the config
	config, err := loadConfig(configPath)
	if err != nil {
		return fmt.Errorf("could not load config: %v", err)
	}

	site, err := createSite(baseURL, config)
	if err != nil {
		return fmt.Errorf("could not create site: %v", err)
	}

	fmt.Println(site)

	// load all templates
	templates, err := loadAllTemplates(path.Join(inputFolder, templatesDir))
	if err != nil {
		return err
	}

	fmt.Println(templates)

	// load all posts
	posts, err := loadAllPosts(path.Join(inputFolder, postsDir))
	if err != nil {
		return err
	}

	// get all the tags from posts
	tags := getAllTagsFromPosts(posts)
	for _, tag := range tags.Iterate() {
		fmt.Println(tag)
	}

	// Site contains EVERYTHING to generate the site
	// templates are responsible for accessing this in the right way

	// load/scan entries for tags and years/months...

	// create pages (from templates by executing)...
	// write out to files (index in folder)

	// create entries (from template by executing)
	// write out to files (for each entry to permalink)

	// organise into archive pages and write these out (from templates by executing)

	// create the index page

	return nil
}

func loadAllTemplates(templatesDir string) ([]*template.Template, error) {
	var loadedTemplates []*template.Template
	templates, err := getFilesInDir(templatesDir, "*.html")
	if err != nil {
		return nil, err
	}
	for _, template := range templates {
		loadedTemplate, err := loadTemplate(template)
		if err != nil {
			return nil, fmt.Errorf("could not load template '%s': %v", template, err)
		}
		loadedTemplates = append(loadedTemplates, loadedTemplate)
	}
	return loadedTemplates, nil
}

func loadTemplate(templatePath string) (*template.Template, error) {
	// the file name is the template name, grab it without the extension
	templateName := strings.Split(filepath.Base(templatePath), ".")[0]

	templateContent, err := LoadText(templatePath)
	if err != nil {
		return nil, err
	}

	return template.New(templateName).Parse(templateContent)
}

func loadAllPosts(postDir string) ([]Post, error) {
	var loadedPosts []Post
	posts, err := getFilesInDir(postDir, "*.md")
	if err != nil {
		return nil, err
	}
	for _, post := range posts {
		loadedPost, err := loadPost(post)
		if err != nil {
			return nil, fmt.Errorf("could not load post '%s': %v", post, err)
		}
		loadedPosts = append(loadedPosts, loadedPost)
	}
	return loadedPosts, nil
}

func getAllTagsFromPosts(posts []Post) *TagSet {
	tags := NewTagSet()
	for _, post := range posts {
		for _, tag := range post.Header.Tags {
			tags.Add(tag)
		}
	}
	return tags
}
