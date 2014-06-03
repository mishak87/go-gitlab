// Copyright 2014 The go-gitlab AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gitlab

import "fmt"

// ProjectEvent represents a GitLab project event.
type ProjectEvent struct {
	Title       *string           `json:"title,omitempty"`
	ProjectID   *int              `json:"project_id,omitempty"`
	ActionName  *string           `json:"action_name,omitempty"`
	TargetID    *int              `json:"target_id,omitempty"`
	AuthorID    *int              `json:"author_id,omitempty"`
	Data        *ProjectEventData `json:"data,omitempty"`
	TargetTitle *string           `json:"target_title"`
	// TargetType ???
}

// ProjectEventData represents a project event data blob.
type ProjectEventData struct {
	Before           *string                     `json:"before,omitempty"`
	After            *string                     `json:"after,omitempty"`
	Ref              *string                     `json:"ref,omitempty"`
	UserID           *int                        `json:"user_id,omitempty"`
	UserName         *string                     `json:"user_name,omitempty"`
	Repository       *ProjectEventDataRepository `json:"repository,omitempty"`
	Commits          []ProjectEventDataCommit    `json:"commits,omitempty"`
	TotalCommitCount *int                        `json:"total_commits_count,omitempty"`
}

// ProjectEventDataRepository represents a repository in a project event data blob.
type ProjectEventDataRepository struct {
	Name        *string `json:"name,omitempty"`
	URL         *string `json:"url,omitempty"`
	Description *string `json:"description,omitempty"`
	Homepage    *string `json:"homepage,omitempty"`
}

// ProjectEventDataCommit represents a commit in a project event data blob.
type ProjectEventDataCommit struct {
	ID        *string                       `json:"id,omitempty"`
	Message   *string                       `json:"message,omitempty"`
	Timestamp *Timestamp                    `json:"timestamp,omitempty"`
	URL       *string                       `json:"url,omitempty"`
	Author    *ProjectEventDataCommitAuthor `json:"author,omitempty"`
}

// ProjectEventDataCommitAuthor represents the author of a commit in a project event data blob.
type ProjectEventDataCommitAuthor struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}

func (p ProjectEvent) String() string {
	return Stringify(p)
}

// Get project events.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/projects.html#get-project-events
func (s *ProjectsService) ListEvents(projectid int) ([]ProjectEvent, *Response, error) {
	u := fmt.Sprintf("projects/%v/events", projectid)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	events := new([]ProjectEvent)
	resp, err := s.client.Do(req, events)
	if err != nil {
		return nil, resp, err
	}

	return *events, resp, err
}
