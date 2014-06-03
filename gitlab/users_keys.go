// Copyright 2014 The go-gitlab AUTHORS. All rights reserved.
// Copyright 2013 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gitlab

import "fmt"

// Key represents a public SSH key used to authenticate a user or deploy script.
type Key struct {
	ID        *int       `json:"id,omitempty"`
	Key       *string    `json:"key,omitempty"`
	CreatedAt *Timestamp `json:"created_at,omitempty"`
	Title     *string    `json:"title,omitempty"`
}

func (k Key) String() string {
	return Stringify(k)
}

// ListKeys lists the verified public keys for a user.  Passing the empty
// string will fetch keys for the authenticated user.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/users.html#list-ssh-keys
func (s *UsersService) ListKeys(user string, opt *ListOptions) ([]Key, *Response, error) {
	var u string
	if user != "" {
		u = fmt.Sprintf("users/%v/keys", user)
	} else {
		u = "user/keys"
	}
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	keys := new([]Key)
	resp, err := s.client.Do(req, keys)
	if err != nil {
		return nil, resp, err
	}

	return *keys, resp, err
}

// GetKey fetches a single public key.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/users.html#single-ssh-key
func (s *UsersService) GetKey(id int) (*Key, *Response, error) {
	u := fmt.Sprintf("user/keys/%v", id)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	key := new(Key)
	resp, err := s.client.Do(req, key)
	if err != nil {
		return nil, resp, err
	}

	return key, resp, err
}

// CreateKey adds a public key for the authenticated user.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/users.html#add-ssh-key
func (s *UsersService) CreateKey(key *Key) (*Key, *Response, error) {
	return s.CreateKeyForUser(0, key)
}

// CreateKeyForUser adds a public key for the specified user.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/users.html#add-ssh-key-for-user
func (s *UsersService) CreateKeyForUser(uid int, key *Key) (*Key, *Response, error) {
	var u string
	if uid != 0 {
		u = fmt.Sprintf("users/%v/keys", uid)
	} else {
		u = "user/keys"
	}

	req, err := s.client.NewRequest("POST", u, key)
	if err != nil {
		return nil, nil, err
	}

	k := new(Key)
	resp, err := s.client.Do(req, k)
	if err != nil {
		return nil, resp, err
	}

	return k, resp, err
}

// DeleteKey deletes a public key.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/users.html#delete-ssh-key
func (s *UsersService) DeleteKey(id int) (*Response, error) {
	return s.DeleteKeyForUser(0, id)
}

// DeleteKeyForUser deletes a public key for the specified user.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/users.html#delete-ssh-key
func (s *UsersService) DeleteKeyForUser(uid int, id int) (*Response, error) {
	var u string
	if uid != 0 {
		u = fmt.Sprintf("users/%v/keys/%v", uid, id)
	} else {
		u = fmt.Sprintf("user/keys/%v", id)
	}

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
