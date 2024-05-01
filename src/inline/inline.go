package inline

import (
	"errors"
	"statgen/src/htmlnode"
	"statgen/src/md"
	"statgen/src/textnode"
	"strings"
)

func TextNodeToHTMLNode(tn textnode.TextNode) (*htmlnode.LeafNode, error) {
	switch tn.TextType {
	case md.TEXT_TYPE_TEXT:
		return &htmlnode.LeafNode{htmlnode.HTMLNode{Value: tn.Text}}, nil
	case md.TEXT_TYPE_BOLD:
		return &htmlnode.LeafNode{htmlnode.HTMLNode{Tag: "b", Value: tn.Text}}, nil
	case md.TEXT_TYPE_ITALIC:
		return &htmlnode.LeafNode{htmlnode.HTMLNode{Tag: "i", Value: tn.Text}}, nil
	case md.TEXT_TYPE_CODE:
		return &htmlnode.LeafNode{htmlnode.HTMLNode{Tag: "code", Value: tn.Text}}, nil
	case md.TEXT_TYPE_LINK:
		return &htmlnode.LeafNode{htmlnode.HTMLNode{Tag: "a", Value: tn.Text, Props: map[string]string{"href": tn.Url}}}, nil
	case md.TEXT_TYPE_IMAGE:
		return &htmlnode.LeafNode{htmlnode.HTMLNode{Tag: "img", Props: map[string]string{"src": tn.Url, "alt": tn.Text}}}, nil
	default:
		return nil, errors.New("Text Node type is not valid")
	}
}

func SplitNodesDelimiter(oldNodes []*textnode.TextNode, delimiter, textType string) ([]*textnode.TextNode, error) {
	textNodes := []*textnode.TextNode{}

	for _, node := range oldNodes {
		if node.TextType != md.TEXT_TYPE_TEXT {
			textNodes = append(textNodes, node)
			continue
		}

		strs := strings.Split(node.Text, delimiter)

		if len(strs) % 2 == 0 {
			return nil, errors.New("Invalid Markdown syntax")
		}

		for idx, str := range strs {
			if idx % 2 == 0 {
				textNodes = append(textNodes, &textnode.TextNode{Text: str, TextType: md.TEXT_TYPE_TEXT})
				continue
			}

			textNodes = append(textNodes, &textnode.TextNode{Text: str, TextType: textType})
		}
	}

	return textNodes, nil
}
