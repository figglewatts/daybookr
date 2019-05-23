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
	Tags           []tag
	FooterLinks    []Link
	Conf           simpleyaml.Yaml
	BaseURL        *url.URL
	GenerationTime time.Time
}

type tag struct {
	Name  string
	Posts []Post
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

	posts, err := loadAllPosts(path.Join(inputDir, postsDir), &site)
	if err != nil {
		return Site{}, fmt.Errorf("could not load posts: %v", err)
	}

	// sort posts by date (most recent first)
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date.After(posts[j].Date)
	})
	site.Posts = posts

	// link tags to posts
	site.Tags = compileTags(site)

	site.GenerationTime = time.Now()

	return site, nil
}

func compileTags(site Site) []tag {
	// first link all unique tag names to their tagged posts
	tagsUnsorted := make(map[string][]Post)
	for _, post := range site.Posts {
		for _, tag := range post.Tags {
			tagsUnsorted[tag] = append(tagsUnsorted[tag], post)
		}
	}

	// now sort tag names alphabetically
	var tagNames []string
	for tagName := range tagsUnsorted {
		tagNames = append(tagNames, tagName)
	}
	sort.Slice(tagNames, func(i, j int) bool {
		return strings.ToLower(tagNames[i]) < strings.ToLower(tagNames[j])
	})

	// finally, make an alphabetical array of tags with their posts
	var tags []tag
	for _, tagName := range tagNames {
		tags = append(tags, tag{
			tagName,
			tagsUnsorted[tagName],
		})
	}
	return tags
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
