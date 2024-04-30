package htmlnode

import (
	"testing"
)

func TestPropToHTML(t *testing.T) {
	props := map[string]string{
		"href": "https://www.google.com",
	}

	node := &HTMLNode{tag: "p", value: "This is a paragraph", children: []Node{}, props: props}
	attr := node.propsToHTML()
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

	node := &HTMLNode{tag: "p", value: "This is a paragraph", children: []Node{}, props: props}
	attr := node.propsToHTML()
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

	node := &HTMLNode{tag: "h1", value: "This is a paragraph", children: []Node{}, props: props}
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
	node := HTMLNode{tag: "h1", value: "This is a paragraph", children: []Node{}, props: props}

	expected := "HTMLNode(h1, This is a paragraph, [], map[href:https://www.google.com])"

	if node.toString() == expected {
		return
	}

	t.Errorf("String values don't match. Expcted: %s, Got: %s", expected, node.toString())
}

func TestLeafNode(t *testing.T) {
	node := &LeafNode{HTMLNode{tag: "a", value: "Just text, fam"}}
	expected := "<a>Just text, fam</a>"

	value, err := node.toHTML()
	if err != nil {
		t.Errorf("String values don't match. Expcted: %s, Got: %s", expected, node.toString())
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

	node := &LeafNode{HTMLNode{tag: "a", value: "Just text, fam", props: props}}
	expected := "<a class='anchor' href='https://www.google.com' id='random'>Just text, fam</a>"

	value, err := node.toHTML()
	if err != nil {
		t.Errorf("String values don't match. Expcted: %s, Got: %s", expected, value)
	}

	if value == expected {
		return
	}

	t.Errorf("Leaf node value does not match. Expected: %s, Got: %s", expected, value)
}

func TestLeafNodeWithoutTag(t *testing.T) {
	node := &LeafNode{HTMLNode{value: "Just text, fam"}}
	expected := "Just text, fam"

	value, err := node.toHTML()
	if err != nil {
		t.Errorf("String values don't match. Expcted: %s, Got: %s", expected, value)
	}

	if value == expected {
		return
	}

	t.Errorf("Leaf node value does not match. Expected: %s, Got: %s", expected, value)
}

func TestLeafNodeWithoutValue(t *testing.T) {
	node := &LeafNode{HTMLNode{tag: "a"}}
	expected := "Leaf node does not have value"

	_, err := node.toHTML()
	if err.Error() == expected {
		return
	}

	t.Errorf("Leaf node without props does not give expected error. Expected: %s, Got: %s", expected, err.Error())
}

func TestParentNode(t *testing.T) {
	props := map[string]string{
		"href":  "https://www.google.com",
		"id":    "random",
		"class": "anchor",
	}

	leaf1 := &LeafNode{HTMLNode{value: "Just text, fam"}}
	leaf2 := &LeafNode{HTMLNode{tag: "a", value: "Just text, fam", props: props}}

	node := &ParentNode{HTMLNode{tag: "div", children: []Node{leaf1, leaf2}}}
	expected := "<div>Just text, fam<a class='anchor' href='https://www.google.com' id='random'>Just text, fam</a></div>"

	value, err := node.toHTML()
	if err != nil {
		t.Errorf("String values don't match. Expected: %s, Got: %s", expected, node.toString())
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

	leaf1 := &LeafNode{HTMLNode{value: "Just text, fam"}}
	node := &ParentNode{HTMLNode{tag: "a", children: []Node{leaf1}, props: props}}
	expected := "<a class='anchor' href='https://www.google.com' id='random'>Just text, fam</a>"

	value, err := node.toHTML()
	if err != nil {
		t.Errorf("String values don't match. Expcted: %s, Got: %s", expected, node.toString())
	}

	if value == expected {
		return
	}

	t.Errorf("Leaf node value does not match. Expected: %s, Got: %s", expected, value)
}

func TestParentNodeWithoutTag(t *testing.T) {
	leaf1 := &LeafNode{HTMLNode{value: "Just text, fam"}}
	node := &ParentNode{HTMLNode{children: []Node{leaf1}}}
	expected := "Parent node does not have a tag"

	_, err := node.toHTML()
	if err.Error() == expected {
    return
	}

	t.Errorf("Parent node without tag does not give expected error. Expected: %s, Got: %s", expected, err.Error())
}


func TestNestedParentNode(t *testing.T) {
  leafh1 := &LeafNode{HTMLNode{tag: "h1", value: "Heading 1"}}
  leafh2 := &LeafNode{HTMLNode{tag: "h2", value: "Heading 2"}}
  p1 := &LeafNode{HTMLNode{tag: "p", value: "Paragraph 1"}}
  divP := &LeafNode{HTMLNode{tag: "p", value: "Complex content", props: map[string]string{"class": "new", "id": "unique"}}}
  articleDiv := &ParentNode{HTMLNode{tag: "div", children: []Node{divP}}}
  boldLi := &LeafNode{HTMLNode{tag: "b", value: "Bold text"}}
  italicLi := &LeafNode{HTMLNode{tag: "i", value: "Italic text"}}
  codeLi := &LeafNode{HTMLNode{tag: "code", value: "print('Hello, World!')"}}
  li1 := &ParentNode{HTMLNode{tag: "li", children: []Node{boldLi, italicLi}}}
  li2 := &ParentNode{HTMLNode{tag: "li", children: []Node{codeLi}}}
  ul := &ParentNode{HTMLNode{tag: "ul", children: []Node{li1, li2}}}
  article := &ParentNode{HTMLNode{tag: "article", children: []Node{p1, ul, articleDiv}}}
  section := &ParentNode{HTMLNode{tag: "section", children: []Node{leafh2, article}}}
  footerP := &LeafNode{HTMLNode{tag: "p", value: "Footer text"}}
  footer := &ParentNode{HTMLNode{tag: "footer", children: []Node{footerP}}}
  node := &ParentNode{HTMLNode{tag: "div", children: []Node{leafh1, section, footer}}}

  expected := "<div><h1>Heading 1</h1><section><h2>Heading 2</h2><article><p>Paragraph 1</p><ul><li><b>Bold text</b><i>Italic text</i></li><li><code>print('Hello, World!')</code></li></ul><div><p class='new' id='unique'>Complex content</p></div></article></section><footer><p>Footer text</p></footer></div>"

  value, err := node.toHTML()
	if err != nil {
		t.Errorf("String values don't match. Expcted: %s, Got: %s", expected, value)
	}

  if value == expected {
    return
  }

	t.Errorf("Parent node value does not match. Expected: %s, Got: %s", expected, value)
}
