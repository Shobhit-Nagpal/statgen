package textnode

import "fmt"

type TextNode struct {
  Text string
  TextType string
  Url string
}

func (t *TextNode) ToString() string {
  return fmt.Sprintf("TextNode(%s, %s, %s)", t.Text, t.TextType, t.Url)
}

func (t *TextNode) IsEqual(tn *TextNode) bool {
  if t.Text == tn.Text && t.TextType == tn.TextType && t.Url == tn.Url {
    return true
  }
  return false
}
