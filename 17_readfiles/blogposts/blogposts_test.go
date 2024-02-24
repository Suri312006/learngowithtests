package blogposts_test

import (
	"errors"
	"testing"
	"testing/fstest"
	"io/fs"
	blogposts "github.com/suri312006/blogposts"
)

type StubFailingFS struct {

}

func (s StubFailingFS) Open(name string) (fs.File, error){
	return nil, errors.New("oi maite, we dont work ova here")
}

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello_world.md":  {Data: []byte("hi")},
		"hello_world2.md": {Data: []byte("weeniw")},
	}

	posts, err := blogposts.NewPostFromFS(fs)

	if err != nil{
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}
}
