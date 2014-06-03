# go-gitlab #

go-gitlab is a Go client library for accessing the GitLab API, it's a fork of the go-github library [go-github](https://github.com/google/go-github). Thanks go to everyone who has contributed to the go-github project for creating the client library used to build this client.

go-gitlab requires Go version 1.1 or greater.

## Usage ##

```go
import "github.com/bigkraig/go-gitlab/gitlab"
```

Create a new GitLab client and search for a project:

```go
client := gitlab.NewClient("http://gitlab.example.com/api/v3/", "API SECRET KEY")
opts := &gitlab.SearchOptions{gitlab.ListOptions{Page: 1}}
projects, _, err := client.Search.Projects("my project", opts)
```

For complete usage of go-gitlab, see the full [package docs][].

[GitLab API]: http://doc.gitlab.com/ce/api/
[package docs]: http://godoc.org/github.com/bigkraig/go-gitlab/gitlab



## License ##

This library is distributed under the BSD-style license found in the [LICENSE](./LICENSE)
file.
