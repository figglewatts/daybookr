package daybookr

import (
	"fmt"
	"os"
	"path"
)

const templatesDir = "templates"
const postsDir = "posts"
const includesDir = "includes"
const pagesDir = "pages"

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

	site, err := createSite(baseURL, config, inputFolder)
	if err != nil {
		return fmt.Errorf("could not create site: %v", err)
	}

	fmt.Println(site)

	// get includes filenames
	includes, err := getFilesInDir(path.Join(inputFolder, includesDir), "*.html")
	if err != nil {
		return err
	}

	index, err := loadTemplate(path.Join(inputFolder, "index.html"), includes)
	if err != nil {
		return err
	}

	result, err := renderTemplate(index, site)
	if err != nil {
		return err
	}

	fmt.Println(result)

	// copy styles folder to output

	// load all templates
	/*templates, err := loadAllTemplates(path.Join(inputFolder, templatesDir), includes)
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
	fmt.Println(tags)

	// get all years from posts
	years := getAllYearsFromPosts(posts)
	fmt.Println(years)

	rendered, err := renderTemplate(templates[2], site)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	fmt.Println(rendered)*/

	// generate templates...

	// Site contains EVERYTHING to generate the site
	// templates are responsible for accessing this in the right way

	// create pages (from templates by executing)...
	// write out to files (index in folder)

	// create entries (from template by executing)
	// write out to files (for each entry to permalink)

	// organise into archive pages and write these out (from templates by executing)

	// create the index page

	return nil
}
