package htmlnode

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"unicode"
)

type HTMLTag string

type Node interface {
	ToString() string
	ToHTML() (string, error)
}

//HTMLNode

type HTMLNode struct {
	Tag      HTMLTag
	Value    string
	Children []Node
	Props    map[string]string
}

func (h *HTMLNode) ToHTML() (string, error) {
	return "", errors.New("Not implemented")
}

func (h *HTMLNode) PropsToHTML() string {
	if len(h.Props) == 0 {
		return ""
	}

	keys := make([]string, 0, len(h.Props))

	for k := range h.Props {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	htmlAttr := ""
	for _, key := range keys {
		htmlAttr += fmt.Sprintf("%s='%s' ", key, h.Props[key])
	}

	htmlAttr = strings.TrimRightFunc(htmlAttr, unicode.IsSpace)

	return htmlAttr
}

func (h *HTMLNode) ToString() string {
	return fmt.Sprintf("HTMLNode(%s, %s, %s, %s)", h.Tag, h.Value, h.Children, h.Props)
}

// LeafNode

type LeafNode struct {
	HTMLNode
}

func (l *LeafNode) ToString() string {
	return fmt.Sprintf("LeafNode(%s, %s, %s, %s)", l.Tag, l.Value, l.Children, l.Props)
}

func (l *LeafNode) ToHTML() (string, error) {
	if l.Tag != "img" {

		if l.Value == "" {
			return "", errors.New(fmt.Sprintf("Leaf node %s does not have value. %s", l.Tag, l.ToString()))
		}
	}

	if l.Tag == "" {
		return l.Value, nil
	}

	props := l.PropsToHTML()
	if props != "" {
		return fmt.Sprintf("<%s %s>%s</%s>", l.Tag, props, l.Value, l.Tag), nil
	}

	return fmt.Sprintf("<%s>%s</%s>", l.Tag, l.Value, l.Tag), nil
}

// ParentNode

type ParentNode struct {
	HTMLNode
}

func (p *ParentNode) toString() string {
	return fmt.Sprintf("ParentNode(%s, %s, %s, %s)", p.Tag, p.Value, p.Children, p.Props)
}

func (p *ParentNode) ToHTML() (string, error) {
	if p.Tag == "" {
		return "", errors.New("Parent node does not have a Tag")
	}

	if len(p.Children) == 0 {
		return "", errors.New("Parent node does not have Children")
	}

	childNodes := ""
	for _, child := range p.Children {
		childHTML, err := child.ToHTML()
		if err != nil {
			return "", err
		}
		childNodes += childHTML
	}

	props := p.PropsToHTML()
	if props != "" {
		return fmt.Sprintf("<%s %s>%s</%s>", p.Tag, props, childNodes, p.Tag), nil
	}

	return fmt.Sprintf("<%s>%s</%s>", p.Tag, childNodes, p.Tag), nil
}
