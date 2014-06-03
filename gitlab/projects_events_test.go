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

func TestProjectsService_ListEvents(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/projects/1/events", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"project_id":1},{"project_id":1}]`)
	})

	events, _, err := client.Projects.ListEvents(1)

	if err != nil {
		t.Errorf("Projects.GetEvents returned error: %v", err)
	}

	want := []ProjectEvent{{ProjectID: Int(1)}, {ProjectID: Int(1)}}
	if !reflect.DeepEqual(events, want) {
		t.Errorf("Projects.GetEvents returned %+v, want %+v", events, want)
	}
}
