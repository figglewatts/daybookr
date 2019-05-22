package daybookr

import (
	"fmt"
	"net/url"
	"path"
	"sort"
	"strings"
	"time"

	"github.com/smallfish/simpleyaml"
)

type Site struct {
	Title          string
	Subtitle       string
	Author         string
	Pages          []Page
	Posts          []Post
	Tags           map[string][]Post
	FooterLinks    []Link
	Conf           simpleyaml.Yaml
	BaseURL        *url.URL
	GenerationTime time.Time
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

	pages, err := loadAllPages(path.Join(inputDir, pagesDir), &site)
	if err != nil {
		return Site{}, fmt.Errorf("could not load pages: %v", err)
	}
	site.Pages = pages

	site.Tags = make(map[string][]Post)

	posts, err := loadAllPosts(path.Join(inputDir, postsDir), &site)
	if err != nil {
		return Site{}, fmt.Errorf("could not load posts: %v", err)
	}

	// sort posts by date (most recent first)
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date.After(posts[j].Date)
	})
	site.Posts = posts

	site.GenerationTime = time.Now()

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

	site.Title = strings.Title(title)
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
