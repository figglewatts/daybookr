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
	for _, siteTemplate := range templates {
		loadedTemplate, err := loadTemplate(siteTemplate, includes)
		if err != nil {
			fmt.Println("err")
			return nil, fmt.Errorf("could not load siteTemplate '%s': %v", siteTemplate, err)
		}
		// siteTemplate name is siteTemplate filename without extension
		templateName := path.Base(siteTemplate)
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

	templateLoaded, err := template.New(templateName).Funcs(createFuncMap()).Parse(templateContent)
	if err != nil {
		return nil, err
	}

	templateLoaded, err = loadTemplateString(templateName, templateContent, includes)
	if err != nil {
		return nil, err
	}

	return templateLoaded, nil
}

func loadTemplateString(templateName string, templateStr string, includes []string) (*template.Template, error) {
	templateLoaded, err := template.New(templateName).Funcs(createFuncMap()).Parse(templateStr)
	if err != nil {
		return nil, err
	}

	// now load the includes into this template if they're given
	if includes != nil {
		templateLoaded, err = templateLoaded.ParseFiles(includes...)
		if err != nil {
			return nil, fmt.Errorf("unable to load template includes: %v", err)
		}
	}

	return templateLoaded, nil
}

func createFuncMap() template.FuncMap {
	return template.FuncMap{
		"Title":       strings.Title,
		"From":        From,
		"To":          To,
		"FromTo":      FromTo,
		"PostsByYear": PostsByYear,
	}
}
