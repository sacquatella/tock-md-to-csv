# tock-md-to-csv

Simple tool to transform a list of markdown files to a csv file adapted to **Tock** rag import.

Markdown files should have mininum metadata in their header, `tiltle` and `site_url` like this:


```markdown
---
title: My Website online Documentation
site_url: https://www.mywebsite.com
---

# My doc title

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


# Usage

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

Usage:
  md-to-csv [flags]
  md-to-csv [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  version     Provide md-to-csv version and build number

Flags:
  -c, --csv string      CSV file to generate (default "output.csv")
  -f, --folder string   Folder containing markdown files (default ".")
  -h, --help            help for md-to-csv

Use "md-to-csv [command] --help" for more information about a command.
```