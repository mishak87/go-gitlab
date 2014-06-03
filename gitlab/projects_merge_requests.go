// Copyright 2014 The go-gitlab AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gitlab

// Get all merge requests for this project. The state parameter can be
// used to get only merge requests with a given state (opened, closed,
// or merged) or all of them (all). The pagination parameters page and
// per_page can be used to restrict the list of merge requests.
//
// TODO: GitLab API docs: http://doc.gitlab.com/ce/api/merge_requests.html#list-merge-requests

// Shows information about a single merge request.
//
// TODO: GitLab API docs: http://doc.gitlab.com/ce/api/merge_requests.html#get-single-mr

// Creates a new merge request.
//
// TODO: GitLab API docs: http://doc.gitlab.com/ce/api/merge_requests.html#create-mr

// Updates an existing merge request. You can change branches, title, or
// even close the merge request.
//
// TODO: GitLab API docs: http://doc.gitlab.com/ce/api/merge_requests.html#update-mr

// Merge changes submitted with MR usign this API.
//
// TODO: GitLab API docs: http://doc.gitlab.com/ce/api/merge_requests.html#accept-mr

// Adds a comment to a merge request.
//
// TODO: GitLab API docs: http://doc.gitlab.com/ce/api/merge_requests.html#post-comment-to-mr

// Gets all the comments associated with a merge request.
//
// TODO: GitLab API docs: http://doc.gitlab.com/ce/api/merge_requests.html#get-the-comments-on-a-mr
