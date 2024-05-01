package block

import "strings"

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
