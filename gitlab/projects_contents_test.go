// Copyright 2014 The go-gitlab AUTHORS. All rights reserved.
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

func TestProjectsService_GetFileContents(t *testing.T) {
	setup()
	defer teardown()

	// given
	mux.HandleFunc("/projects/1/repository/blobs/sha", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"filepath": "fp",
		})
		fmt.Fprintf(w, "file data")
	})

	content, _, err := client.Projects.GetFileContents(1, "sha", "fp")
	if err != nil {
		t.Errorf("Projects.GetFileContents returned error: %v", err)
	}

	want := "file data"
	if !reflect.DeepEqual(content.String(), want) {
		t.Errorf("Projects.GetFileContents returned %+v, want %+v", content, want)
	}
}
