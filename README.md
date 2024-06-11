
# markdown-auto-catalog
GitHub action automatically update folder-based table of contents in documents.

This robot will help you scan the folder and add the directory in the form of a list to the specified location of the specified document.

An example repository using this tool: [minoic/stackoverflow-go-top-qa](https://github.com/minoic/stackoverflow-go-top-qa)

## Usage

Step 1

Add two tags `<!-- catalog -->` to the appropriate location in the document where you want to add the catalog. This way, when the GitHub Action workflow job is completed, the catalog will be automatically added or updated between these two tags.

```markdown
## Catalog

<!-- catalog -->

<!-- catalog -->
``` 

Step 2

Add workflow in your project by creating workflow file such as `.github/workflows/catalog.yml` with the code below:

```yaml
name: markdown-auto-catalog

on:
  push:
    branches: [ "master" ]
  workflow_dispatch:

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: minoic/markdown-auto-catalog@v1.0.2
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
| filter        |     Filename Regex filter   | Optional |    \\(.*\\).md    |

## Preview

**Code view:**

```markdown
## Catalog

<!-- catalog -->

- [0-intro.md](folder/0-intro.md)
- 1-chapter1
  - [article1.md](folder/1-chapter1/article1.md)
  - [article2.md](folder/1-chapter1/article2.md)
  - [article3.md](folder/1-chapter1/article3.md)
  - sub-catalog
    - [sub-catalog.md](folder/1-chapter1/sub-catalog/sub-catalog.md)
- 2-chapter2
  - [article4.md](folder/2-chapter2/article4.md)
  - [article5.md](folder/2-chapter2/article5.md)
  - [article6.md](folder/2-chapter2/article6.md)
  - [article7.md](folder/2-chapter2/article7.md)

<!-- catalog -->
```

**Reader view:**

>## Catalog
> 
> <!-- catalog -->
> 
> - [0-intro.md](folder/0-intro.md)
> - 1-chapter1
>   - [article1.md](folder/1-chapter1/article1.md)
>   - [article2.md](folder/1-chapter1/article2.md)
>   - [article3.md](folder/1-chapter1/article3.md)
>   - sub-catalog
>     - [sub-catalog.md](folder/1-chapter1/sub-catalog/sub-catalog.md)
> - 2-chapter2
>   - [article4.md](folder/2-chapter2/article4.md)
>   - [article5.md](folder/2-chapter2/article5.md)
>   - [article6.md](folder/2-chapter2/article6.md)
>   - [article7.md](folder/2-chapter2/article7.md)
> 
> <!-- catalog -->

## Notice

1. When there are no files to change, there will be no PRs.
2. The previous PR is updated when it has not been merged and there are new changes.
3. If you don't like the red × when no files to change, add `continue-on-error: true` below the step.
4. Empty folders or folders containing filtered files are also displayed in the directory, so save unnecessary files in other directories.
5. Tables of contents are sorted in ascending lexicographical order and can be named using a format such as "01 - article title.md" to ensure the table of contents is in the order you want.
