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
   daybookr [global options] <BASE-URL>

VERSION:
   0.1

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

## Design
- includes folder that are loaded for every template
- pages in pages folder in root
- layouts
   - layouts of pages are HTML files with templated values available in Page struct
   - pages say which layout they want to use in YAML front matter
   - posts say which layout they want to use in YAML front matter (the same...)
   - template loaded with template.ParseFiles(layout, content)
   - executed against Page struct
- handling index
   - index is a page at the root of the folder, will be templated from Site struct

## Todo
- Pagination for long pages
- Generate front page (with all posts)
- Years are not sorted on the archive page
- Tags are not sorted on the tags page
- Potential for removal of site base url?