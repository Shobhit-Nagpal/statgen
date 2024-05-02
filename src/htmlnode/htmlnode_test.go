package htmlnode

import (
	"fmt"
	"testing"
)

func TestPropToHTML(t *testing.T) {
	props := map[string]string{
		"href": "https://www.google.com",
	}

	node := &HTMLNode{Tag: "p", Value: "This is a paragraph", Children: []Node{}, Props: props}
	attr := node.PropsToHTML()
	expected := "href='https://www.google.com'"

	if attr == expected {
		return
	}

	t.Errorf("Props not converted correctly. Expected: %s, got: %s", expected, attr)
}

func TestPropsToHTML(t *testing.T) {
	props := map[string]string{
		"href":   "https://www.google.com",
		"target": "_blank",
		"class":  "para",
	}

	node := &HTMLNode{Tag: "p", Value: "This is a paragraph", Children: []Node{}, Props: props}
	attr := node.PropsToHTML()
	expected := "class='para' href='https://www.google.com' target='_blank'"

	if attr == expected {
		return
	}

	t.Errorf("Props not converted correctly. Expected: %s, got: %s", expected, attr)
}

func TestPropsToHTMLHeading(t *testing.T) {
	props := map[string]string{
		"href":   "https://www.google.com",
		"target": "_blank",
		"class":  "heading another-one",
	}

	node := &HTMLNode{Tag: "h1", Value: "This is a paragraph", Children: []Node{}, Props: props}
	attr := node.PropsToHTML()
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
	node := HTMLNode{Tag: "h1", Value: "This is a paragraph", Children: []Node{}, Props: props}

	expected := "HTMLNode(h1, This is a paragraph, [], map[href:https://www.google.com])"

	if node.ToString() == expected {
		return
	}

	t.Errorf("String values don't match. Expcted: %s, Got: %s", expected, node.ToString())
}

func TestLeafNode(t *testing.T) {
	node := &LeafNode{HTMLNode{Tag: "a", Value: "Just text, fam"}}
	expected := "<a>Just text, fam</a>"

	value, err := node.ToHTML()
	if err != nil {
		t.Errorf("String values don't match. Expcted: %s, Got: %s", expected, node.ToString())
	}

	if value == expected {
		return
	}

	t.Errorf("Leaf node value does not match. Expected: %s, Got: %s", expected, value)
}

func TestLeafNodeWithProps(t *testing.T) {
	props := map[string]string{
		"href":  "https://www.google.com",
		"id":    "random",
		"class": "anchor",
	}

	node := &LeafNode{HTMLNode{Tag: "a", Value: "Just text, fam", Props: props}}
	expected := "<a class='anchor' href='https://www.google.com' id='random'>Just text, fam</a>"

	value, err := node.ToHTML()
	if err != nil {
		t.Errorf("String values don't match. Expcted: %s, Got: %s", expected, value)
	}

	if value == expected {
		return
	}

	t.Errorf("Leaf node value does not match. Expected: %s, Got: %s", expected, value)
}

func TestLeafNodeWithoutTag(t *testing.T) {
	node := &LeafNode{HTMLNode{Value: "Just text, fam"}}
	expected := "Just text, fam"

	value, err := node.ToHTML()
	if err != nil {
		t.Errorf("String values don't match. Expcted: %s, Got: %s", expected, value)
	}

	if value == expected {
		return
	}

	t.Errorf("Leaf node value does not match. Expected: %s, Got: %s", expected, value)
}

func TestLeafNodeWithoutValue(t *testing.T) {
	node := &LeafNode{HTMLNode{Tag: "a"}}
  expected := fmt.Sprintf("Leaf node %s does not have value. %s", node.Tag, node.ToString())

	_, err := node.ToHTML()
	if err.Error() == expected {
		return
	}

	t.Errorf("Leaf node without Value does not give expected error. Expected: %s, Got: %s", expected, err.Error())
}

func TestParentNode(t *testing.T) {
	props := map[string]string{
		"href":  "https://www.google.com",
		"id":    "random",
		"class": "anchor",
	}

	leaf1 := &LeafNode{HTMLNode{Value: "Just text, fam"}}
	leaf2 := &LeafNode{HTMLNode{Tag: "a", Value: "Just text, fam", Props: props}}

	node := &ParentNode{HTMLNode{Tag: "div", Children: []Node{leaf1, leaf2}}}
	expected := "<div>Just text, fam<a class='anchor' href='https://www.google.com' id='random'>Just text, fam</a></div>"

	value, err := node.ToHTML()
	if err != nil {
		t.Errorf("String values don't match. Expected: %s, Got: %s", expected, node.ToString())
	}

	if value == expected {
		return
	}

	t.Errorf("Leaf node value does not match. Expected: %s, Got: %s", expected, value)
}

func TestParentNodeWithProps(t *testing.T) {
	props := map[string]string{
		"href":  "https://www.google.com",
		"id":    "random",
		"class": "anchor",
	}

	leaf1 := &LeafNode{HTMLNode{Value: "Just text, fam"}}
	node := &ParentNode{HTMLNode{Tag: "a", Children: []Node{leaf1}, Props: props}}
	expected := "<a class='anchor' href='https://www.google.com' id='random'>Just text, fam</a>"

	value, err := node.ToHTML()
	if err != nil {
		t.Errorf("String values don't match. Expcted: %s, Got: %s", expected, node.ToString())
	}

	if value == expected {
		return
	}

	t.Errorf("Leaf node value does not match. Expected: %s, Got: %s", expected, value)
}

func TestParentNodeWithoutTag(t *testing.T) {
	leaf1 := &LeafNode{HTMLNode{Value: "Just text, fam"}}
	node := &ParentNode{HTMLNode{Children: []Node{leaf1}}}
	expected := "Parent node does not have a Tag"

	_, err := node.ToHTML()
	if err.Error() == expected {
    return
	}

	t.Errorf("Parent node without Tag does not give expected error. Expected: %s, Got: %s", expected, err.Error())
}


func TestNestedParentNode(t *testing.T) {
  leafh1 := &LeafNode{HTMLNode{Tag: "h1", Value: "Heading 1"}}
  leafh2 := &LeafNode{HTMLNode{Tag: "h2", Value: "Heading 2"}}
  p1 := &LeafNode{HTMLNode{Tag: "p", Value: "Paragraph 1"}}
  divP := &LeafNode{HTMLNode{Tag: "p", Value: "Complex content", Props: map[string]string{"class": "new", "id": "unique"}}}
  articleDiv := &ParentNode{HTMLNode{Tag: "div", Children: []Node{divP}}}
  boldLi := &LeafNode{HTMLNode{Tag: "b", Value: "Bold text"}}
  italicLi := &LeafNode{HTMLNode{Tag: "i", Value: "Italic text"}}
  codeLi := &LeafNode{HTMLNode{Tag: "code", Value: "print('Hello, World!')"}}
  li1 := &ParentNode{HTMLNode{Tag: "li", Children: []Node{boldLi, italicLi}}}
  li2 := &ParentNode{HTMLNode{Tag: "li", Children: []Node{codeLi}}}
  ul := &ParentNode{HTMLNode{Tag: "ul", Children: []Node{li1, li2}}}
  article := &ParentNode{HTMLNode{Tag: "article", Children: []Node{p1, ul, articleDiv}}}
  section := &ParentNode{HTMLNode{Tag: "section", Children: []Node{leafh2, article}}}
  footerP := &LeafNode{HTMLNode{Tag: "p", Value: "Footer text"}}
  footer := &ParentNode{HTMLNode{Tag: "footer", Children: []Node{footerP}}}
  node := &ParentNode{HTMLNode{Tag: "div", Children: []Node{leafh1, section, footer}}}

  expected := "<div><h1>Heading 1</h1><section><h2>Heading 2</h2><article><p>Paragraph 1</p><ul><li><b>Bold text</b><i>Italic text</i></li><li><code>print('Hello, World!')</code></li></ul><div><p class='new' id='unique'>Complex content</p></div></article></section><footer><p>Footer text</p></footer></div>"

  value, err := node.ToHTML()
	if err != nil {
		t.Errorf("String values don't match. Expcted: %s, Got: %s", expected, value)
	}

  if value == expected {
    return
  }

	t.Errorf("Parent node value does not match. Expected: %s, Got: %s", expected, value)
}
