package daybookr

import (
	"fmt"
	"net/url"
	"path"

	"github.com/smallfish/simpleyaml"
)

// TODO(sam): fields for archives
type Site struct {
	Title       string
	Subtitle    string
	Author      string
	Pages       []Page
	Posts       []Post
	HeaderLinks []Link
	FooterLinks []Link
	Conf        simpleyaml.Yaml
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

func createSite(baseURL string, config *simpleyaml.Yaml, inputDir string) (Site, error) {
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

	pages, err := loadAllPages(path.Join(inputDir, pagesDir))
	if err != nil {
		return Site{}, fmt.Errorf("could not load pages: %v", err)
	}
	site.Pages = pages

	posts, err := loadAllPosts(path.Join(inputDir, postsDir))
	if err != nil {
		return Site{}, fmt.Errorf("could not load posts: %v", err)
	}
	site.Posts = posts

	return site, nil
}

func (site *Site) populateWithConfig(config *simpleyaml.Yaml) error {
	title, err := config.Get(configTitleField).String()
	if err != nil {
		return err
	}
	subtitle, err := config.Get(configSubtitleField).String()
	if err != nil {
		return err
	}
	author, err := config.Get(configAuthorField).String()
	if err != nil {
		return err
	}

	footerLinksArrayYAML := config.Get(configFooterLinksField)

	site.Title = title
	site.Subtitle = subtitle
	site.Author = author

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
