package daybookr

import (
	"io/ioutil"
	"net/url"
	"os"

	"gopkg.in/russross/blackfriday.v2"
)

// convert markdown into HTML
func htmlFromMarkdown(markdown string) string {
	markdownBytes := []byte(markdown)
	htmlBytes := blackfriday.Run(markdownBytes)
	return string(htmlBytes)
}

func makeURL(urlString string) (*url.URL, error) {
	return url.Parse(urlString)
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func LoadText(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
