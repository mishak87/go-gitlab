// Copyright 2014 The go-gitlab AUTHORS. All rights reserved.
// Copyright 2013 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Repository contents API methods.
// http://developer.github.com/v3/repos/contents/package gitlab

package gitlab

import (
	"fmt"
	"net/http"
	"reflect"

	"testing"
)

func TestSearchService_Projects(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/projects/search/blah", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"page":     "2",
			"per_page": "2",
		})

		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	opts := &SearchOptions{ListOptions: ListOptions{Page: 2, PerPage: 2}}
	result, _, err := client.Search.Projects("blah", opts)
	if err != nil {
		t.Errorf("Search.Projects returned error: %v", err)
	}
	want := &[]Project{
		Project{ID: Int(1)},
		Project{ID: Int(2)},
	}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("Search.Projects returned %+v, want %+v", result, want)
	}
}
