package textnode

import "fmt"

type TextNode struct {
  text string
  textType string
  url string
}

func (t *TextNode) toString() string {
  return fmt.Sprintf("TextNode(%s, %s, %s)", t.text, t.textType, t.url)
}

func (t *TextNode) isEqual(tn TextNode) bool {
  if t.text == tn.text && t.textType == tn.textType && t.url == tn.url {
    return true
  }
  return false
}
