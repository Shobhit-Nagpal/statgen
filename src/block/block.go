package block

import (
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
  headingRe := regexp.MustCompile(`^##?#?#?#?#?\s.*$`)
  quoteRe := regexp.MustCompile(`(?m:^>.*?$)`)
  ulRe := regexp.MustCompile(`(?m:^[*-]\s.*?$)`)
  olRe := regexp.MustCompile(`(?m:^(1|[2-9]\d*)\.\s.*?$)`)
  codeRe := regexp.MustCompile("^```[\\s\\S]*?```$")

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
  } else {
    return BLOCK_TYPE_PARAGRAPH
  }
}
