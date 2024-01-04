package fns

import (
  "os"
  "bufio"
  "github.com/adrg/frontmatter"
)

func GetPost(title string) (*Matter, []byte, error) {
  front := &Matter{}
  
  file, err := os.Open("./posts/" + title + ".md")
  if err != nil {
    return nil, nil, err
  }

  body, err := frontmatter.Parse(bufio.NewReader(file), &front)
  if err != nil {
    return nil, nil, err
  }

  return front, body, nil
}
