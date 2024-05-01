package block

import "testing"

func TestMarkdownToBlocks(t *testing.T) {
	markdown := `
This is **bolded** paragraph

This is another paragraph with *italic* text here
This is the same paragraph on a new line

* This is a list
* with items
  `

	blocks := MarkdownToBlocks(markdown)
	expected := []string{
		"This is **bolded** paragraph",
		"This is another paragraph with *italic* text here\nThis is the same paragraph on a new line",
		"* This is a list\n* with items",
	}

	for idx, block := range expected {
		if block == blocks[idx] {
			continue
		} else {
			t.Errorf("Block does not match. Expected: %s, Got: %s", block, blocks[idx])
		}
	}
}

func TestMarkdownToBlocksWithExtraNewLines(t *testing.T) {
	markdown := `
This is **bolded** paragraph








This is another paragraph with *italic* text here
This is the same paragraph on a new line










* This is a list
* with items
  `

	blocks := MarkdownToBlocks(markdown)
	expected := []string{
		"This is **bolded** paragraph",
		"This is another paragraph with *italic* text here\nThis is the same paragraph on a new line",
		"* This is a list\n* with items",
	}

	for idx, block := range expected {
		if block == blocks[idx] {
			continue
		} else {
			t.Errorf("Block does not match. Expected: %s, Got: %s", block, blocks[idx])
		}
	}
}
