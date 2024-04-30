package htmlnode

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"unicode"
)

type HTMLNode struct {
	tag      string
	value    string
	children []string
	props    map[string]string
}

func (h *HTMLNode) toHTML() error {
	return errors.New("Not implemented")
}

func (h *HTMLNode) propsToHTML() string {
	if len(h.props) == 0 {
		return ""
	}

  keys := make([]string, 0, len(h.props))

  for k := range h.props {
    keys = append(keys, k)
  }
  sort.Strings(keys)

	htmlAttr := ""
	for _, key := range keys {
		htmlAttr += fmt.Sprintf("%s='%s' ", key, h.props[key])
	}

	htmlAttr = strings.TrimRightFunc(htmlAttr, unicode.IsSpace)

	return htmlAttr
}

func (h *HTMLNode) toString() string {
	return fmt.Sprintf("HTMLNode(%s, %s, %s, %s)", h.tag, h.value, h.children, h.props)
}
