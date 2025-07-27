# tock-md-to-csv

Simple tool to transform a list of markdown files to a csv file adapted to **Tock** rag import.

Markdown files should have mininum metadata in their header, `tiltle` and `site_url` like this:


```markdown
---
title: My Website online Documentation
site_url: https://www.mywebsite.com
---

## My doc title

This is the content of my doc.


ccxcccakldklzedkzed
zedd;,el,ekzekd,kfkrnfjrnhhriéh"héuheé"e

```

Transformation will result in a csv file with the following columns:

```csv
title|source|text
My Website online Documentation|https://www.mywebsite.com|"# My doc title ......
My second doc title|https://www.mywebsite.com/second-doc|"# My second doc title ......
```

## Mkdocs UseCase

if you are using [Mkdocs](https://www.mkdocs.org/) to generate your documentation, you can use the option `--ismkdoc --baseurl` to generate a csv where urls are computed base on mkdocs pages folder.
```shell
$ md-to-csv -f docs -c output.csv --ismkdoc --baseurl https://www.mywebsite.com/
```

## Usage

If your markdown files are in a folder called `samples` and you want to create a csv file called `output.csv`, you can run the following command:

```shell
$ md-to-csv -f samples -c output.csv
```

More options are available, you can see them by running:

```shell
$ md-to-csv --help
```

```console
$ md-to-csv -h 
Générate a CSV file from markdown files. The CSV file will contain the title, the source and the text of each markdown file.
        
        Example usage:
        md-to-csv -f samples -c output.csv

Usage:
  md-to-csv [flags]
  md-to-csv [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  version     Provide md-to-csv version and build number

Flags:
  -u, --base string     Base URL for Mkdocs files (default "http://localhost:9000")
  -c, --csv string      CSV file to generate (default "output.csv")
  -f, --folder string   Folder containing markdown files (default ".")
  -h, --help            help for md-to-csv
  -m, --ismkdoc         Folder is an Mkdocs source

Use "md-to-csv [command] --help" for more information about a command.
```