package inline

import (
	"errors"
	"fmt"
	"regexp"
	"statgen/src/htmlnode"
	"statgen/src/md"
	"statgen/src/textnode"
	"strings"
)

func TextNodeToHTMLNode(tn *textnode.TextNode) (*htmlnode.LeafNode, error) {
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

		if len(strs)%2 == 0 {
			return nil, errors.New("Invalid Markdown syntax")
		}

		for idx, str := range strs {
			if idx%2 == 0 {
				textNodes = append(textNodes, &textnode.TextNode{Text: str, TextType: md.TEXT_TYPE_TEXT})
				continue
			}

			textNodes = append(textNodes, &textnode.TextNode{Text: str, TextType: textType})
		}
	}

	return textNodes, nil
}

func ExtractMarkdownImages(text string) []md.MarkdownImage {
	mdImages := []md.MarkdownImage{}

	re := regexp.MustCompile(`!\[(.*?)\]\((.*?)\)`)
	textRe := regexp.MustCompile(`!\[(.*?)\]`)
	urlRe := regexp.MustCompile(`\((.*?)\)`)

	images := re.FindAllString(text, -1)

	for _, image := range images {
		img := md.MarkdownImage{}
		imageText := textRe.FindAllString(image, 1)[0]
		imageUrl := urlRe.FindAllString(image, 1)[0]
		img.Text = strings.Trim(imageText, "![]")
		img.Url = strings.Trim(imageUrl, "()")
		mdImages = append(mdImages, img)
	}

	return mdImages
}

func ExtractMarkdownLinks(text string) []md.MarkdownLink {
	mdLinks := []md.MarkdownLink{}

	re := regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)
	textRe := regexp.MustCompile(`\[(.*?)\]`)
	urlRe := regexp.MustCompile(`\((.*?)\)`)

	links := re.FindAllString(text, -1)

	for _, link := range links {
		lnk := md.MarkdownLink{}
		lnkText := textRe.FindAllString(link, 1)[0]
		lnkUrl := urlRe.FindAllString(link, 1)[0]
		lnk.Text = strings.Trim(lnkText, "![]")
		lnk.Url = strings.Trim(lnkUrl, "()")
		mdLinks = append(mdLinks, lnk)
	}
	return mdLinks
}

func SplitNodesImage(nodes []*textnode.TextNode) ([]*textnode.TextNode, error) {
	newNodes := []*textnode.TextNode{}

	for _, node := range nodes {
		if node.TextType != md.TEXT_TYPE_TEXT {
			newNodes = append(newNodes, node)
			continue
		}

		images := ExtractMarkdownImages(node.Text)
		size := len(images)
		if size == 0 {
			newNodes = append(newNodes, node)
			continue
		}

		imageText := node.Text
		for _, image := range images {
			sep := fmt.Sprintf("![%s](%s)", image.Text, image.Url)
			sections := strings.Split(imageText, sep)

			if len(sections) != 2 {
				return nil, errors.New("Invalid Markdown syntax. Image section is not closed")
			}

			if sections[0] != "" {
				newNodes = append(newNodes, &textnode.TextNode{Text: sections[0], TextType: md.TEXT_TYPE_TEXT})
			}

			newNodes = append(newNodes, &textnode.TextNode{Text: image.Text, TextType: md.TEXT_TYPE_IMAGE, Url: image.Url})
			imageText = sections[1]

		}

		if imageText != "" {
			newNodes = append(newNodes, &textnode.TextNode{Text: imageText, TextType: md.TEXT_TYPE_TEXT})
		}
	}

	return newNodes, nil
}

func SplitNodesLink(nodes []*textnode.TextNode) ([]*textnode.TextNode, error) {
	newNodes := []*textnode.TextNode{}

	for _, node := range nodes {
		if node.TextType != md.TEXT_TYPE_TEXT {
			newNodes = append(newNodes, node)
			continue
		}

		links := ExtractMarkdownLinks(node.Text)
		size := len(links)
		if size == 0 {
			newNodes = append(newNodes, node)
			continue
		}

		linkText := node.Text
		for _, link := range links {
			sep := fmt.Sprintf("[%s](%s)", link.Text, link.Url)
			sections := strings.Split(linkText, sep)

			if len(sections) != 2 {
				return nil, errors.New("Invalid Markdown syntax. Link section is not closed")
			}

			if sections[0] != "" {
				newNodes = append(newNodes, &textnode.TextNode{Text: sections[0], TextType: md.TEXT_TYPE_TEXT})
			}

			newNodes = append(newNodes, &textnode.TextNode{Text: link.Text, TextType: md.TEXT_TYPE_LINK, Url: link.Url})
			linkText = sections[1]

		}

		if linkText != "" {
			newNodes = append(newNodes, &textnode.TextNode{Text: linkText, TextType: md.TEXT_TYPE_TEXT})
		}
	}

	return newNodes, nil
}

func TextToTextNodes(text string) ([]*textnode.TextNode, error) {
  newNodes, err := SplitNodesDelimiter([]*textnode.TextNode{&textnode.TextNode{Text: text, TextType: md.TEXT_TYPE_TEXT}}, "**", md.TEXT_TYPE_BOLD)
  if err != nil {
    return nil, err
  }


  newNodes, err = SplitNodesDelimiter(newNodes, "*", md.TEXT_TYPE_ITALIC)
  if err != nil {
    return nil, err
  }

  newNodes, err = SplitNodesDelimiter(newNodes, "`", md.TEXT_TYPE_CODE)
  if err != nil {
    return nil, err
  }

  newNodes, err = SplitNodesImage(newNodes)
  if err != nil {
    return nil, err
  }

  newNodes, err = SplitNodesLink(newNodes)
  if err != nil {
    return nil, err
  }
  return newNodes, nil
}
