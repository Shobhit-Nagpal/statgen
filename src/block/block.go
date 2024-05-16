package block

import (
	"errors"
//	"fmt"
	"regexp"
	"strings"
)

func MarkdownToBlocks(markdown string) []string {
	blocks := []string{}
	mdBlocks := strings.Split(markdown, "\n\n")
	for _, mdBlock := range mdBlocks {
		if mdBlock != "" {
			mdBlock = strings.TrimSpace(mdBlock)
			blocks = append(blocks, mdBlock)
		}
	}
	return blocks
}

func BlockToBlockType(block string) string {
	headingRe := regexp.MustCompile(`^#{1,6}\s.*$`)
	quoteRe := regexp.MustCompile(`(?m:^>.*?$)`)
	ulRe := regexp.MustCompile(`(?m:^[*-]\s.*?$)`)
	olRe := regexp.MustCompile(`(?m:^(1|[2-9]\d*)\.\s.*?$)`)
	codeRe := regexp.MustCompile("^```[\\s\\S]*?```$")
	tableRe := regexp.MustCompile(`(?m:^\|\s.*?\s\|\s.*?\s\|$)`)
	hrRe := regexp.MustCompile(`^---$`)

	if headingRe.MatchString(block) {
		return BLOCK_TYPE_HEADING
	} else if quoteRe.MatchString(block) {
		return BLOCK_TYPE_QUOTE
	} else if ulRe.MatchString(block) {
		return BLOCK_TYPE_UNORDERED_LIST
	} else if olRe.MatchString(block) {
		return BLOCK_TYPE_ORDERED_LIST
	} else if codeRe.MatchString(block) {
		return BLOCK_TYPE_CODE
	} else if tableRe.MatchString(block) {
		return BLOCK_TYPE_TABLE
	} else if hrRe.MatchString(block) {
    return BLOCK_TYPE_HR
  } else {
		return BLOCK_TYPE_PARAGRAPH
	}
}

func MarkdownToHTMLNode(markdown string) (string, error) {
	html := "<div>"

	blocks := MarkdownToBlocks(markdown)

	for _, block := range blocks {
		blockType := BlockToBlockType(block)
		switch blockType {
		case BLOCK_TYPE_HEADING:
			heading, err := CreateHTMLHeading(block)
			if err != nil {
				return "", err
			}

			html += heading
		case BLOCK_TYPE_PARAGRAPH:
			p, err := CreateHTMLParagraph(block)
			if err != nil {
				return "", err
			}

			html += p
		case BLOCK_TYPE_QUOTE:
			quote, err := CreateHTMLQuote(block)
			if err != nil {
				return "", err
			}

			html += quote
		case BLOCK_TYPE_CODE:
			code, err := CreateHTMLCode(block)
			if err != nil {
				return "", err
			}

			html += code
		case BLOCK_TYPE_UNORDERED_LIST:
			ul, err := CreateHTMLUnorderedList(block)
			if err != nil {
				return "", err
			}

			html += ul
		case BLOCK_TYPE_ORDERED_LIST:
			ol, err := CreateHTMLOrderedList(block)
			if err != nil {
				return "", err
			}

			html += ol
		case BLOCK_TYPE_TABLE:
			table, err := CreateHTMLTable(block)
			if err != nil {
				return "", err
			}

			html += table
		case BLOCK_TYPE_HR:
			hr, err := CreateHTMLHr(block)
			if err != nil {
				return "", err
			}

			html += hr
		default:
			return "", errors.New("Block type not recognized")
		}
	}

	html += "</div>"

//  fmt.Println(html)

	return html, nil
}
