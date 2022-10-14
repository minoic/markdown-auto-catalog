# markdown-auto-catalog
GitHub action automatically update folder-based table of contents in documents.

This bot will help you scan folders and add directories to the specified location of the specified document as a list.

## Usage

Step 1

Add two flags `<!-- catalog -->` to the document you want to add catalog.So that the catalog while be automatically update while the job works.

```mardown

## Catalog

<!-- catalog -->

<!-- catalog -->

``` 

Step 2

Add workflow in your project by creating workflow file such as `.github/workflows/catalog.yml` with the code below:

```yaml
name: minoic-auto-catalog

on:
  push:
    branches: [ "master" ]
  workflow_dispatch:

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: minoic/markdown-auto-catalog@master
        with:
          content-path: 'test/folder'
          document-path: 'test/README.md'
          filter: '\(.*\).md'
        continue-on-error: true

```

| Argument      | Description | Required | Default |
|---------------|--------|----------|--------|
| content-path  |   Path of the documents to be listed in catalog.     | Yes      |        |
| document-path |     Path of the catalog document.   | Yes      |        |
| filter        |     Filename Regex filter   | Optional |    \(.*\).md    |

## Preview

```markdown

## Catalog

<!-- catalog -->

- [0-intro.md](test/folder/0-intro.md)
- 1-chapter1
  - [article1.md](test/folder/1-chapter1/article1.md)
  - [article2.md](test/folder/1-chapter1/article2.md)
  - [article3.md](test/folder/1-chapter1/article3.md)
- 2-chapter2
  - [article4.md](test/folder/2-chapter2/article4.md)
  - [article5.md](test/folder/2-chapter2/article5.md)
  - [article6.md](test/folder/2-chapter2/article6.md)

<!-- catalog -->

```
