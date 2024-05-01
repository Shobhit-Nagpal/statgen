package md

import "fmt"

const (
  TEXT_TYPE_TEXT = "text"
  TEXT_TYPE_BOLD = "bold"
  TEXT_TYPE_ITALIC = "italic"
  TEXT_TYPE_CODE = "code"
  TEXT_TYPE_LINK = "link"
  TEXT_TYPE_IMAGE = "image"
)


type MarkdownImage struct {
	Text string
	Url  string
}

func (m MarkdownImage) IsEqual(mi MarkdownImage) bool {
  if m.Text == mi.Text && m.Url == mi.Url {
    return true
  }

  return false
}

func (m MarkdownImage) ToString() string {
  return fmt.Sprintf("MarkdownImage(%s, %s)", m.Text, m.Url)
}

type MarkdownLink struct {
	Text string
	Url  string
}


func (m MarkdownLink) IsEqual(ml MarkdownLink) bool {
  if m.Text == ml.Text && m.Url == ml.Url {
    return true
  }

  return false
}

func (m MarkdownLink) ToString() string {
  return fmt.Sprintf("MarkdownLink(%s, %s)", m.Text, m.Url)
}
