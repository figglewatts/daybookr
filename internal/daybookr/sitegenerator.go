package daybookr

import (
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

func LoadText(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func CreatePost(templateFile string, postFile string) (string, error) {
	templateTxt, err := LoadText(templateFile)
	if err != nil {
		return "", fmt.Errorf("could not load template: %v", err)
	}

	post, err := LoadPost(postFile)
	if err != nil {
		return "", fmt.Errorf("error while loading: %v", err)
	}

	tmpl, err := template.New("post").Parse(templateTxt)
	if err != nil {
		return "", err
	}

	err = tmpl.Execute(os.Stdout, post)
	if err != nil {
		return "", err
	}

	return "", nil
}
