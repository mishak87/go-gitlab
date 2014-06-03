// Copyright 2014 The go-gitlab AUTHORS. All rights reserved.
// Copyright 2013 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gitlab

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestUser_marshall(t *testing.T) {
	testJSONMarshal(t, &User{}, "{}")

	u := &User{
		ID:               Int(1),
		Username:         String("l"),
		Email:            String("e"),
		Name:             String("n"),
		State:            String("s"),
		CreatedAt:        &Timestamp{referenceTime},
		Bio:              String("b"),
		Skype:            String("s"),
		Linkedin:         String("l"),
		Twitter:          String("t"),
		WebsiteURL:       String("w"),
		ExternUID:        String("e"),
		Provider:         String("p"),
		ThemeID:          Int(1),
		ColorSchemeID:    Int(1),
		IsAdmin:          Bool(false),
		CanCreateGroup:   Bool(true),
		CanCreateProject: Bool(false),
		AvatarURL:        String("a"),
		PrivateToken:     String("p"),
	}
	want := `{
		"username": "l",
		"id": 1,
		"email": "e",
		"name": "n",
		"state": "s",
		"created_at": ` + referenceTimeStr + `,
		"bio": "b",
		"skype": "s",
		"linkedin": "l",
		"twitter": "t",
		"website_url": "w",
		"extern_uid": "e",
		"provider": "p",
		"theme_id": 1,
		"color_scheme_id": 1,
		"is_admin": false,
		"can_create_group": true,
		"can_create_project": false,
		"avatar_url": "a",
		"private_token": "p"
	}`
	testJSONMarshal(t, u, want)
}

func TestUsersService_Get_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":1}`)
	})

	user, _, err := client.Users.Get("")
	if err != nil {
		t.Errorf("Users.Get returned error: %v", err)
	}

	want := &User{ID: Int(1)}
	if !reflect.DeepEqual(user, want) {
		t.Errorf("Users.Get returned %+v, want %+v", user, want)
	}
}

func TestUsersService_Get_specifiedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/u", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":1}`)
	})

	user, _, err := client.Users.Get("u")
	if err != nil {
		t.Errorf("Users.Get returned error: %v", err)
	}

	want := &User{ID: Int(1)}
	if !reflect.DeepEqual(user, want) {
		t.Errorf("Users.Get returned %+v, want %+v", user, want)
	}
}

func TestUsersService_Get_invalidUser(t *testing.T) {
	_, _, err := client.Users.Get("%")
	testURLParseError(t, err)
}

func TestUsersService_Edit(t *testing.T) {
	setup()
	defer teardown()

	input := &User{Name: String("n")}

	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		v := new(User)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "PUT")
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"id":1}`)
	})

	user, _, err := client.Users.Edit(input)
	if err != nil {
		t.Errorf("Users.Edit returned error: %v", err)
	}

	want := &User{ID: Int(1)}
	if !reflect.DeepEqual(user, want) {
		t.Errorf("Users.Edit returned %+v, want %+v", user, want)
	}
}

func TestUsersService_ListAll(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id":2}]`)
	})

	users, _, err := client.Users.ListAll()
	if err != nil {
		t.Errorf("Users.Get returned error: %v", err)
	}

	want := []User{{ID: Int(2)}}
	if !reflect.DeepEqual(users, want) {
		t.Errorf("Users.ListAll returned %+v, want %+v", users, want)
	}
}
