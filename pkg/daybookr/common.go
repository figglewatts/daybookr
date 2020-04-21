package daybookr

import (
	"bytes"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"path/filepath"

	"github.com/yuin/goldmark"
)

// the goldmark markdown renderer with additional options and extensions
var mdRenderer = goldmark.New(
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
		goldmark.WithExtensions(
			extension.GFM,
			extension.Footnote,
		),
	)

// convert markdown into HTML
func htmlFromMarkdown(markdown string) string {
	markdownBytes := []byte(markdown)
	var buf bytes.Buffer
	if err := mdRenderer.Convert(markdownBytes, &buf); err != nil {
		panic(err)
	}
	return buf.String()
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

func getFilesInDir(dirPath string, pattern string) ([]string, error) {
	var files []string
	dir, err := os.Open(dirPath)
	if err != nil {
		return files, err
	}
	fileInfos, err := dir.Readdir(-1)
	dir.Close()
	if err != nil {
		return files, err
	}

	for _, file := range fileInfos {
		fileNamePath := path.Join(dirPath, file.Name())
		fileNameMatches, _ := filepath.Match(pattern, file.Name())
		if fileNameMatches {
			files = append(files, fileNamePath)
		}
	}
	return files, nil
}

func LoadText(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
