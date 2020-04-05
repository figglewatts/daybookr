package daybookr

import (
	"sort"
)

func To(end int, s []string) []string {
	return s[:end]
}

func From(start int, s []string) []string {
	return s[start:]
}

func FromTo(start int, end int, s []string) []string {
	return s[start:end]
}

type yearPosts struct {
	Year  int
	Posts []Post
}

func PostsByYear(site Site) []yearPosts {
	postsByYearUnsorted := make(map[int][]Post)
	for _, post := range site.Posts {
		postsByYearUnsorted[post.Date.Year()] = append(postsByYearUnsorted[post.Date.Year()], post)
	}

	// sort the years in descending order
	var years []int
	for year := range postsByYearUnsorted {
		years = append(years, year)
	}
	sort.Slice(years, func(i, j int) bool {
		return years[i] > years[j]
	})

	var postsByYearSorted []yearPosts
	for _, year := range years {
		postsByYearSorted = append(postsByYearSorted, yearPosts{
			year,
			postsByYearUnsorted[year],
		})
	}

	return postsByYearSorted
}