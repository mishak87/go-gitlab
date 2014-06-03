// Copyright 2014 The go-gitlab AUTHORS. All rights reserved.
// Copyright 2013 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gitlab

import "fmt"

// ProjectCommit represents a commit in a project.
type ProjectCommit struct {
	ID          *string `json:"id,omitempty"`
	AuthorName  *string `json:"author_name,omitempty"`
	AuthorEmail *string `json:"author_email,omitempty"`

	AuthoredDate  *Timestamp `json:"authored_date,omitempty"`
	CommittedDate *Timestamp `json:"comitted_date,omitempty"`
	CreatedAt     *Timestamp `json:"created_at,omitempty"`

	ParentIDs []string `json:"parent_ids,omitempty"`
	ShortID   *string  `json:"short_id,omitempty"`
	Title     *string  `json:"title,omitempty"`
}

func (p ProjectCommit) String() string {
	return Stringify(p)
}

// ProjectCommitDiff represents a file diff inside a commit
type ProjectCommitDiff struct {
	AMode       *string `json:"a_mode,omitempty"`
	BMode       *string `json:"b_mode,omitempty"`
	Diff        *string `json:"diff,omitempty"`
	NewPath     *string `json:"new_path,omitempty"`
	OldPath     *string `json:"old_path,omitempty"`
	NewFile     *bool   `json:"new_file,omitempty"`
	DeletedFile *bool   `json:"deleted_file,omitempty"`
	RenamedFile *bool   `json:"renamed_file,omitempty"`
}

func (p ProjectCommitDiff) String() string {
	return Stringify(p)
}

// ListCommits lists the commits of a project.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/commits.html#list-repository-commits
func (s *ProjectsService) ListCommits(projectID int) ([]ProjectCommit, *Response, error) {
	u := fmt.Sprintf("projects/%v/repository/commits", projectID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	commits := new([]ProjectCommit)
	resp, err := s.client.Do(req, commits)
	if err != nil {
		return nil, resp, err
	}

	return *commits, resp, err
}

// GetCommit fetches the specified commit, including all details about it.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/commits.html#get-a-single-commit
func (s *ProjectsService) GetCommit(projectID int, commitID string) (*ProjectCommit, *Response, error) {
	u := fmt.Sprintf("projects/%v/repository/commits/%v", projectID, commitID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	commit := new(ProjectCommit)
	resp, err := s.client.Do(req, commit)
	if err != nil {
		return nil, resp, err
	}

	return commit, resp, err
}

// Get the diff of a commit
//
// GitLab API docs: http://doc.gitlab.com/ce/api/commits.html#get-the-diff-of-a-commit
func (s *ProjectsService) GetCommitDiff(projectID int, commitID string) (*ProjectCommitDiff, *Response, error) {
	u := fmt.Sprintf("projects/%v/repository/commits/%v/diff", projectID, commitID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	diff := new(ProjectCommitDiff)
	resp, err := s.client.Do(req, diff)
	if err != nil {
		return nil, resp, err
	}

	return diff, resp, err
}

// Compares two different branches, tag, or commits
//
// TODO: GitLab API docs: http://doc.gitlab.com/ce/api/repositories.html#compare-branches-tags-or-commits
