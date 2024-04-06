package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/suri312006/blogposts"
)

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oi maite, we dont work ova here")
}

func TestNewBlogPosts(t *testing.T) {

	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: rust, borrowchecker
---
man im so tired of everything
what is wron gwith me`

		secondBody = `Title: Post 2
Description: Description 2
Tags: lonely
---
wahoo!! i love 
my
girlfriend`
	)

	fs := fstest.MapFS{
		"hello_world.md":  {Data: []byte(firstBody)},
		"hello_world2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}

	assertPost(t, posts[0], blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags: []string{"rust", "borrowchecker" },
		Body: `man im so tired of everything
what is wron gwith me`,

	})
	assertPost(t, posts[1], blogposts.Post{
		Title:       "Post 2",
		Description: "Description 2",
		Tags: []string{"lonely"},
		Body: `wahoo!! i love 
my
girlfriend`,
	})
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v posts, wanted %+v posts", got, want)
	}
}
