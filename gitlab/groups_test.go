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

func TestGroupsService_List_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/groups", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	groups, _, err := client.Groups.List()
	if err != nil {
		t.Errorf("Groups.List returned error: %v", err)
	}

	want := []Group{{ID: Int(1)}, {ID: Int(2)}}
	if !reflect.DeepEqual(groups, want) {
		t.Errorf("Groups.List returned %+v, want %+v", groups, want)
	}
}

func TestGroupsService_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/groups/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":1, "name":"n"}`)
	})

	group, _, err := client.Groups.Get(1)
	if err != nil {
		t.Errorf("Groups.Get returned error: %v", err)
	}

	want := &Group{ID: Int(1), Name: String("n")}
	if !reflect.DeepEqual(group, want) {
		t.Errorf("Groups.Get returned %+v, want %+v", group, want)
	}
}
