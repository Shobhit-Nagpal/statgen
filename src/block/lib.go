package block

import (
	"errors"
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
	BLOCK_TYPE_TABLE          = "table"
)

func CreateHTMLHeading(block string) (string, error) {
	content := strings.SplitN(block, " ", 2)
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
		content := strings.SplitN(item, " ", 2)
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
		}

		listItems += fmt.Sprintf("<li>%s</li>", listValue)
	}

	ol := fmt.Sprintf("<ol>%s</ol>", listItems)
	return ol, nil
}

func CreateHTMLUnorderedList(block string) (string, error) {
	items := strings.Split(block, "\n")
	listItems := ""

	for _, item := range items {
		content := strings.SplitN(item, " ", 2)
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
		}

		listItems += fmt.Sprintf("<li>%s</li>", listValue)
	}

	ul := fmt.Sprintf("<ul>%s</ul>", listItems)
	return ul, nil
}

func CreateHTMLTable(block string) (string, error) {
	items := strings.Split(block, "\n")

	if len(items) < 2 {
		return "", errors.New("Table must have atleast two lines")
	}

	if strings.HasPrefix(items[1], "|---") || strings.HasPrefix(items[1], "| --") {
		items = append(items[:1], items[1+1:]...)
	}

	tableItems := ""

	tableHeadings := items[0]
	//Parse table heading here and append to tableItems
	tableHeaders := strings.Split(tableHeadings, "|")
	tableHeaders = tableHeaders[1:]
	tableHeaders = tableHeaders[:len(tableHeaders)-1]

	headers := ""
	for _, tableHeader := range tableHeaders {
		tableHeading, err := CreateHTMLHeading(fmt.Sprintf("### %s", strings.TrimSpace(tableHeader)))
		if err != nil {
			return "", err
		}

		headers += fmt.Sprintf("<th>%s</th>", tableHeading)
	}

	tableItems += fmt.Sprintf("<tr>%s</tr>", headers)

	//Remove the table headings
	items = items[1:]

	for _, tableRow := range items {
		//Parse each table row here and append to table items
		tableData := strings.Split(tableRow, "|")
		tableData = tableData[1:]
		tableData = tableData[:len(tableData)-1]

		td := ""
		for _, data := range tableData {
			tableDataItem, err := CreateHTMLParagraph(strings.TrimSpace(data))
			if err != nil {
				return "", err
			}

			td += fmt.Sprintf("<td>%s</td>", tableDataItem)
		}

		tableItems += fmt.Sprintf("<tr>%s</tr>", td)
	}

	table := fmt.Sprintf("<table>%s</table>", tableItems)
	return table, nil
}
