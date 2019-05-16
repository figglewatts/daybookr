package daybookr

import (
	"fmt"
	"time"
)

const (
	tagsFieldName = "tags"
	dateFieldName = "date"
	iso8601Date   = "2006-01-02"
)

type Post struct {
	Page
	Tags []string
	Date time.Time
}

func loadAllPosts(postsDir string) ([]Post, error) {
	var loadedPosts []Post
	posts, err := getFilesInDir(postsDir, "*.md")
	if err != nil {
		return nil, err
	}
	for _, post := range posts {
		loadedPost, err := loadPost(post)
		if err != nil {
			return nil, fmt.Errorf("could not load post '%s': %v", post, err)
		}
		loadedPosts = append(loadedPosts, loadedPost)
	}
	return loadedPosts, nil
}

func loadPost(postPath string) (Post, error) {
	post := Post{}
	page, err := loadPage(postPath)
	post.Page = page

	err = post.validatePostMetadata()
	if err != nil {
		return Post{}, err
	}

	return post, nil
}

func (post *Post) validatePostMetadata() error {
	metadata := post.Metadata

	// metadata must be a map
	metadataMap, err := metadata.Map()
	if err != nil {
		return fmt.Errorf("malformed header: %v", err)
	}

	// the metadata must have a tags field
	if _, ok := metadataMap[tagsFieldName]; !ok {
		return fmt.Errorf("post metadata needs %s value", tagsFieldName)
	}
	// try and get the tags field
	tagsYaml := metadata.Get(tagsFieldName)
	tagsArr, err := tagsYaml.Array()
	if err != nil {
		return fmt.Errorf("post %s must be array", tagsFieldName)
	}
	tags := make([]string, len(tagsArr))
	for i := range tagsArr {
		tag, err := tagsYaml.GetIndex(i).String()
		if err != nil {
			return fmt.Errorf("post tag %d was not a string", i)
		}
		tags[i] = tag
	}

	// the metadata must have a date field
	if _, ok := metadataMap[dateFieldName]; !ok {
		return fmt.Errorf("post metadata needs %s value", dateFieldName)
	}
	// try and get the date field
	dateStr, err := metadata.Get(dateFieldName).String()
	if err != nil {
		return fmt.Errorf("post %s must be string", dateFieldName)
	}
	date, err := time.Parse(iso8601Date, dateStr)
	if err != nil {
		return fmt.Errorf("malformed date: %v", err)
	}

	post.Tags = tags
	post.Date = date
	return nil
}

func getAllTagsFromPosts(posts []Post) map[string][]Post {
	tags := make(map[string][]Post)
	for _, post := range posts {
		for _, tag := range post.Tags {
			tags[tag] = append(tags[tag], post)
		}
	}
	return tags
}

func getAllYearsFromPosts(posts []Post) map[int][]Post {
	years := make(map[int][]Post)
	for _, post := range posts {
		years[post.Date.Year()] = append(years[post.Date.Year()], post)
	}
	return years
}
