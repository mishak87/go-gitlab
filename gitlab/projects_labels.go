// Copyright 2014 The go-gitlab AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gitlab

import "fmt"

// Label represents a GitLab label.
type Label struct {
	Name *string `json:"name,omitempty"`
}

func (l Label) String() string {
	return Stringify(l)
}

// List project labels.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/projects.html#list-project-labels
func (s *ProjectsService) ListLabels(projectid int) ([]Label, *Response, error) {
	u := fmt.Sprintf("projects/%v/labels", projectid)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	labels := new([]Label)
	resp, err := s.client.Do(req, labels)
	if err != nil {
		return nil, resp, err
	}

	return *labels, resp, err
}
