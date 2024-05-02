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

func TestSplitBoldDelimiterWithNoTrails(t *testing.T) {
  tn := &textnode.TextNode{Text: "This is a text with a **bold**", TextType: md.TEXT_TYPE_TEXT}

  newNodes, err := SplitNodesDelimiter([]*textnode.TextNode{tn}, "**", md.TEXT_TYPE_BOLD)
  if err != nil {
    t.Errorf("%s", err.Error())
  }
  
  expectedNodes := []*textnode.TextNode{
    &textnode.TextNode{Text: "This is a text with a ", TextType: md.TEXT_TYPE_TEXT},
    &textnode.TextNode{Text: "bold", TextType: md.TEXT_TYPE_BOLD},
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

func TestExtractMarkdownImages(t *testing.T) {
  text := "This is text with an ![image](https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/zjjcJKZ.png) and ![another](https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/dfsdkjfd.png)"
  images := ExtractMarkdownImages(text)

  expected := []md.MarkdownImage{
    md.MarkdownImage{Text: "image", Url: "https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/zjjcJKZ.png"},
    md.MarkdownImage{Text: "another", Url: "https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/dfsdkjfd.png"},
  }

  for idx, img := range expected {
    if img.IsEqual(images[idx]) {
      continue
    } else {
      t.Errorf("Images are not the same. Expected: %s, Got: %s", img.ToString(), images[idx].ToString())
    }
  }
}


func TestExtractMarkdownLinks(t *testing.T) {
  text := "This is text with an [link](https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/zjjcJKZ.png) and ![another](https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/dfsdkjfd.png)"
  links := ExtractMarkdownLinks(text)

  expected := []md.MarkdownLink{
    md.MarkdownLink{Text: "link", Url: "https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/zjjcJKZ.png"},
    md.MarkdownLink{Text: "another", Url: "https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/dfsdkjfd.png"},
  }

  for idx, link := range expected {
    if link.IsEqual(links[idx]) {
      continue
    } else {
      t.Errorf("Links are not the same. Expected: %s, Got: %s", link.ToString(), links[idx].ToString())
    }
  }

}

func TestSplitNodeImages(t *testing.T) {
  node := &textnode.TextNode{Text: "This is text with an ![image](https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/zjjcJKZ.png) and another ![second image](https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/3elNhQu.png)", TextType: md.TEXT_TYPE_TEXT}

  expected := []*textnode.TextNode{
    &textnode.TextNode{Text: "This is text with an ", TextType: md.TEXT_TYPE_TEXT},
    &textnode.TextNode{Text: "image", TextType: md.TEXT_TYPE_IMAGE, Url: "https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/zjjcJKZ.png"},
    &textnode.TextNode{Text: " and another ", TextType: md.TEXT_TYPE_TEXT},
    &textnode.TextNode{Text: "second image", TextType: md.TEXT_TYPE_IMAGE, Url: "https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/3elNhQu.png"},
  }

  splitNodes, err := SplitNodesImage([]*textnode.TextNode{node})
  if err != nil {
    t.Errorf(err.Error())
  }

  for idx, txtNode := range expected {
    if txtNode.IsEqual(splitNodes[idx]) {
      continue
    } else {
      t.Errorf("Text nodes are not the same. Expected: %s, Got: %s", txtNode.ToString(), splitNodes[idx].ToString())
    }
  }
}

func TestSplitNodeLinks(t *testing.T) {
  node := &textnode.TextNode{Text: "This is text with a [link](https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/zjjcJKZ.png) and another [second link](https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/3elNhQu.png)", TextType: md.TEXT_TYPE_TEXT}

  expected := []*textnode.TextNode{
    &textnode.TextNode{Text: "This is text with a ", TextType: md.TEXT_TYPE_TEXT},
    &textnode.TextNode{Text: "link", TextType: md.TEXT_TYPE_LINK, Url: "https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/zjjcJKZ.png"},
    &textnode.TextNode{Text: " and another ", TextType: md.TEXT_TYPE_TEXT},
    &textnode.TextNode{Text: "second link", TextType: md.TEXT_TYPE_LINK, Url: "https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/3elNhQu.png"},
  }

  splitNodes, err := SplitNodesLink([]*textnode.TextNode{node})
  if err != nil {
    t.Errorf(err.Error())
  }

  for idx, txtNode := range expected {
    if txtNode.IsEqual(splitNodes[idx]) {
      continue
    } else {
      t.Errorf("Text nodes are not the same. Expected: %s, Got: %s", txtNode.ToString(), splitNodes[idx].ToString())
    }
  }
}

func TestTextToTextNodes(t *testing.T) {
  text := "This is **text** with an *italic* word and a `code block` and an ![image](https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/zjjcJKZ.png) and a [link](https://boot.dev)"
  expected := []*textnode.TextNode{
    &textnode.TextNode{Text: "This is ", TextType: md.TEXT_TYPE_TEXT},
    &textnode.TextNode{Text: "text", TextType: md.TEXT_TYPE_BOLD},
    &textnode.TextNode{Text: " with an ", TextType: md.TEXT_TYPE_TEXT},
    &textnode.TextNode{Text: "italic", TextType: md.TEXT_TYPE_ITALIC},
    &textnode.TextNode{Text: " word and a ", TextType: md.TEXT_TYPE_TEXT},
    &textnode.TextNode{Text: "code block", TextType: md.TEXT_TYPE_CODE},
    &textnode.TextNode{Text: " and an ", TextType: md.TEXT_TYPE_TEXT},
    &textnode.TextNode{Text: "image", TextType: md.TEXT_TYPE_IMAGE, Url: "https://storage.googleapis.com/qvault-webapp-dynamic-assets/course_assets/zjjcJKZ.png"},
    &textnode.TextNode{Text: " and a ", TextType: md.TEXT_TYPE_TEXT},
    &textnode.TextNode{Text: "link", TextType: md.TEXT_TYPE_LINK, Url: "https://boot.dev"},
  }

  textNodes, err := TextToTextNodes(text)
  if err != nil {
    t.Errorf(err.Error())
  }

  for idx, node := range expected {
    if node.IsEqual(textNodes[idx]) {
      continue
    } else {
      t.Errorf("Text nodes are not the same. Expected: %s, Got: %s", node.ToString(), textNodes[idx].ToString())
    }
  }
}


func TestTextToTextNodesBoldWordTrail(t *testing.T) {
  text := "A word **Test bold word**"
  expected := []*textnode.TextNode{
    &textnode.TextNode{Text: "A word ", TextType: md.TEXT_TYPE_TEXT},
    &textnode.TextNode{Text: "Test bold word", TextType: md.TEXT_TYPE_BOLD},
 }

  textNodes, err := TextToTextNodes(text)
  if err != nil {
    t.Errorf(err.Error())
  }

  for idx, node := range expected {
    if node.IsEqual(textNodes[idx]) {
      continue
    } else {
      t.Errorf("Text nodes are not the same. Expected: %s, Got: %s", node.ToString(), textNodes[idx].ToString())
    }
  }
}

func TestTextToTextNodesOneItalicWord(t *testing.T) {
  text := "*Test bold word*"
  expected := []*textnode.TextNode{
    &textnode.TextNode{Text: "Test bold word", TextType: md.TEXT_TYPE_ITALIC},
  }

  textNodes, err := TextToTextNodes(text)
  if err != nil {
    t.Errorf(err.Error())
  }

  for idx, node := range expected {
    if node.IsEqual(textNodes[idx]) {
      continue
    } else {
      t.Errorf("Text nodes are not the same. Expected: %s, Got: %s", node.ToString(), textNodes[idx].ToString())
    }
  }
}

func TestTextToTextNodesOneBoldWord(t *testing.T) {
  text := "**Test bold word**"
  expected := []*textnode.TextNode{
    &textnode.TextNode{Text: "Test bold word", TextType: md.TEXT_TYPE_BOLD},
  }

  textNodes, err := TextToTextNodes(text)
  if err != nil {
    t.Errorf(err.Error())
  }

  for idx, node := range expected {
    if node.IsEqual(textNodes[idx]) {
      continue
    } else {
      t.Errorf("Text nodes are not the same. Expected: %s, Got: %s", node.ToString(), textNodes[idx].ToString())
    }
  }
}
