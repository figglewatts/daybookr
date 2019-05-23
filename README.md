# daybookr
SSG for journals. Quick and dirty, extremely barebones. Great for scatterbrained people in a rush.

## Installation
### Prerequisites
- Go
- `GOPATH` set up correctly
- `$GOPATH/bin` is on your `$PATH`

```bash
$ go install github.com/figglewatts/daybookr/cmd/daybookr
```

## Usage
```
NAME:
   daybookr - generate a journal based on some content, config and templates

USAGE:
   daybookr [global options]

VERSION:
   1.0

AUTHOR:
   Figglewatts

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --input FOLDER, -i FOLDER   use FOLDER to generate the site (default: ".")
   --output FOLDER, -o FOLDER  FOLDER relative to --input to build the site into (default: "static")
   --config CONFIG, -c CONFIG  CONFIG file relative to --input to build the site from (default: "daybook.yml")
   --help, -h                  show help
   --version, -v               print the version

COPYRIGHT:
   (c) 2019 Figglewatts
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

# Todo
- Documentation