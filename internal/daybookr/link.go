package daybookr

import (
	"net/url"

	"github.com/smallfish/simpleyaml"
)

type Link struct {
	Text   string
	Target *url.URL
}

const (
	LinkYAMLTextField   = "text"
	LinkYAMLTargetField = "target"
)

func CreateLink(text string, target string) (Link, error) {
	link := Link{Text: text}
	url, err := url.Parse(target)
	if err != nil {
		return Link{}, err
	}
	link.Target = url
	return link, nil
}

func CreateLinkFromYAML(yaml *simpleyaml.Yaml) (Link, error) {
	text, err := yaml.Get(LinkYAMLTextField).String()
	if err != nil {
		return Link{}, err
	}

	target, err := yaml.Get(LinkYAMLTargetField).String()
	if err != nil {
		return Link{}, err
	}
	return CreateLink(text, target)
}
