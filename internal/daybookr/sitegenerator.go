package daybookr

import "gopkg.in/russross/blackfriday.v2"

// Convert a markdown string to HTML
func htmlFromMarkdown(markdown string) string {
	markdownBytes := []byte(markdown)
	htmlBytes := blackfriday.Run(markdownBytes)
	return string(htmlBytes)
}
