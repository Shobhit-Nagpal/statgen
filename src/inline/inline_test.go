package inline

import (
	"statgen/src/md"
	"statgen/src/textnode"
	"testing"
)

func TestSplitBoldDelimiter(t *testing.T) {
  tn := &textnode.TextNode{Text: "This is a text with a **bold** word", TextType: md.TEXT_TYPE_TEXT}

  newNodes, err := SplitNodesDelimiter([]*textnode.TextNode{tn}, "**", md.TEXT_TYPE_BOLD)
  if err != nil {
    t.Errorf("%s", err.Error())
  }
  
  expectedNodes := []*textnode.TextNode{
    &textnode.TextNode{Text: "This is a text with a ", TextType: md.TEXT_TYPE_TEXT},
    &textnode.TextNode{Text: "bold", TextType: md.TEXT_TYPE_BOLD},
    &textnode.TextNode{Text: " word", TextType: md.TEXT_TYPE_TEXT},
  }

  for idx, node := range expectedNodes {
    if node.IsEqual(newNodes[idx]) {
      continue
    } else {
      t.Errorf("Nodes are not the same. Expected: %s, Got: %s", node.ToString(), newNodes[idx].ToString())
    }
  }
}

func TestSplitItalicDelimiter(t *testing.T) {
  tn := &textnode.TextNode{Text: "This is a text with a *italic* word", TextType: md.TEXT_TYPE_TEXT}

  newNodes, err := SplitNodesDelimiter([]*textnode.TextNode{tn}, "*", md.TEXT_TYPE_ITALIC)
  if err != nil {
    t.Errorf("%s", err.Error())
  }
  
  expectedNodes := []*textnode.TextNode{
    &textnode.TextNode{Text: "This is a text with a ", TextType: md.TEXT_TYPE_TEXT},
    &textnode.TextNode{Text: "italic", TextType: md.TEXT_TYPE_ITALIC},
    &textnode.TextNode{Text: " word", TextType: md.TEXT_TYPE_TEXT},
  }

  for idx, node := range expectedNodes {
    if node.IsEqual(newNodes[idx]) {
      continue
    } else {
      t.Errorf("Nodes are not the same. Expected: %s, Got: %s", node.ToString(), newNodes[idx].ToString())
    }
  }
}

func TestSplitCodeDelimiter(t *testing.T) {
  tn := &textnode.TextNode{Text: "This is a text with a `code` word", TextType: md.TEXT_TYPE_TEXT}

  newNodes, err := SplitNodesDelimiter([]*textnode.TextNode{tn}, "`", md.TEXT_TYPE_CODE)
  if err != nil {
    t.Errorf("%s", err.Error())
  }
  
  expectedNodes := []*textnode.TextNode{
    &textnode.TextNode{Text: "This is a text with a ", TextType: md.TEXT_TYPE_TEXT},
    &textnode.TextNode{Text: "code", TextType: md.TEXT_TYPE_CODE},
    &textnode.TextNode{Text: " word", TextType: md.TEXT_TYPE_TEXT},
  }

  for idx, node := range expectedNodes {
    if node.IsEqual(newNodes[idx]) {
      continue
    } else {
      t.Errorf("Nodes are not the same. Expected: %s, Got: %s", node.ToString(), newNodes[idx].ToString())
    }
  }
}

func TestSplitMultipleBoldDelimiter(t *testing.T) {
  tn := &textnode.TextNode{Text: "This is text with a **bold** word and yet **another bold** word.", TextType: md.TEXT_TYPE_TEXT}

  newNodes, err := SplitNodesDelimiter([]*textnode.TextNode{tn}, "**", md.TEXT_TYPE_BOLD)
  if err != nil {
    t.Errorf("%s", err.Error())
  }
  
  expectedNodes := []*textnode.TextNode{
    &textnode.TextNode{Text: "This is text with a ", TextType: md.TEXT_TYPE_TEXT},
    &textnode.TextNode{Text: "bold", TextType: md.TEXT_TYPE_BOLD},
    &textnode.TextNode{Text: " word and yet ", TextType: md.TEXT_TYPE_TEXT},
    &textnode.TextNode{Text: "another bold", TextType: md.TEXT_TYPE_BOLD},
    &textnode.TextNode{Text: " word.", TextType: md.TEXT_TYPE_TEXT},
  }

  for idx, node := range expectedNodes {
    if node.IsEqual(newNodes[idx]) {
      continue
    } else {
      t.Errorf("Nodes are not the same. Expected: %s, Got: %s", node.ToString(), newNodes[idx].ToString())
    }
  }
}
