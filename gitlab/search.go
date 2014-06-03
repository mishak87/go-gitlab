// Copyright 2014 The go-gitlab AUTHORS. All rights reserved.
// Copyright 2013 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gitlab

import (
	"fmt"

	qs "github.com/google/go-querystring/query"
)

// SearchService provides access to the search related functions
// in the GitLab API.
//
type SearchService struct {
	client *Client
}

// SearchOptions specifies optional parameters to the SearchService methods.
type SearchOptions struct {
	ListOptions
}

// Projects searches projects by name.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/projects.html#search-for-projects-by-name
func (s *SearchService) Projects(name string, opt *SearchOptions) (*[]Project, *Response, error) {
	result := new([]Project)
	resp, err := s.search("projects", name, opt, result)
	return result, resp, err
}

// Users searches users by name.
//
// TODO: GitLab API docs: http://doc.gitlab.com/ce/api/users.html#list-users

// Helper function that executes search queries against different
// GitLab search types (project, users)
func (s *SearchService) search(searchType string, query string, opt *SearchOptions, result interface{}) (*Response, error) {
	params, err := qs.Values(opt)
	if err != nil {
		return nil, err
	}
	u := fmt.Sprintf("%s/search/%s?%s", searchType, query, params.Encode())

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, result)
}
