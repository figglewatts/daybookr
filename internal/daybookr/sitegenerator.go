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
const entriesDir = "entries"

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

	templates, err := loadAllTemplates(path.Join(inputFolder, templatesDir))
	if err != nil {
		return err
	}

	// load/scan entries for tags and years/months...

	// create pages (from templates)...
	// write out to files (index in folder)

	// create entries (from template)
	// write out to files (for each entry to permalink)

	// organise into archive pages and write these out

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
