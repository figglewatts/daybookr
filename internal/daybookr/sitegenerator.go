package daybookr

import (
	"fmt"
	"os"
)

func Generate(baseURL string, inputFolder string, outputFolder string, configPath string) error {
	// check to see if the input folder exists
	inputFolderExists, err := exists(inputFolder)
	if err != nil || !inputFolderExists {
		return fmt.Errorf("input folder '%s' did not exist", inputFolder)
	}

	// check to see if the config file exists
	configPathExists, err := exists(configPath)
	if err != nil || !configPathExists {
		return fmt.Errorf("config file '%s' did not exist", configPath)
	}

	// check to see if the output folder exists, and create it if not
	outputFolderExists, err := exists(outputFolder)
	if err != nil {
		return err
	}
	if !outputFolderExists {
		os.MkdirAll(outputFolder, os.ModePerm)
	}

	// load the config
	config, err := loadConfig(configPath)
	if err != nil {
		return fmt.Errorf("could not load config: %v", err)
	}

	site, err := createSite(baseURL, config)
	if err != nil {
		return fmt.Errorf("could not create site: %v", err)
	}

	fmt.Println(site)

	return nil
}
