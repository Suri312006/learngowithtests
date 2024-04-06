package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
	// "testing/fstest"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func NewPostFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")

	if err != nil {
		return nil, err
	}

	var posts []Post

	for _, f := range dir {
		post, err := getPost(fileSystem, f.Name())

		if err != nil {
			return nil, err

		}

		posts = append(posts, post)
	}

	return posts, nil

}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, err := fileSystem.Open(fileName)

	if err != nil {
		return Post{}, err
	}
	// like how you should close the file after reading
	defer postFile.Close()

	return newPost(postFile)

}

const (
	titleSeperator       = "Title: "
	descriptionSeperator = "Description: "
	TagsSeperator        = "Tags: "
)


func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readLine := func(seperatorName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), seperatorName)
	}

	return Post{
		Title:       readLine(titleSeperator),
		Description: readLine(descriptionSeperator),
		Tags:        strings.Split(readLine(TagsSeperator), ", "),
		Body: readBody(scanner),
	}, nil

}

// extract tags from text?
func readBody(scanner *bufio.Scanner) string { scanner.Scan() // scan past the body line

	buf := bytes.Buffer{}

	// will keep scanning until theres nothing to can??
	for scanner.Scan(){
		// it prints to this buffer instead of actual console output
		fmt.Fprintln(&buf, scanner.Text())

	}
	// trims spaces from output

	body := strings.TrimSuffix(buf.String(), "\n")
	return body
}
