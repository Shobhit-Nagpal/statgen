package htmlnode

import (
	"testing"
)

func TestPropToHTML(t *testing.T) {
  props := map[string]string{
    "href": "https://www.google.com",
  }

  node := HTMLNode{tag: "p", value: "This is a paragraph", children: []string{}, props: props}
  attr := node.propsToHTML()
  expected := "href='https://www.google.com'"
  
  if attr == expected {
    return
  }

  t.Errorf("Props not converted correctly. Expected: %s, got: %s", expected, attr)
}

func TestPropsToHTML(t *testing.T) {
  props := map[string]string{
    "href": "https://www.google.com",
    "target": "_blank",
    "class": "para",
  }

  node := HTMLNode{tag: "p", value: "This is a paragraph", children: []string{}, props: props}
  attr := node.propsToHTML()
  expected := "class='para' href='https://www.google.com' target='_blank'"
  
  if attr == expected {
    return
  }

  t.Errorf("Props not converted correctly. Expected: %s, got: %s", expected, attr)
}

func TestPropsToHTMLHeading(t *testing.T) {
  props := map[string]string{
    "href": "https://www.google.com",
    "target": "_blank",
    "class": "heading another-one",
  }


  node := HTMLNode{tag: "h1", value: "This is a paragraph", children: []string{}, props: props}
  attr := node.propsToHTML()
  expected := "class='heading another-one' href='https://www.google.com' target='_blank'"
  
  if attr == expected {
    return
  }

  t.Errorf("Props not converted correctly. Expected: %s, got: %s", expected, attr)
}

func TestToString(t *testing.T) {
  props := map[string]string{
    "href": "https://www.google.com",
  }
  node := HTMLNode{tag: "h1", value: "This is a paragraph", children: []string{}, props: props}

  expected := "HTMLNode(h1, This is a paragraph, [], map[href:https://www.google.com])"

  if node.toString() == expected {
    return
  }

  t.Errorf("String values don't match. Expcted: %s, Got: %s", expected, node.toString())
}
