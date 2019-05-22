package daybookr

func To(end int, s []string) []string {
	return s[:end]
}

func From(start int, s []string) []string {
	return s[start:]
}

func FromTo(start int, end int, s []string) []string {
	return s[start:end]
}

func PostsByYear(site Site) map[int][]Post {
	output := make(map[int][]Post)
	for _, post := range site.Posts {
		output[post.Date.Year()] = append(output[post.Date.Year()], post)
	}
	return output
}
