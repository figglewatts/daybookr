package main

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/figglewatts/daybookr/internal/daybookr"
	"gopkg.in/urfave/cli.v1"
)

type daybookrFlags struct {
	InputFolder  string
	OutputFolder string
	ConfigPath   string
}

const (
	inputFolderFlagName   = "input"
	inputFolderShortName  = "i"
	outputFolderFlagName  = "output"
	outputFolderShortName = "o"
	configPathFlagName    = "config"
	configPathShortName   = "c"
)

func getFlags(c *cli.Context) daybookrFlags {
	inputFolder := c.String(inputFolderShortName)
	return daybookrFlags{
		InputFolder:  inputFolder,
		OutputFolder: path.Join(inputFolder, c.String(outputFolderFlagName)),
		ConfigPath:   path.Join(inputFolder, c.String(configPathFlagName)),
	}
}

func runDaybookr(c *cli.Context) error {
	// get optional flags
	flags := getFlags(c)

	// generate the site
	err := daybookr.Generate(flags.InputFolder, flags.OutputFolder, flags.ConfigPath)
	if err != nil {
		return err
	}

	fmt.Printf("Successfully generated site in %s\n", flags.OutputFolder)

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "daybookr"
	app.Usage = "generate a journal based on some content, config and templates"
	app.UsageText = "daybookr [global options]"
	app.Action = runDaybookr
	app.Author = "Figglewatts"
	app.Version = "1.0"
	app.Compiled = time.Now()
	app.Copyright = "(c) 2019 Figglewatts"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  fmt.Sprintf("%s, %s", inputFolderFlagName, inputFolderShortName),
			Usage: "use `FOLDER` to generate the site",
			Value: ".",
		},
		cli.StringFlag{
			Name:  fmt.Sprintf("%s, %s", outputFolderFlagName, outputFolderShortName),
			Usage: "`FOLDER` relative to --input to build the site into",
			Value: "static",
		},
		cli.StringFlag{
			Name:  fmt.Sprintf("%s, %s", configPathFlagName, configPathShortName),
			Usage: "`CONFIG` file relative to --input to build the site from",
			Value: "daybook.yml",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
}
