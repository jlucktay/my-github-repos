# GitHub GraphQL API client

[![godoc](https://img.shields.io/badge/pkg.go.dev-godoc-00ADD8?logo=go)](https://pkg.go.dev/go.jlucktay.dev/my-github-repos)

## Goal

Get a list of all of repositories owned by me from GitHub, alongside and/or sorted by creation date.

## From the Explorer

```GraphQL
{
  repositoryOwner(login: "jlucktay") {
    login
    repositories(first: 100, isFork: false, orderBy: {field: CREATED_AT, direction: ASC}) {
      edges {
        node {
          createdAt
          name
        }
      }
    }
  }
}
```

## GitHub token

The [autoload](https://github.com/joho/godotenv#usage) feature of
[`github.com/joho/godotenv`](https://github.com/joho/godotenv) is used to look for a `.env` file and read a value for
`GITHUB_TOKEN` to authenticate with GitHub.

## TODO

### Doing

### Done

- get forks and not-forks as two separate queries
  - in Terraform, these would be maintained as two separate resources, one volatile and one less so
- Pagination (starter limit is 100 and we're almost there already)
