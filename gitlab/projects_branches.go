// Copyright 2014 The go-gitlab AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gitlab

// Get a list of repository branches from a project, sorted by name alphabetically.
//
// TODO: GitLab API docs: http://doc.gitlab.com/ce/api/branches.html#list-repository-branches

// Get a singple project repository branch
//
// TODO: GitLab API docs: http://doc.gitlab.com/ce/api/branches.html#get-single-repository-branch

// Protects a single project repository branch. This is an idempotent
// function, protecting an already protected repository branch still
// returns a 200 Ok status code.
//
// TODO: GitLab API docs: http://doc.gitlab.com/ce/api/branches.html#protect-repository-branch

// Unprotects a single project repository branch. This is an idempotent
// function, unprotecting an already unprotected repository branch still
// returns a 200 Ok status code.
//
// TODO: GitLab API docs: http://doc.gitlab.com/ce/api/branches.html#unprotect-repository-branch

// Create a project repository branch
//
// TODO: GitLab API docs: http://doc.gitlab.com/ce/api/branches.html#create-repository-branch

// Delete a project repository branch
//
// TODO: GitLab API docs: http://doc.gitlab.com/ce/api/branches.html#delete-repository-branch
