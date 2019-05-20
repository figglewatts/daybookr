package daybookr

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

func renderTemplate(t *template.Template, data interface{}) (string, error) {
	var builder strings.Builder
	err := t.Execute(&builder, data)
	return builder.String(), err
}

func loadAllTemplates(templatesDir string, includes []string) (map[string]*template.Template, error) {
	loadedTemplates := make(map[string]*template.Template)
	templates, err := getFilesInDir(templatesDir, "*.html")
	if err != nil {
		return nil, err
	}
	for _, template := range templates {
		loadedTemplate, err := loadTemplate(template, includes)
		if err != nil {
			return nil, fmt.Errorf("could not load template '%s': %v", template, err)
		}
		// template name is template filename without extension
		templateName := path.Base(template)
		templateName = strings.TrimSuffix(templateName, filepath.Ext(templateName))
		loadedTemplates[templateName] = loadedTemplate
	}
	return loadedTemplates, nil
}

func loadTemplate(templatePath string, includes []string) (*template.Template, error) {
	// the file name is the template name, grab it without the extension
	templateName := strings.Split(filepath.Base(templatePath), ".")[0]

	templateContent, err := LoadText(templatePath)
	if err != nil {
		return nil, err
	}

	template, err := template.New(templateName).Parse(templateContent)
	if err != nil {
		return nil, err
	}

	// now load the includes into this template
	template, err = template.ParseFiles(includes...)
	if err != nil {
		return nil, fmt.Errorf("unable to load template includes: %v", err)
	}

	return template, err
}
