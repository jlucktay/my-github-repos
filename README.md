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

[A personal access token (classic) is used to access
GitHub](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens),
and must be set in the `GITHUB_TOKEN` environment variable.

## TODO

### Doing

### Done

- get forks and not-forks as two separate queries
  - in Terraform, these would be maintained as two separate resources, one volatile and one less so
- Pagination (starter limit is 100 and we're almost there already)
