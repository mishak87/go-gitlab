// Copyright 2014 The go-gitlab AUTHORS. All rights reserved.
// Copyright 2013 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gitlab

import "fmt"

// ListMembers lists the members for a group.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/groups.html#group-members
func (s *GroupsService) ListMembers(groupid int) ([]User, *Response, error) {
	u := fmt.Sprintf("groups/%v/members", groupid)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	members := new([]User)
	resp, err := s.client.Do(req, members)
	if err != nil {
		return nil, resp, err
	}

	return *members, resp, err
}

// AddMember adds a member to a group.
//
// TODO: GitLab API docs: http://doc.gitlab.com/ce/api/groups.html#add-group-member

// RemoveMember removes a user from a group.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/groups.html#remove-user-team-member
func (s *GroupsService) RemoveMember(groupid, userid int) (*Response, error) {
	u := fmt.Sprintf("groups/%v/members/%v", groupid, userid)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
