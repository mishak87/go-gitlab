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

func TestProjectsService_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	projects, _, err := client.Projects.List()
	if err != nil {
		t.Errorf("Projects.List returned error: %v", err)
	}

	want := []Project{{ID: Int(1)}, {ID: Int(2)}}
	if !reflect.DeepEqual(projects, want) {
		t.Errorf("Projects.List returned %+v, want %+v", projects, want)
	}
}

func TestProjectsService_ListOwned(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/projects/owned", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id":1}]`)
	})

	projects, _, err := client.Projects.ListOwned()
	if err != nil {
		t.Errorf("Projects.ListAll returned error: %v", err)
	}

	want := []Project{{ID: Int(1)}}
	if !reflect.DeepEqual(projects, want) {
		t.Errorf("Projects.ListAll returned %+v, want %+v", projects, want)
	}
}

func TestProjectsService_ListAll(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/projects/all", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id":1}]`)
	})

	projects, _, err := client.Projects.ListAll()
	if err != nil {
		t.Errorf("Projects.ListAll returned error: %v", err)
	}

	want := []Project{{ID: Int(1)}}
	if !reflect.DeepEqual(projects, want) {
		t.Errorf("Projects.ListAll returned %+v, want %+v", projects, want)
	}
}

func TestProjectsService_Create_user(t *testing.T) {
	setup()
	defer teardown()

	input := &Project{Name: String("n"), UserID: Int(1), NamespaceID: Int(2)}

	mux.HandleFunc("/projects/user/1", func(w http.ResponseWriter, r *http.Request) {
		v := new(Project)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"id":1}`)
	})

	project, _, err := client.Projects.Create(input)
	if err != nil {
		t.Errorf("Projects.Create returned error: %v", err)
	}

	want := &Project{ID: Int(1)}
	if !reflect.DeepEqual(project, want) {
		t.Errorf("Projects.Create returned %+v, want %+v", project, want)
	}
}

func TestProjectsService_Create(t *testing.T) {
	setup()
	defer teardown()

	input := &Project{Name: String("n")}

	mux.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		v := new(Project)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"id":1}`)
	})

	project, _, err := client.Projects.Create(input)
	if err != nil {
		t.Errorf("Projects.Create returned error: %v", err)
	}

	want := &Project{ID: Int(1)}
	if !reflect.DeepEqual(project, want) {
		t.Errorf("Projects.Create returned %+v, want %+v", project, want)
	}
}

func TestRepositoriesService_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/projects/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":1,"name":"n","description":"d","owner":{"username":"l"}}`)
	})

	project, _, err := client.Projects.Get(1)
	if err != nil {
		t.Errorf("Projects.Get returned error: %v", err)
	}

	want := &Project{ID: Int(1), Name: String("n"), Description: String("d"), Owner: &User{Username: String("l")}}
	if !reflect.DeepEqual(project, want) {
		t.Errorf("Projects.Get returned %+v, want %+v", project, want)
	}
}

// func TestRepositoriesService_Edit(t *testing.T) {
// 	setup()
// 	defer teardown()

// 	i := true
// 	input := &Repository{IssuesEnabled: &i}

// 	mux.HandleFunc("/repos/o/r", func(w http.ResponseWriter, r *http.Request) {
// 		v := new(Repository)
// 		json.NewDecoder(r.Body).Decode(v)

// 		testMethod(t, r, "PATCH")
// 		if !reflect.DeepEqual(v, input) {
// 			t.Errorf("Request body = %+v, want %+v", v, input)
// 		}
// 		fmt.Fprint(w, `{"id":1}`)
// 	})

// 	repo, _, err := client.Repositories.Edit("o", "r", input)
// 	if err != nil {
// 		t.Errorf("Repositories.Edit returned error: %v", err)
// 	}

// 	want := &Repository{ID: Int(1)}
// 	if !reflect.DeepEqual(repo, want) {
// 		t.Errorf("Repositories.Edit returned %+v, want %+v", repo, want)
// 	}
// }

func TestProjectsService_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/projects/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Projects.Delete(1)
	if err != nil {
		t.Errorf("Projects.Delete returned error: %v", err)
	}
}

// func TestRepositoriesService_Get_invalidOwner(t *testing.T) {
// 	_, _, err := client.Repositories.Get("%", "r")
// 	testURLParseError(t, err)
// }

// func TestRepositoriesService_Edit_invalidOwner(t *testing.T) {
// 	_, _, err := client.Repositories.Edit("%", "r", nil)
// 	testURLParseError(t, err)
// }

// func TestRepositoriesService_ListContributors(t *testing.T) {
// 	setup()
// 	defer teardown()

// 	mux.HandleFunc("/repos/o/r/contributors", func(w http.ResponseWriter, r *http.Request) {
// 		testMethod(t, r, "GET")
// 		testFormValues(t, r, values{
// 			"anon": "true",
// 			"page": "2",
// 		})
// 		fmt.Fprint(w, `[{"contributions":42}]`)
// 	})

// 	opts := &ListContributorsOptions{Anon: "true", ListOptions: ListOptions{Page: 2}}
// 	contributors, _, err := client.Repositories.ListContributors("o", "r", opts)

// 	if err != nil {
// 		t.Errorf("Repositories.ListContributors returned error: %v", err)
// 	}

// 	want := []Contributor{{Contributions: Int(42)}}
// 	if !reflect.DeepEqual(contributors, want) {
// 		t.Errorf("Repositories.ListContributors returned %+v, want %+v", contributors, want)
// 	}
// }

// func TestRepositoriesService_ListLanguages(t *testing.T) {
// 	setup()
// 	defer teardown()

// 	mux.HandleFunc("/repos/o/r/languages", func(w http.ResponseWriter, r *http.Request) {
// 		testMethod(t, r, "GET")
// 		fmt.Fprint(w, `{"go":1}`)
// 	})

// 	languages, _, err := client.Repositories.ListLanguages("o", "r")
// 	if err != nil {
// 		t.Errorf("Repositories.ListLanguages returned error: %v", err)
// 	}

// 	want := map[string]int{"go": 1}
// 	if !reflect.DeepEqual(languages, want) {
// 		t.Errorf("Repositories.ListLanguages returned %+v, want %+v", languages, want)
// 	}
// }

// func TestRepositoriesService_ListTags(t *testing.T) {
// 	setup()
// 	defer teardown()

// 	mux.HandleFunc("/repos/o/r/tags", func(w http.ResponseWriter, r *http.Request) {
// 		testMethod(t, r, "GET")
// 		testFormValues(t, r, values{"page": "2"})
// 		fmt.Fprint(w, `[{"name":"n", "commit" : {"sha" : "s", "url" : "u"}, "zipball_url": "z", "tarball_url": "t"}]`)
// 	})

// 	opt := &ListOptions{Page: 2}
// 	tags, _, err := client.Repositories.ListTags("o", "r", opt)
// 	if err != nil {
// 		t.Errorf("Repositories.ListTags returned error: %v", err)
// 	}

// 	want := []RepositoryTag{
// 		{
// 			Name: String("n"),
// 			Commit: &Commit{
// 				SHA: String("s"),
// 				URL: String("u"),
// 			},
// 			ZipballURL: String("z"),
// 			TarballURL: String("t"),
// 		},
// 	}
// 	if !reflect.DeepEqual(tags, want) {
// 		t.Errorf("Repositories.ListTags returned %+v, want %+v", tags, want)
// 	}
// }

// func TestRepositoriesService_ListBranches(t *testing.T) {
// 	setup()
// 	defer teardown()

// 	mux.HandleFunc("/repos/o/r/branches", func(w http.ResponseWriter, r *http.Request) {
// 		testMethod(t, r, "GET")
// 		testFormValues(t, r, values{"page": "2"})
// 		fmt.Fprint(w, `[{"name":"master", "commit" : {"sha" : "a57781", "url" : "https://api.github.com/repos/o/r/commits/a57781"}}]`)
// 	})

// 	opt := &ListOptions{Page: 2}
// 	branches, _, err := client.Repositories.ListBranches("o", "r", opt)
// 	if err != nil {
// 		t.Errorf("Repositories.ListBranches returned error: %v", err)
// 	}

// 	want := []Branch{{Name: String("master"), Commit: &Commit{SHA: String("a57781"), URL: String("https://api.github.com/repos/o/r/commits/a57781")}}}
// 	if !reflect.DeepEqual(branches, want) {
// 		t.Errorf("Repositories.ListBranches returned %+v, want %+v", branches, want)
// 	}
// }

// func TestRepositoriesService_GetBranch(t *testing.T) {
// 	setup()
// 	defer teardown()

// 	mux.HandleFunc("/repos/o/r/branches/b", func(w http.ResponseWriter, r *http.Request) {
// 		testMethod(t, r, "GET")
// 		fmt.Fprint(w, `{"name":"n", "commit":{"sha":"s"}}`)
// 	})

// 	branch, _, err := client.Repositories.GetBranch("o", "r", "b")
// 	if err != nil {
// 		t.Errorf("Repositories.GetBranch returned error: %v", err)
// 	}

// 	want := &Branch{Name: String("n"), Commit: &Commit{SHA: String("s")}}
// 	if !reflect.DeepEqual(branch, want) {
// 		t.Errorf("Repositories.GetBranch returned %+v, want %+v", branch, want)
// 	}
// }

// func TestRepositoriesService_ListLanguages_invalidOwner(t *testing.T) {
// 	_, _, err := client.Repositories.ListLanguages("%", "%")
// 	testURLParseError(t, err)
// }
