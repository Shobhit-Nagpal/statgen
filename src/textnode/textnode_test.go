package textnode

import "testing"

func TestIsEqual(t *testing.T) {
  textnode1 := &TextNode{Text: "This is a text", TextType: "bold"}
  textnode2 := &TextNode{Text: "This is a text", TextType: "bold"}

  equal := textnode1.IsEqual(textnode2)

  if equal {
    return
  }

  t.Errorf("Text nodes are not equal. Textnode1: %s, Textnode2: %s", textnode1.ToString(), textnode2.ToString())
}

func TestIsNotEqual(t *testing.T) {
  textnode1 := &TextNode{Text: "This is a text", TextType: "bold"}
  textnode2 := &TextNode{Text: "This is another text", TextType: "bold"}

  equal := textnode1.IsEqual(textnode2)

  if !equal {
    return
  }

  t.Errorf("Text nodes are equal. Textnode1: %s, Textnode2: %s", textnode1.ToString(), textnode2.ToString())
}


func TestIsEqualWithUrl(t *testing.T) {
  textnode1 := &TextNode{Text: "This is a text", TextType: "bold", Url: "https://www.shobhitnagpal.com"}
  textnode2 := &TextNode{Text: "This is a text", TextType: "bold", Url: "https://www.shobhitnagpal.com"}

  equal := textnode1.IsEqual(textnode2)

  if equal {
    return
  }

  t.Errorf("Text nodes are not equal. Textnode1: %s, Textnode2: %s", textnode1.ToString(), textnode2.ToString())
}


func TestIsNotEqualWithUrl(t *testing.T) {
  textnode1 := &TextNode{Text: "This is a text", TextType: "bold", Url: "https://www.blogshobhitnagpal.com"}
  textnode2 := &TextNode{Text: "This is a text", TextType: "bold", Url: "https://www.shobhitnagpal.com"}

  equal := textnode1.IsEqual(textnode2)

  if !equal {
    return
  }

  t.Errorf("Text nodes are equal. Textnode1: %s, Textnode2: %s", textnode1.ToString(), textnode2.ToString())
}
