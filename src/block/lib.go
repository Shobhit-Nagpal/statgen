package block

import (
	"fmt"
	"statgen/src/htmlnode"
	"statgen/src/inline"
	"strings"
)

const (
	BLOCK_TYPE_PARAGRAPH      = "paragraph"
	BLOCK_TYPE_HEADING        = "heading"
	BLOCK_TYPE_CODE           = "code"
	BLOCK_TYPE_QUOTE          = "quote"
	BLOCK_TYPE_UNORDERED_LIST = "unordered_list"
	BLOCK_TYPE_ORDERED_LIST   = "ordered_list"
)

func CreateHTMLHeading(block string) (string, error) {
	content := strings.SplitN(block, " ", 1)
	headingNumber := len(content[0])

	textNodes, err := inline.TextToTextNodes(content[1])
	if err != nil {
		return "", err
	}

	leafNodes := []*htmlnode.LeafNode{}
	for _, textNode := range textNodes {
		leafNode, err := inline.TextNodeToHTMLNode(textNode)
		if err != nil {
			return "", err
		}

		leafNodes = append(leafNodes, leafNode)
	}

	headingValue := ""

	for _, leafNode := range leafNodes {
		leafHTML, err := leafNode.ToHTML()
		if err != nil {
			return "", err
		}
		headingValue += leafHTML
	}

	heading := fmt.Sprintf("<h%d>%s</h%d>", headingNumber, headingValue, headingNumber)
	return heading, nil
}

func CreateHTMLParagraph(block string) (string, error) {
	pValue := ""

	textNodes, err := inline.TextToTextNodes(block)
	if err != nil {
		return "", err
	}

	leafNodes := []*htmlnode.LeafNode{}
	for _, textNode := range textNodes {
		leafNode, err := inline.TextNodeToHTMLNode(textNode)
		if err != nil {
			return "", err
		}

		leafNodes = append(leafNodes, leafNode)
	}

	for _, leafNode := range leafNodes {
		leafHTML, err := leafNode.ToHTML()
		if err != nil {
			return "", err
		}
		pValue += leafHTML
	}

	p := fmt.Sprintf("<p>%s</p>", pValue)
	return p, nil
}

func CreateHTMLQuote(block string) (string, error) {
	contents := strings.Split(block, "\n")
	quoteBlocks := []string{}

	for _, content := range contents {
		quoteBlocks = append(quoteBlocks, strings.Split(content, ">")[1])
	}
	quoteValue := ""
	for _, block := range quoteBlocks {
		quoteValue += block + "\n"
	}

	quote := "<blockquote>"
	quote += quoteValue + "</blockquote>"

	return quote, nil
}

func CreateHTMLCode(block string) (string, error) {
	codeValue := ""

	textNodes, err := inline.TextToTextNodes(block)
	if err != nil {
		return "", err
	}

	leafNodes := []*htmlnode.LeafNode{}
	for _, textNode := range textNodes {
		leafNode, err := inline.TextNodeToHTMLNode(textNode)
		if err != nil {
			return "", err
		}

		leafNodes = append(leafNodes, leafNode)
	}

	for _, leafNode := range leafNodes {
		leafHTML, err := leafNode.ToHTML()
		if err != nil {
			return "", err
		}
		codeValue += leafHTML
	}

	code := fmt.Sprintf("<pre>%s</pre>", strings.TrimSpace(codeValue))
	return code, nil

}

func CreateHTMLOrderedList(block string) (string, error) {
	items := strings.Split(block, "\n")
  listItems := ""

	for _, item := range items {
		content := strings.SplitN(item, " ", 1)
		listValue := ""

		textNodes, err := inline.TextToTextNodes(content[1])
		if err != nil {
			return "", err
		}

		leafNodes := []*htmlnode.LeafNode{}
		for _, textNode := range textNodes {
			leafNode, err := inline.TextNodeToHTMLNode(textNode)
			if err != nil {
				return "", err
			}

			leafNodes = append(leafNodes, leafNode)
		}

		for _, leafNode := range leafNodes {
			leafHTML, err := leafNode.ToHTML()
			if err != nil {
				return "", err
			}
			listValue += leafHTML
      listItems += fmt.Sprintf("<li>%s</li>", listValue)
		}

	}

  ol := fmt.Sprintf("<ol>%s</ol>", listItems)
  return ol, nil
}

func CreateHTMLUnorderedList(block string) (string, error) {
	items := strings.Split(block, "\n")
  listItems := ""

	for _, item := range items {
		content := strings.SplitN(item, " ", 1)
		listValue := ""

		textNodes, err := inline.TextToTextNodes(content[1])
		if err != nil {
			return "", err
		}

		leafNodes := []*htmlnode.LeafNode{}
		for _, textNode := range textNodes {
			leafNode, err := inline.TextNodeToHTMLNode(textNode)
			if err != nil {
				return "", err
			}

			leafNodes = append(leafNodes, leafNode)
		}

		for _, leafNode := range leafNodes {
			leafHTML, err := leafNode.ToHTML()
			if err != nil {
				return "", err
			}
			listValue += leafHTML
      listItems += fmt.Sprintf("<li>%s</li>", listValue)
		}

	}

  ol := fmt.Sprintf("<ul>%s</ul>", listItems)
  return ol, nil
}
