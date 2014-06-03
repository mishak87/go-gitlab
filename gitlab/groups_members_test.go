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

func TestGroupsService_ListMembers(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/groups/1/members", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id":1}]`)
	})

	members, _, err := client.Groups.ListMembers(1)
	if err != nil {
		t.Errorf("Groups.ListMembers returned error: %v", err)
	}

	want := []User{{ID: Int(1)}}
	if !reflect.DeepEqual(members, want) {
		t.Errorf("Groups.ListMembers returned %+v, want %+v", members, want)
	}
}

func TestGroupsService_RemoveMember(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/groups/1/members/2", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Groups.RemoveMember(1, 2)
	if err != nil {
		t.Errorf("Groups.RemoveMember returned error: %v", err)
	}
}
