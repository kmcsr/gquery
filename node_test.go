
package gquery_test

import (
	testing "testing"
	. "github.com/kmcsr/gquery"
)

func TestTextNode(t *testing.T){
	tt := "test text right"
	t.Logf(`text := NewTextNode("%s")`, tt)
	text := NewTextNode(tt)
	if text.Name() != "#text" {
		t.Fatalf(`text.Name() == "%s", want "%s"`, text.Name(), "#text")
	}
	if !text.IsSimple() {
		t.Fatalf("Text node must be simple")
	}
	if text.String() != tt {
		t.Errorf(`text.String() == "%s", want "%s"`, text.String(), tt)
	}
	if text.GetText() != tt {
		t.Errorf(`text.GetText() == "%s", want "%s"`, text.GetText(), tt)
	}
	if text.GetValue() != tt {
		t.Errorf(`text.GetValue() == "%s", want "%s"`, text.GetValue(), tt)
	}
	tt2 := "hhh aaa mmm www"
	text.SetValue(tt2)
	if text.String() != tt2 {
		t.Errorf(`text.String() == "%s", want "%s"`, text.String(), tt2)
	}
	if text.GetText() != tt2 {
		t.Errorf(`text.GetText() == "%s", want "%s"`, text.GetText(), tt2)
	}
	if text.GetValue() != tt2 {
		t.Errorf(`text.GetValue() == "%s", want "%s"`, text.GetValue(), tt2)
	}
}

func TestCommentNode(t *testing.T){
	tt := "test comment right"
	t.Logf(`common := NewCommentNode("%s")`, tt)
	common := NewCommentNode(tt)
	if common.Name() != "#comment" {
		t.Fatalf(`common.Name() == "%s", want "%s"`, common.Name(), "#comment")
	}
	if common.String() != "<!--" + tt + "-->" {
		t.Errorf(`common.String() == "%s", want %s`, common.String(), "<!--" + tt + "-->")
	}
	if common.GetValue() != tt {
		t.Errorf(`common.GetValue() == "%s", want %s`, common.GetValue(), tt)
	}
	tt2 := "hhh aaa mmm www"
	common.SetValue(tt2)
	if common.String() != "<!--" + tt2 + "-->" {
		t.Errorf(`common.String() == "%s", want %s`, common.String(), "<!--" + tt2 + "-->")
	}
	if common.GetValue() != tt2 {
		t.Errorf(`common.GetValue() == "%s", want %s`, common.GetValue(), tt2)
	}
}

func TestSpanNode(t *testing.T){
	sp1 := NewSpanNode("SpAn")
	if sp1.Name() != "span" {
		t.Fatalf(`sp1.Name() == "%s", want "%s"`, sp1.Name(), "span")
	}
	if sp1.IsSimple() {
		t.Fatalf("Span node must not be simple")
	}
	if sp1.String() != "<span></span>" {
		t.Errorf(`sp1.String() != <span></span>, has %s`, sp1.String())
	}
	sp1.SetAttr("id", "test-span")
	if sp1.String() != `<span id="test-span"></span>` {
		t.Errorf(`sp1.String() != <span id="test-span"></span>, has %s`, sp1.String())
	}
	sp1.SetAttr("class", "test-cls aabbcc")
	if sp1.String() != `<span id="test-span" class="test-cls aabbcc"></span>`{
		t.Errorf(`sp1.String() != <span id="test-span" class="test-cls aabbcc"></span>, has %s`, sp1.String())
	}
}

func TestSimpleSpanNode(t *testing.T){
	sp1 := NewSimpleSpanNode("inPUt")
	if sp1.Name() != "input" {
		t.Fatalf(`sp1.Name() == "%s", want "%s"`, sp1.Name(), "input")
	}
	if !sp1.IsSimple() {
		t.Fatalf("Simple span node must be simple")
	}
	if sp1.String() != "<input/>" {
		t.Errorf(`sp1.String() != <input/>, has %s`, sp1.String())
	}
	sp1.SetAttr("id", "test-input")
	if sp1.String() != `<input id="test-input"/>` {
		t.Errorf(`sp1.String() != <input id="test-input"/>, has %s`, sp1.String())
	}
	sp1.SetAttr("class", "test-cls aabbcc")
	tt := `<input id="test-input" class="test-cls aabbcc"/>`
	ss := sp1.String()
	if ss != tt{
		t.Logf("tlen: %d, sslen: %d", len(tt), len(ss))
		t.Errorf(`sp1.String() != %s, has %s`, tt, ss)
	}
}

func TestDocumentNode(t *testing.T){
	doc := NewDocumentNode()
	html := NewHtmlNode()
	doc.SetHtmlNode(html)
	t.Log("document:", doc.String())
	head := NewHeadNode()
	html.AppendChild(head)
	t.Log("document:", doc.String())
	body := NewBodyNode()
	html.AppendChild(body)
	t.Log("document:", doc.String())
	div1 := NewSpanNode("div")
	body.AppendChild(div1)
	div1.SetAttr("id", "test1")
	t.Log("document:", doc.String())
	div2 := NewSpanNode("div")
	div2.SetAttr("id", "test2")
	body.AppendChild(div2)
	div2.SetText("hello&copy; abc")
	t.Log("document:", doc.String())
	t.Log("text:", doc.GetText())
	div1.SetText("  emm div1 ")
	t.Log("document:", doc.String())
	t.Log("text:", doc.GetText())
	txt := NewTextNode()
	txt.SetValue(" That's copy sym: &copy")
	body.AppendChild(txt)
	t.Log("document:", doc.String())
	t.Log("text:", doc.GetText())
}
