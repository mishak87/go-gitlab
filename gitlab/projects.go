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

// ProjectsService handles communication with the project related
// methods of the GitLab API.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/projects.html
type ProjectsService struct {
	client *Client
}

// ForkedFromProject represents the source project for a project.
type ForkedFromProject struct {
	ID                *int    `json:"id,omitempty"`
	Name              *string `json:"name,omitempty"`
	NameWithNamespace *string `json:"name_with_namespace,omitempty"`
	Path              *string `json:"path,omitempty"`
	PathWithNamespace *string `json:"path_with_namespace,omitempty"`
}

// Project represents a GitLab project.
type Project struct {
	ID                *int               `json:"id,omitempty"`
	Archived          *bool              `json:"archived,omitempty"`
	DefaultBranch     *string            `json:"default_branch,omitempty"`
	Description       *string            `json:"description,omitempty"`
	HTTPURLToRepo     *string            `json:"http_url_to_repo,omitempty"`
	ImportURL         *string            `json:"import_url,omitempty"`
	LastActivityAt    *Timestamp         `json:"last_activity_at,omitempty"`
	ForkedFromProject *ForkedFromProject `json:"forked_from_project,omitempty"`
	Name              *string            `json:"name,omitempty"`
	NamespaceID       *int               `json:"namespace_id,omitempty"`
	NameWithNamespace *string            `json:"name_with_namespace,omitempty"`
	Namespace         *Namespace         `json:"namespace,omitempty"`
	Owner             *User              `json:"owner,omitempty"`
	Path              *string            `json:"path,omitempty"`
	PathWithNamespace *string            `json:"path_with_namespace,omitempty"`
	SSHURLToRepo      *string            `json:"ssh_url_to_repo,omitempty"`
	UserID            *int               `json:"user_id,omitempty"`
	VisibilityLevel   *int               `json:"visibility_level,omitempty"`
	WebURL            *string            `json:"web_url,omitempty"`

	// Additional mutable fields when creating and editing a project
	IssuesEnabled        *bool `json:"issues_enabled,omitempty"`
	MergeRequestsEnabled *bool `json:"merge_requests_enabled,omitempty"`
	Public               *bool `json:"public,omitempty"`
	SnippetsEnabled      *bool `json:"snippets_enabled,omitempty"`
	WallEnabled          *bool `json:"wall_enabled,omitempty"`
	WikiEnabled          *bool `json:"wiki_enabled,omitempty"`
}

func (p Project) String() string {
	return Stringify(p)
}

// List the repositories accessible by authenticated user. Passing the empty string will list
// repositories for the authenticated user.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/projects.html#list-projects
func (s *ProjectsService) List(opt *ListOptions) ([]Project, *Response, error) {
	return s.list("projects", opt)
}

// ListOwned lists all GitLab projects that are owned by the user.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/projects.html#list-owned-projects
func (s *ProjectsService) ListOwned(opt *ListOptions) ([]Project, *Response, error) {
	return s.list("projects/owned", opt)
}

// ListAll lists all GitLab projects. For admins only.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/projects.html#list-all-projects
func (s *ProjectsService) ListAll(opt *ListOptions) ([]Project, *Response, error) {
	return s.list("projects/all", opt)
}

// Helper function for list functions
func (s *ProjectsService) list(listType string, opt *ListOptions) ([]Project, *Response, error) {
	params, err := qs.Values(opt)
	if err != nil {
		return nil, nil, err
	}

	u := fmt.Sprintf("%s?%s", listType, params.Encode())
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	projects := new([]Project)
	resp, err := s.client.Do(req, projects)
	if err != nil {
		return nil, resp, err
	}

	return *projects, resp, err
}

// Create a new project. If Project.UserID is defined, it will create
// the project under that user, provided you are an admin.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/projects.html#create-project
//                  http://doc.gitlab.com/ce/api/projects.html#create-project-for-user
func (s *ProjectsService) Create(project *Project) (*Project, *Response, error) {
	u := "projects"
	if project.UserID != nil {
		project.NamespaceID = nil
		u = fmt.Sprintf("projects/user/%v", *project.UserID)
	}

	req, err := s.client.NewRequest("POST", u, project)
	if err != nil {
		return nil, nil, err
	}

	p := new(Project)
	resp, err := s.client.Do(req, p)
	if err != nil {
		return nil, resp, err
	}

	return p, resp, err
}

// Get fetches a project.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/projects.html#get-single-project
func (s *ProjectsService) Get(projectid int) (*Project, *Response, error) {
	u := fmt.Sprintf("projects/%v", projectid)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	project := new(Project)
	resp, err := s.client.Do(req, project)
	if err != nil {
		return nil, resp, err
	}

	return project, resp, err
}

// Delete a project.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/projects.html#remove-project
func (s *ProjectsService) Delete(project int) (*Response, error) {
	u := fmt.Sprintf("projects/%v", project)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
