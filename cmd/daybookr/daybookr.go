package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"gopkg.in/urfave/cli.v1"
)

type daybookrArgs struct {
	BaseURL string
}

const (
	baseURLArgIndex = iota
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

func runDaybookr(c *cli.Context) error {
	args, err := getArgs(c)
	if err != nil {
		return err
	}

	fmt.Printf(args.BaseURL)

	// _, err = daybookr.CreatePost("template.txt", "first-entry.md")
	// if err != nil {
	// 	return fmt.Errorf("Could not open post: %v", err)
	// }
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
			Name:  "input, i",
			Usage: "folder to build the site from",
			Value: ".",
		},
		cli.StringFlag{
			Name:  "output, o",
			Usage: "folder to output static site to",
			Value: "static",
		},
		cli.StringFlag{
			Name:  "config, c",
			Usage: "config file to use when building site",
			Value: "daybook.yml",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
}
