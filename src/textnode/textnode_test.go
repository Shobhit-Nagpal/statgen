package textnode

import "testing"

func TestIsEqual(t *testing.T) {
  textnode1 := TextNode{text: "This is a text", textType: "bold"}
  textnode2 := TextNode{text: "This is a text", textType: "bold"}

  equal := textnode1.isEqual(textnode2)

  if equal {
    return
  }

  t.Errorf("Text nodes are not equal. Textnode1: %s, Textnode2: %s", textnode1.toString(), textnode2.toString())
}

func TestIsNotEqual(t *testing.T) {
  textnode1 := TextNode{text: "This is a text", textType: "bold"}
  textnode2 := TextNode{text: "This is another text", textType: "bold"}

  equal := textnode1.isEqual(textnode2)

  if !equal {
    return
  }

  t.Errorf("Text nodes are equal. Textnode1: %s, Textnode2: %s", textnode1.toString(), textnode2.toString())
}


func TestIsEqualWithUrl(t *testing.T) {
  textnode1 := TextNode{text: "This is a text", textType: "bold", url: "https://www.shobhitnagpal.com"}
  textnode2 := TextNode{text: "This is a text", textType: "bold", url: "https://www.shobhitnagpal.com"}

  equal := textnode1.isEqual(textnode2)

  if equal {
    return
  }

  t.Errorf("Text nodes are not equal. Textnode1: %s, Textnode2: %s", textnode1.toString(), textnode2.toString())
}


func TestIsNotEqualWithUrl(t *testing.T) {
  textnode1 := TextNode{text: "This is a text", textType: "bold", url: "https://www.blogshobhitnagpal.com"}
  textnode2 := TextNode{text: "This is a text", textType: "bold", url: "https://www.shobhitnagpal.com"}

  equal := textnode1.isEqual(textnode2)

  if !equal {
    return
  }

  t.Errorf("Text nodes are equal. Textnode1: %s, Textnode2: %s", textnode1.toString(), textnode2.toString())
}
