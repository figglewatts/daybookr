package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")
	markdown := []byte(`
# Markdown test
https://www.google.com
	`)
	output := daybookr.htmlFromMarkdown(markdown)
	fmt.Println(string(output))
}
