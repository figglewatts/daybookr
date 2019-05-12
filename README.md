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
```bash
$ daybookr
```

## Design
- `daybookr` executable installable via `go install ...` can be run to generate a site based on a bunch of templates in the folder it's run in (or optionally by giving a folder as a cmd line argument)
    - site gets output to `./static/` by default, or a folder given by a cmd line arg
    - need method for converting a page into a URL?
    - site base name passed in as cmd line arg (for building URLs)
- `./templates/*.html` - a bunch of templates that the site uses to generate content
    - other templates can be included in these templates
- `./entries/*.md` - journal entries
- `./pages/*.md` - custom pages content
- `daybook.yml` - config file for a particular site
    - loaded as arbitrary YAML
    - could include:
        - title and subtitle of journal
        - pages and their templates/content (to show up in the navbar)
        - author name
        - links in footer

## Todo
- Command line integration
    - args for custom folder
    - args for config override
    - args for custom output dir
    - arg for site base name
- Load in config file
- Site data structure to hold config/args
- URL generator based on site base name
- Render site pages based on template combination and config
- Generate front page (with all posts)
- Generate bespoke pages from config
- Generate archive page
- Generate tag pages