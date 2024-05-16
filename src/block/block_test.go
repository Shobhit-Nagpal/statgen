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

func TestBlockToBlockTypeHeading(t *testing.T) {
  text := "## Heading 2"
  blockType := BlockToBlockType(text)

  if blockType == BLOCK_TYPE_HEADING {
    return
  }

  t.Errorf("Block type not same. Expected: %s, Got: %s", BLOCK_TYPE_HEADING, blockType)
}

func TestBlockToBlockTypeParagraph(t *testing.T) {
  text := "Classic paragraph"
  blockType := BlockToBlockType(text)

  if blockType == BLOCK_TYPE_PARAGRAPH {
    return
  }

  t.Errorf("Block type not same. Expected: %s, Got: %s", BLOCK_TYPE_PARAGRAPH, blockType)
}

func TestBlockToBlockTypeQuote(t *testing.T) {
  text := ">This is a whole quote\n> - Socrates probably\n> Who knows ackshually"

  blockType := BlockToBlockType(text)

  if blockType == BLOCK_TYPE_QUOTE {
    return
  }

  t.Errorf("Block type not same. Expected: %s, Got: %s", BLOCK_TYPE_QUOTE, blockType)
}

func TestBlockToBlockTypeUnorderedList(t *testing.T) {
  text := "- Brocolli\n- Onion"

  blockType := BlockToBlockType(text)

  if blockType == BLOCK_TYPE_UNORDERED_LIST {
    return
  }

  t.Errorf("Block type not same. Expected: %s, Got: %s", BLOCK_TYPE_UNORDERED_LIST, blockType)
}

func TestBlockToBlockTypeOrderedList(t *testing.T) {
  text := "1. Brocolli\n2. Onion"

  blockType := BlockToBlockType(text)

  if blockType == BLOCK_TYPE_ORDERED_LIST {
    return
  }

  t.Errorf("Block type not same. Expected: %s, Got: %s", BLOCK_TYPE_ORDERED_LIST, blockType)
}

func TestBlockToBlockTypeCode(t *testing.T) {
  text := "```print('Checkmate')\nprint('Gottem')```"

  blockType := BlockToBlockType(text)

  if blockType == BLOCK_TYPE_CODE {
    return
  }

  t.Errorf("Block type not same. Expected: %s, Got: %s", BLOCK_TYPE_CODE, blockType)
}

func TestBlockToBlockTypeHr(t *testing.T) {
  text := "---"

  blockType := BlockToBlockType(text)

  if blockType == BLOCK_TYPE_HR {
    return
  }

  t.Errorf("Block type not same. Expected: %s, Got: %s", BLOCK_TYPE_CODE, blockType)
}
