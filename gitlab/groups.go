// Copyright 2014 The go-gitlab AUTHORS. All rights reserved.
// Copyright 2013 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gitlab

import "fmt"

// GroupsService provides access to the group related functions
// in the GitLab API.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/groups.html
type GroupsService struct {
	client *Client
}

// Group represents a GitLab group.
type Group struct {
	ID      *int    `json:"id,omitempty"`
	Name    *string `json:"name,omitempty"`
	Path    *string `json:"path,omitempty"`
	OwnerID *int    `json:"owner_id,omitempty"`
}

func (g Group) String() string {
	return Stringify(g)
}

// List the groups. When the token is an admin it will list all groups,
// when it is a user, it will list only their groups.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/groups.html#list-project-groups
func (s *GroupsService) List() ([]Group, *Response, error) {
	u := "groups"

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	groups := new([]Group)
	resp, err := s.client.Do(req, groups)
	if err != nil {
		return nil, resp, err
	}

	return *groups, resp, err
}

// Get fetches a group by name.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/groups.html#details-of-a-group
func (s *GroupsService) Get(groupid int) (*Group, *Response, error) {
	u := fmt.Sprintf("groups/%v", groupid)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	group := new(Group)
	resp, err := s.client.Do(req, group)
	if err != nil {
		return nil, resp, err
	}

	return group, resp, err
}

// Create a group
//
// TODO: GitLab API docs: http://doc.gitlab.com/ce/api/groups.html#new-group

// Transfer a project to a group
//
// TODO: GitLab API docs: http://doc.gitlab.com/ce/api/groups.html#transfer-project-to-group

// Remove a group
//
// TODO: GitLab API docs: http://doc.gitlab.com/ce/api/groups.html#remove-group
