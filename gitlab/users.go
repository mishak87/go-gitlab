// Copyright 2014 The go-gitlab AUTHORS. All rights reserved.
// Copyright 2013 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gitlab

import "fmt"

// UsersService handles communication with the user related
// methods of the GitLab API.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/users.html
type UsersService struct {
	client *Client
}

// User represents a GitLab user.
type User struct {
	Username         *string    `json:"username,omitempty"`
	ID               *int       `json:"id,omitempty"`
	Email            *string    `json:"email,omitempty"`
	Name             *string    `json:"name,omitempty"`
	State            *string    `json:"state,omitempty"`
	CreatedAt        *Timestamp `json:"created_at,omitempty"`
	Bio              *string    `json:"bio,omitempty"`
	Skype            *string    `json:"skype,omitempty"`
	Linkedin         *string    `json:"linkedin,omitempty"`
	Twitter          *string    `json:"twitter,omitempty"`
	WebsiteURL       *string    `json:"website_url,omitempty"`
	ExternUID        *string    `json:"extern_uid,omitempty"`
	Provider         *string    `json:"provider,omitempty"`
	ThemeID          *int       `json:"theme_id,omitempty"`
	ColorSchemeID    *int       `json:"color_scheme_id,omitempty"`
	IsAdmin          *bool      `json:"is_admin,omitempty"`
	CanCreateGroup   *bool      `json:"can_create_group,omitempty"`
	CanCreateProject *bool      `json:"can_create_project,omitempty"`
	AvatarURL        *string    `json:"avatar_url,omitempty"`
	PrivateToken     *string    `json:"private_token,omitempty"`
	AccessLevel      *int       `json:"access_level,omitempty"`
}

func (u User) String() string {
	return Stringify(u)
}

// Get fetches a user.  Passing the empty string will fetch the authenticated
// user.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/users.html#single-user
func (s *UsersService) Get(user string) (*User, *Response, error) {
	var u string
	if user != "" {
		u = fmt.Sprintf("users/%v", user)
	} else {
		u = "user"
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(User)
	resp, err := s.client.Do(req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, err
}

// Edit the authenticated user.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/users.html#user-modification
func (s *UsersService) Edit(user *User) (*User, *Response, error) {
	u := "user"
	req, err := s.client.NewRequest("PUT", u, user)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(User)
	resp, err := s.client.Do(req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, err
}

// User creation
//
// TODO: GitLab API docs: http://doc.gitlab.com/ce/api/users.html#user-creation

// User deletion
//
// TODO: GitLab API docs: http://doc.gitlab.com/ce/api/users.html#user-deletion

// ListAll lists all GitLab users.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/users.html#list-users
func (s *UsersService) ListAll() ([]User, *Response, error) {
	u := "users"
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	users := new([]User)
	resp, err := s.client.Do(req, users)
	if err != nil {
		return nil, resp, err
	}

	return *users, resp, err
}
