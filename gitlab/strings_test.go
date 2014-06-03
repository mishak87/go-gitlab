// Copyright 2014 The go-gitlab AUTHORS. All rights reserved.
// Copyright 2013 The go-gitlab AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gitlab

import (
	"fmt"
	"testing"
	"time"
)

func TestStringify(t *testing.T) {
	var nilPointer *string

	var tests = []struct {
		in  interface{}
		out string
	}{
		// basic types
		{"foo", `"foo"`},
		{123, `123`},
		{1.5, `1.5`},
		{false, `false`},
		{
			[]string{"a", "b"},
			`["a" "b"]`,
		},
		{
			struct {
				A []string
			}{nil},
			// nil slice is skipped
			`{}`,
		},
		{
			struct {
				A string
			}{"foo"},
			// structs not of a named type get no prefix
			`{A:"foo"}`,
		},

		// pointers
		{nilPointer, `<nil>`},
		{String("foo"), `"foo"`},
		{Int(123), `123`},
		{Bool(false), `false`},
		{
			[]*string{String("a"), String("b")},
			`["a" "b"]`,
		},

		// actual GitHub structs
		{
			Timestamp{time.Date(2006, 01, 02, 15, 04, 05, 0, time.UTC)},
			`gitlab.Timestamp{2006-01-02 15:04:05 +0000 UTC}`,
		},
		{
			&Timestamp{time.Date(2006, 01, 02, 15, 04, 05, 0, time.UTC)},
			`gitlab.Timestamp{2006-01-02 15:04:05 +0000 UTC}`,
		},
		{
			User{ID: Int(123), Name: String("n")},
			`gitlab.User{ID:123, Name:"n"}`,
		},
		{
			Project{Owner: &User{ID: Int(123)}},
			`gitlab.Project{Owner:gitlab.User{ID:123}}`,
		},
	}

	for i, tt := range tests {
		s := Stringify(tt.in)
		if s != tt.out {
			t.Errorf("%d. Stringify(%q) => %q, want %q", i, tt.in, s, tt.out)
		}
	}
}

// Directly test the String() methods on various GitHub types.  We don't do an
// exaustive test of all the various field types, since TestStringify() above
// takes care of that.  Rather, we just make sure that Stringify() is being
// used to build the strings, which we do by verifying that pointers are
// stringified as their underlying value.
func TestString(t *testing.T) {
	var tests = []struct {
		in  interface{}
		out string
	}{
		// {CommitAuthor{Name: String("n")}, `gitlab.CommitAuthor{Name:"n"}`},
		// {CommitFile{SHA: String("s")}, `gitlab.CommitFile{SHA:"s"}`},
		// {CommitStats{Total: Int(1)}, `gitlab.CommitStats{Total:1}`},
		// {CommitsComparison{TotalCommits: Int(1)}, `gitlab.CommitsComparison{TotalCommits:1}`},
		// {Commit{SHA: String("s")}, `gitlab.Commit{SHA:"s"}`},
		// {Event{ID: String("1")}, `gitlab.Event{ID:"1"}`},
		// {GistComment{ID: Int(1)}, `gitlab.GistComment{ID:1}`},
		// {GistFile{Size: Int(1)}, `gitlab.GistFile{Size:1}`},
		// {Gist{ID: String("1")}, `gitlab.Gist{ID:"1", Files:map[]}`},
		// {GitObject{SHA: String("s")}, `gitlab.GitObject{SHA:"s"}`},
		// {Gitignore{Name: String("n")}, `gitlab.Gitignore{Name:"n"}`},
		// {Hook{ID: Int(1)}, `gitlab.Hook{Config:map[], ID:1}`},
		// {IssueComment{ID: Int(1)}, `gitlab.IssueComment{ID:1}`},
		// {Issue{Number: Int(1)}, `gitlab.Issue{Number:1}`},
		{Key{ID: Int(1)}, `gitlab.Key{ID:1}`},
		// {Label{Name: String("l")}, "l"},
		{Group{ID: Int(1)}, `gitlab.Group{ID:1}`},
		// {PullRequestComment{ID: Int(1)}, `gitlab.PullRequestComment{ID:1}`},
		// {PullRequest{Number: Int(1)}, `gitlab.PullRequest{Number:1}`},
		// {PushEventCommit{SHA: String("s")}, `gitlab.PushEventCommit{SHA:"s"}`},
		// {PushEvent{PushID: Int(1)}, `gitlab.PushEvent{PushID:1}`},
		// {Reference{Ref: String("r")}, `gitlab.Reference{Ref:"r"}`},
		// {ReleaseAsset{ID: Int(1)}, `gitlab.ReleaseAsset{ID:1}`},
		// {RepoStatus{ID: Int(1)}, `gitlab.RepoStatus{ID:1}`},
		// {RepositoryComment{ID: Int(1)}, `gitlab.RepositoryComment{ID:1}`},
		// {RepositoryCommit{SHA: String("s")}, `gitlab.RepositoryCommit{SHA:"s"}`},
		// {RepositoryContent{Name: String("n")}, `gitlab.RepositoryContent{Name:"n"}`},
		// {RepositoryRelease{ID: Int(1)}, `gitlab.RepositoryRelease{ID:1}`},
		{Project{ID: Int(1)}, `gitlab.Project{ID:1}`},
		// {Team{ID: Int(1)}, `gitlab.Team{ID:1}`},
		// {TreeEntry{SHA: String("s")}, `gitlab.TreeEntry{SHA:"s"}`},
		// {Tree{SHA: String("s")}, `gitlab.Tree{SHA:"s"}`},
		{User{ID: Int(1)}, `gitlab.User{ID:1}`},
		// {WebHookAuthor{Name: String("n")}, `gitlab.WebHookAuthor{Name:"n"}`},
		// {WebHookCommit{ID: String("1")}, `gitlab.WebHookCommit{ID:"1"}`},
		// {WebHookPayload{Ref: String("r")}, `gitlab.WebHookPayload{Ref:"r"}`},
	}

	for i, tt := range tests {
		s := tt.in.(fmt.Stringer).String()
		if s != tt.out {
			t.Errorf("%d. String() => %q, want %q", i, tt.in, tt.out)
		}
	}
}
