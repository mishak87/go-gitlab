// Copyright 2014 The go-gitlab AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gitlab

// Namespace represents a GitLab namespace.
type Namespace struct {
	CreatedAt   *Timestamp `json:"created_at,omitempty"`
	Description *string    `json:"description,omitempty"`
	ID          *int       `json:"id,omitempty"`
	Name        *string    `json:"name,omitempty"`
	OwnerID     *int       `json:"owner_id,omitempty"`
	Path        *string    `json:"path,omitempty"`
	UpdatedAt   *Timestamp `json:"updated_at,omitempty"`
}
