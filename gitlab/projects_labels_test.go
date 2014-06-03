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

func TestProjectsService_ListLabels(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/projects/1/labels", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"name":"n"}]`)
	})

	labels, _, err := client.Projects.ListLabels(1)

	if err != nil {
		t.Errorf("Projects.ListLabels returned error: %v", err)
	}

	want := []Label{{Name: String("n")}}
	if !reflect.DeepEqual(labels, want) {
		t.Errorf("Projects.ListLabels returned %+v, want %+v", labels, want)
	}
}
