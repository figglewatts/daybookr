package daybookr

import (
	"fmt"

	"github.com/smallfish/simpleyaml"
)

const (
	ConfigTitleField       = "title"
	ConfigSubtitleField    = "subtitle"
	ConfigAuthorField      = "author"
	ConfigPagesField       = "pages"
	ConfigFooterLinksField = "footer-links"
)

var requiredFields = [...]string{
	ConfigTitleField,
	ConfigSubtitleField,
	ConfigAuthorField,
	ConfigPagesField,
	ConfigFooterLinksField,
}

// loadConfig loads a config YAML file into a Yaml structure (from simpleyaml)
func loadConfig(configPath string) (*simpleyaml.Yaml, error) {
	// load the text of the config file
	confYaml, err := LoadText(configPath)
	if err != nil {
		return nil, err
	}

	// now attempt to parse to yaml
	conf, err := simpleyaml.NewYaml([]byte(confYaml))
	if err != nil {
		return nil, err
	}

	// now check to see if it's valid
	err = validConfig(conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

func validConfig(config *simpleyaml.Yaml) error {
	// check through each of the required fields to see if they're present
	for _, field := range requiredFields {
		configMap, _ := config.Map()
		if _, ok := configMap[field]; !ok {
			return fmt.Errorf("field '%s' not present in config", field)
		}
	}
	return nil
}
