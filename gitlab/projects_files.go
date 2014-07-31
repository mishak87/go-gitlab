// Copyright 2014 The go-gitlab AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gitlab

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type ProjectFileParameters struct {
	FilePath      string `json:"file_path,omitempty"`
	BranchName    string `json:"branch_name,omit_empty"`
	Encoding      string `json:"encoding,omit_empty"`
	Content       string `json:"content,omit_empty"`
	CommitMessage string `json:"commit_message,omit_empty"`
}

// Get file from project
//
// TODO: GitLab API docs: http://doc.gitlab.com/ce/api/repository_files.html#get-file-from-repository

// Create a new file in project
//
// TODO: GitLab API docs: http://doc.gitlab.com/ce/api/repository_files.html#create-new-file-in-repository

// Update an existing file in project
//
// TODO: GitLab API docs: http://doc.gitlab.com/ce/api/repository_files.html#update-existing-file-in-repository
func (s *ProjectsService) UpdateFile(projectID int, parameters ProjectFileParameters, content []byte) (interface{}, *Response, error) {
	u := fmt.Sprintf("projects/%d/repository/files", projectID)

	parameters.Encoding = "base64"
	parameters.Content = base64.StdEncoding.EncodeToString(content)
	content, err := json.Marshal(parameters)

	req, err := s.client.NewRequest("PUT", u, parameters)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	data := new(interface{})
	resp, err := s.client.Do(req, data)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, err
}

// Delete existing file in project
//
// TODO: GitLab API docs: http://doc.gitlab.com/ce/api/repository_files.html#delete-existing-file-in-repository
