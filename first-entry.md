---
date: 2019-3-30
tags: [journal]
---

# Journalling

So this is the first entry in my online journal... Except the online portion of the journal doesn't exactly exist yet.
My goal for this is to make a static site generator (SSG) that will take in markdown files like this one and generate
a really simple barebones site out of them. I think I'll probably call it 'daybookr' because it's still hip and cool to
skip the 'e' in a name like that, right?

## Tech
- Go (good templating engine, also I want to practice!)
- https://github.com/urfave/cli for CLI
- https://github.com/russross/blackfriday for markdown
- GitLab pages

## Features
- Dir of entries + Dir of templates + Config file = Online journal site
- Automatic generation of archive page and pages for each entry
- Automatic generation of tag index page
- Looks good (is readable) on mobile

## Useful links
- https://docs.gitlab.com/ee/user/project/pages/getting_started_part_four.html
- https://about.gitlab.com/2016/06/10/ssg-overview-gitlab-pages-part-2/
- https://golang.org/pkg/text/template/
- https://golang.org/pkg/html/template/
- https://docs.gitlab.com/ee/user/project/pages/getting_started_part_two.html#create-a-project-from-scratch
- https://css-tricks.com/snippets/css/a-guide-to-flexbox/
- https://css-tricks.com/couple-takes-sticky-footer/#article-header-id-3

## Considerations
- How to handle multiple entries on a single date?
- How do we make it look good on mobile?

## Work chunks
- Make a website layout (templates)
- Iteratively replace templates with data using Go
- Improve the website to support all features (mobile, archives, tags, entries etc.)
- Improve the SSG to support all of these features
- Hook the SSG up to GitLab pages to generate a site on a push to a repo

So, off to work. I'll add more stuff to the sections above as I find it -- and I'll add more content to this entry
as I develop it!

## GitLab pages deployment
A GitLab pages repo deploys the website from folder `public` in the repo.
For GitLab pages I'll need a CI job that grabs/installs the SSG executable, runs it on the repo folder, and dumps
generated content in `public`.

## Site design
- I wanna go for really simple, like monospace font and no frills and such. Kinda like old livejournals, I guess, but more programmer-y.
- 'Modern' HTML/CSS used for stuff, like `<article>` and `flex`.

## Code
```python
def func(arg1, arg2):
    print(f"This is just a big test {arg1}")
```