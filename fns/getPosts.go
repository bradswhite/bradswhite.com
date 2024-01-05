package fns

import (
  "os"
  "io/ioutil"
)

func GetPosts() ([]Matter, error) {
  var err error
	var fds []os.FileInfo

  posts := []Matter{}
  src := "./posts"

	if fds, err = ioutil.ReadDir(src); err != nil {
		return nil, err
	}

	for _, fd := range fds {
    matter, _, err := GetPost(fd.Name())
    if err != nil {
      return nil, err
    }
    posts = append(posts, *matter)
  }

  return posts, nil
}
