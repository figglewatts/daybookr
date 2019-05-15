package daybookr

import (
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"path/filepath"

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
