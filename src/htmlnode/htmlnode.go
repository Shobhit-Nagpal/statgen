package htmlnode

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"unicode"
)

type Node interface {
  toString() string
  toHTML() (string, error)
}

//HTMLNode

type HTMLTag string

type HTMLNode struct {
	tag      HTMLTag
	value    string
	children []Node
	props    map[string]string
}

func (h *HTMLNode) toHTML() (string, error) {
	return "", errors.New("Not implemented")
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


// LeafNode

type LeafNode struct {
  HTMLNode
}

func (l *LeafNode) toString() string {
  return fmt.Sprintf("LeafNode(%s, %s, %s, %s)", l.tag, l.value, l.children, l.props)
}

func (l *LeafNode) toHTML() (string, error) {
  if l.value == "" {
    return "", errors.New("Leaf node does not have value")
  }

  if l.tag == "" {
    return l.value, nil
  }

  props := l.propsToHTML()
  if props != "" {
    return fmt.Sprintf("<%s %s>%s</%s>", l.tag, props, l.value, l.tag), nil
  }

  return fmt.Sprintf("<%s>%s</%s>", l.tag, l.value, l.tag), nil
}


// ParentNode

type ParentNode struct {
  HTMLNode
}


func (p *ParentNode) toString() string {
  return fmt.Sprintf("ParentNode(%s, %s, %s, %s)", p.tag, p.value, p.children, p.props)
}


func (p *ParentNode) toHTML() (string, error) {
  if p.tag == "" {
    return "", errors.New("Parent node does not have a tag")
  }

  if len(p.children) == 0 {
    return "", errors.New("Parent node does not have children")
  }

  childNodes := ""
  for _, child := range p.children {
    childHTML, err := child.toHTML()
    if err != nil {
      return "", err
    }
    childNodes += childHTML
  }

  props := p.propsToHTML()
  if props != "" {
    return fmt.Sprintf("<%s %s>%s</%s>", p.tag, props, childNodes, p.tag), nil
  }

  return fmt.Sprintf("<%s>%s</%s>", p.tag, childNodes, p.tag), nil
}
