package daybookr

import (
	"fmt"
	"time"
)

const (
	tagsFieldName = "tags"
	dateFieldName = "date"

	// go has weird ways of specifying date formatting
	// this one adheres to ISO-8601
	iso8601Date = "2006-01-02"

	// this one makes it nicely human readable
	humanDate = "Monday, 2 January, 2006"
)

type Post struct {
	Page
	Tags []string
	Date time.Time
}

func (post Post) combine(other Post) Post {
	post.Content += "\n" + other.Content
	post.Tags = append(post.Tags, other.Tags...)
	return post
}

func loadAllPosts(postsDir string, site *Site) ([]Post, error) {
	existingPosts := make(map[time.Time]Post)

	var loadedPosts []Post
	posts, err := getFilesInDir(postsDir, "*.md")
	if err != nil {
		return nil, err
	}
	for _, post := range posts {
		loadedPost, err := loadPost(post, site)
		if err != nil {
			return nil, fmt.Errorf("could not load post '%s': %v", post, err)
		}

		// deduplicate posts with the same date
		if val, ok := existingPosts[loadedPost.Date]; ok {
			loadedPost = val.combine(loadedPost)
		}
		existingPosts[loadedPost.Date] = loadedPost
	}

	for _, v := range existingPosts {
		loadedPosts = append(loadedPosts, v)
	}

	return loadedPosts, nil
}

func loadPost(postPath string, site *Site) (Post, error) {
	post := Post{}
	page, err := loadPage(postPath, site)
	post.Page = page

	tags, postDate, err := getPostMetadata(post)
	post.Tags = tags
	post.Date = postDate
	if err != nil {
		return Post{}, err
	}

	// override the title with the date as a human-readable string
	post.Title = post.Date.Format(humanDate)

	// override the name with the date as a string
	post.Name = post.Date.Format(iso8601Date)

	return post, nil
}

func getPostMetadata(post Post) ([]string, time.Time, error) {
	metadata := post.Metadata

	// metadata must be a map
	metadataMap, err := metadata.Map()
	if err != nil {
		return nil, time.Time{}, fmt.Errorf("malformed header: %v", err)
	}

	// the metadata must have a tags field
	if _, ok := metadataMap[tagsFieldName]; !ok {
		return nil, time.Time{}, fmt.Errorf("post metadata needs %s value", tagsFieldName)
	}
	// try and get the tags field
	tagsYaml := metadata.Get(tagsFieldName)
	tagsArr, err := tagsYaml.Array()
	if err != nil {
		return nil, time.Time{}, fmt.Errorf("post %s must be array", tagsFieldName)
	}
	tags := make([]string, len(tagsArr))
	for i := range tagsArr {
		tag, err := tagsYaml.GetIndex(i).String()
		if err != nil {
			return nil, time.Time{}, fmt.Errorf("post tag %d was not a string", i)
		}
		tags[i] = tag
	}

	// the metadata must have a date field
	if _, ok := metadataMap[dateFieldName]; !ok {
		return nil, time.Time{}, fmt.Errorf("post metadata needs %s value", dateFieldName)
	}
	// try and get the date field
	dateStr, err := metadata.Get(dateFieldName).String()
	if err != nil {
		return nil, time.Time{}, fmt.Errorf("post %s must be string", dateFieldName)
	}
	date, err := time.Parse(iso8601Date, dateStr)
	if err != nil {
		return nil, time.Time{}, fmt.Errorf("malformed date: %v", err)
	}

	return tags, date, nil
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
