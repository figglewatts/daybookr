package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"time"

	_ "github.com/figglewatts/daybookr/cmd/daybookr/statik"
	"github.com/figglewatts/daybookr/pkg/daybookr"
	"github.com/rakyll/statik/fs"
	"github.com/urfave/cli/v2"
	"github.com/pkg/browser"
)

//go:generate statik -f -src=../../test_dir

func generateSite(c *cli.Context) error {
	inputFolder := c.Args().Get(0)
	if len(inputFolder) == 0 {
		inputFolder = "."
	}

	outputFolder := c.Args().Get(1)
	if len(outputFolder) == 0 {
		outputFolder = "static"
	}
	outputFolder = path.Join(inputFolder, outputFolder)

	configPath := c.Args().Get(2)
	if len(configPath) == 0 {
		configPath = "daybook.yml"
	}
	configPath = path.Join(inputFolder, configPath)

	// generate the site
	fmt.Printf("Generating site in '%s'\n", outputFolder)
	err := daybookr.Generate(inputFolder, outputFolder, configPath)
	if err != nil {
		return err
	}

	fmt.Printf("Successfully generated site in %s\n", outputFolder)
	return nil
}

func newProject(c *cli.Context) error {
	outputFolder := c.Args().Get(0)

	outputPathExists, err := exists(outputFolder)
	if err != nil {
		return err
	}
	if outputPathExists {
		outputPathIsEmpty, err := isEmpty(outputFolder)
		if !outputPathIsEmpty || err != nil {
			return fmt.Errorf("can't create a new project in existing non-empty directory '%v'", outputFolder)
		}
	}

	err = os.MkdirAll(outputFolder, 0644)
	if err != nil {
		return err
	}

	statikFs, err := fs.New()
	if err != nil {
		return err
	}

	err = fs.Walk(statikFs, "/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Printf("%v\n", path)
		if !info.IsDir() {
			// it's a file, write it to the output dir
			r, err := statikFs.Open(path)
			if err != nil {
				return err
			}
			fileData, err := ioutil.ReadAll(r)
			if err != nil {
				return err
			}
			targetDirectory := filepath.Dir(filepath.Join(outputFolder, path))
			err = os.MkdirAll(targetDirectory, 0644)
			if err != nil {
				return err
			}
			err = ioutil.WriteFile(filepath.Join(outputFolder, path), fileData, 0644)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	fmt.Printf("Created new daybookr project in '%s'\n", outputFolder)
	return nil
}

func openProject(c *cli.Context) error {
	inputFolder := c.Args().Get(0)
	if len(inputFolder) == 0 {
		inputFolder = "."
	}

	outputFolder := c.Args().Get(1)
	if len(outputFolder) == 0 {
		outputFolder = "static"
	}
	indexPage := path.Join(inputFolder, outputFolder, "index.html")

	err := browser.OpenFile(indexPage)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	app := &cli.App{
		Name: "daybookr",
		Usage: "generate a static site based on some content, config and templates",
		UsageText: "daybookr [global options] COMMAND [options]",
		Version: "2.0.3",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name: "Figglewatts",
			},
		},
		Copyright: fmt.Sprintf("(c) %d Figglewatts", time.Now().Year()),
		Commands: []*cli.Command{
			{
				Name: "generate",
				Usage: "generate the site",
				UsageText: "daybookr generate [INPUT_DIR] [OUTPUT_DIR] [CONFIG]",
				Description: `INPUT_DIR:  the directory containing data files of the site to generate. Default: '.'.
   OUTPUT_DIR: the directory (relative to INPUT_DIR) to generate the site into. Default: 'static'.
   CONFIG:     the path (relative to INPUT_DIR) to the config file to use to generate the site. Default: 'daybook.yml'.`,
				Action: generateSite,
			},
			{
				Name: "new",
				Usage: "create a new site",
				UsageText: "daybookr new DIR",
				Description: "DIR: the directory to generate the new site in.",
				Action: newProject,
			},
			{
				Name: "open",
				Usage: "open a daybookr site in a browser",
				UsageText: "daybookr open [PROJECT_DIR] [OUTPUT_DIR]",
				Description: `PROJECT_DIR: the directory containing data files of the site to generate. Default: '.'.
   OUTPUT_DIR:  the directory (relative to PROJECT_DIR) containing the generated site. Default: 'static'.`,
				Action: openProject,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}
}
