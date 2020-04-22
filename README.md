# daybookr
Simple static site generator. Quick and dirty, extremely barebones. Great for scatterbrained people in a rush.

## Installation
```bash
$ go get -u github.com/figglewatts/daybookr/cmd/daybookr
```

## Usage
```
NAME:
   daybookr - generate a static site based on some content, config and templates

USAGE:
   daybookr [global options] COMMAND [options]

VERSION:
   2.0.3

AUTHOR:
   Figglewatts

COMMANDS:
   generate  generate the site
   new       create a new site
   open      open a daybookr site in a browser
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)

COPYRIGHT:
   (c) 2020 Figglewatts
```

## Advanced usage

### Using custom config fields in layouts
If you have custom fields in your `daybook.yml` file, i.e.:
```yaml
custom_field: this is a string
```
These can be accessed in a layout by doing something like this:
```html
<h1>{{(.Site.Conf.Get "custom_field").String}}</h1>
```
Resulting in:
```html
<h1>this is a string</h1>
```
The config file is parsed with the [simpleyaml](https://github.com/smallfish/simpleyaml) package. For more information on how to get all kinds of YAML values from the config file (i.e. not just a string as above), please see their documentation.
