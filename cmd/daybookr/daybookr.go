package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/figglewatts/daybookr/internal/daybookr"
	"gopkg.in/urfave/cli.v1"
)

type daybookrArgs struct {
	BaseURL string
}

const (
	baseURLArgIndex = iota
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

func getArgs(c *cli.Context) (daybookrArgs, error) {
	baseURL := c.Args().Get(baseURLArgIndex)
	if baseURL == "" {
		return daybookrArgs{}, errors.New("required positional argument 'BASE-URL' must be present")
	}

	return daybookrArgs{
		BaseURL: baseURL,
	}, nil
}

func getFlags(c *cli.Context) daybookrFlags {
	inputFolder := c.String(inputFolderShortName)
	return daybookrFlags{
		InputFolder:  inputFolder,
		OutputFolder: path.Join(inputFolder, c.String(outputFolderFlagName)),
		ConfigPath:   path.Join(inputFolder, c.String(configPathFlagName)),
	}
}

func runDaybookr(c *cli.Context) error {
	// parse command line args
	args, err := getArgs(c)
	if err != nil {
		return err
	}

	// get optional flags
	flags := getFlags(c)

	// generate the site
	err = daybookr.Generate(args.BaseURL, flags.InputFolder, flags.OutputFolder, flags.ConfigPath)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "daybookr"
	app.Usage = "generate a journal based on some content, config and templates"
	app.UsageText = "daybookr [global options] <BASE-URL>"
	app.ArgsUsage = "BASE-URL - the URL that is the base of your site, used when making URLs, i.e. https://yoursite.com"
	app.Action = runDaybookr
	app.Author = "Figglewatts"
	app.Version = "0.1"
	app.Compiled = time.Now()
	app.Copyright = "(c) 2019 Figglewatts"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  fmt.Sprintf("%s, %s", inputFolderFlagName, inputFolderShortName),
			Usage: "folder to build the site from",
			Value: ".",
		},
		cli.StringFlag{
			Name:  fmt.Sprintf("%s, %s", outputFolderFlagName, outputFolderShortName),
			Usage: "folder to output static site to",
			Value: "static",
		},
		cli.StringFlag{
			Name:  fmt.Sprintf("%s, %s", configPathFlagName, configPathShortName),
			Usage: "config file to use when building site",
			Value: "daybook.yml",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
}
