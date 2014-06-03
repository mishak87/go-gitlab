// Copyright 2014 The go-gitlab AUTHORS. All rights reserved.
// Copyright 2013 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gitlab

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestProjectsService_ListCommits(t *testing.T) {
	setup()
	defer teardown()

	// given
	mux.HandleFunc("/projects/1/repository/commits", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprintf(w, `[{"id": "s"}]`)
	})

	commits, _, err := client.Projects.ListCommits(1)
	if err != nil {
		t.Errorf("Projects.ListCommits returned error: %v", err)
	}

	want := []ProjectCommit{{ID: String("s")}}
	if !reflect.DeepEqual(commits, want) {
		t.Errorf("Projects.ListCommits returned %+v, want %+v", commits, want)
	}
}

func TestRepositoriesService_GetCommit(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/projects/1/repository/commits/s", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprintf(w, `{
			"id": "s",
			"short_id": "s",
			"title": "t",
			"author_name": "a",
			"author_email": "a",
			"parent_ids": [
				"a",
				"b"
			]
		}`)
	})

	commit, _, err := client.Projects.GetCommit(1, "s")
	if err != nil {
		t.Errorf("Projects.GetCommit returned error: %v", err)
	}

	want := &ProjectCommit{
		ID:          String("s"),
		AuthorName:  String("a"),
		AuthorEmail: String("a"),
		ShortID:     String("s"),
		ParentIDs:   []string{"a", "b"},
		Title:       String("t"),
	}
	if !reflect.DeepEqual(commit, want) {
		t.Errorf("Projects.GetCommit returned \n%+v, want \n%+v", commit, want)
	}
}

func TestRepositoriesService_CompareCommits(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/projects/1/repository/commits/s/diff", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprintf(w, `{
			"diff": "d",
			"new_path": "n",
			"old_path": "o",
			"a_mode": "1",
			"b_mode": "1",
			"new_file": false,
			"renamed_file": false,
			"deleted_file": false
		}`)
	})

	got, _, err := client.Projects.GetCommitDiff(1, "s")
	if err != nil {
		t.Errorf("Projects.GetCommitDiff returned error: %v", err)
	}

	want := &ProjectCommitDiff{
		Diff:        String("d"),
		NewPath:     String("n"),
		AMode:       String("1"),
		BMode:       String("1"),
		OldPath:     String("o"),
		NewFile:     Bool(false),
		RenamedFile: Bool(false),
		DeletedFile: Bool(false),
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Projects.GetCommitDiff returned \n%+v, want \n%+v", got, want)
	}
}
