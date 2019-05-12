package daybookr

import "github.com/smallfish/simpleyaml"

type Page struct {
	Template string
	Content  string
}

const (
	PageYAMLTemplateField = "template"
	PageYAMLContentField  = "content"
)

func CreatePageFromYAML(yaml *simpleyaml.Yaml) (Page, error) {
	template, err := yaml.Get(PageYAMLTemplateField).String()
	if err != nil {
		return Page{}, err
	}

	content, err := yaml.Get(PageYAMLContentField).String()
	if err != nil {
		return Page{}, err
	}

	return Page{Template: template, Content: content}, nil
}
