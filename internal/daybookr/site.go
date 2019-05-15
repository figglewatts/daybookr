package daybookr

import (
	"fmt"
	"net/url"
	"path"

	"github.com/smallfish/simpleyaml"
)

type Site struct {
	Title       string
	Subtitle    string
	Author      string
	Pages       []Page
	HeaderLinks []Link
	FooterLinks []Link
	Conf        simpleyaml.Yaml
	Tags        []string
	BaseURL     *url.URL
}

func (site Site) MakeSiteURL(relativeURL string) (*url.URL, error) {
	urlString := path.Join(site.BaseURL.Path, relativeURL)
	url, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}
	return url, nil
}

func createSite(baseURL string, config *simpleyaml.Yaml) (Site, error) {
	site := Site{}
	err := site.populateWithConfig(config)
	if err != nil {
		return Site{}, fmt.Errorf("unable to create site from config: %v", err)
	}
	createdBaseURL, err := makeURL(baseURL)
	if err != nil {
		return Site{}, fmt.Errorf("invalid base URL: %v", err)
	}
	site.BaseURL = createdBaseURL
	return site, nil
}

func (site *Site) populateWithConfig(config *simpleyaml.Yaml) error {
	title, err := config.Get(ConfigTitleField).String()
	if err != nil {
		return err
	}
	subtitle, err := config.Get(ConfigSubtitleField).String()
	if err != nil {
		return err
	}
	author, err := config.Get(ConfigAuthorField).String()
	if err != nil {
		return err
	}

	pagesArrayYAML := config.Get(ConfigPagesField)
	footerLinksArrayYAML := config.Get(ConfigFooterLinksField)

	site.Title = title
	site.Subtitle = subtitle
	site.Author = author

	pagesArray, err := pagesArrayYAML.Array()
	if err != nil {
		return err
	}
	site.Pages = make([]Page, len(pagesArray))
	for i := range pagesArray {
		page, err := CreatePageFromYAML(pagesArrayYAML.GetIndex(i))
		if err != nil {
			return fmt.Errorf("could not create page: %v", err)
		}
		site.Pages[i] = page
	}

	footerLinksArray, err := footerLinksArrayYAML.Array()
	if err != nil {
		return err
	}
	site.FooterLinks = make([]Link, len(footerLinksArray))
	for i := range footerLinksArray {
		link, err := CreateLinkFromYAML(footerLinksArrayYAML.GetIndex(i))
		if err != nil {
			return fmt.Errorf("could not create link: %v", err)
		}
		site.FooterLinks[i] = link
	}

	return nil
}
