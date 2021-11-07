
package gquery

import (
	strings "strings"
)

type Document struct{
	Node0
	doctype *DocType
	content *HtmlNode
}

func NewDocumentNode()(p *Document){
	p = &Document{}
	p.Node0 = NewNode0(p)
	return
}

func (p *Document)Name()(string){
	return "#document"
}

func (*Document)HasAttr(k string)(bool){
	return false
}

func (*Document)GetAttrDefault(k string, d string)(v string){
	panic("Document node don't have any attribute")
}

func (*Document)GetAttr(k string)(v string){
	panic("Document node don't have any attribute")
}

func (*Document)SetAttr(k string, v string){
	panic("Document node don't have any attribute")
}

func (*Document)DelAttr(k string)(bool){
	panic("Document node don't have any attribute")
}

func (*Document)Attrs()(map[string]string){
	panic("Document node don't have any attribute")
}

func (n *Document)GetText()(string){
	return n.content.GetText()
}

func (n *Document)SetText(t string){
	n.content.SetText(t)
}

func (n *Document)setParent(p Node){
	panic("Document node must be the root parent")
}

func (n *Document)GetDocType()(*DocType){
	return n.doctype
}

func (n *Document)SetDocType(t *DocType){
	if n.doctype != nil {
		n.doctype.setParent(nil)
	}
	n.doctype = t
	if t != nil {
		t.setParent(n)
		if n.content != nil {
			t.setAfter(n.content)
		}
	}
}

func (n *Document)GetHtmlNode()(*HtmlNode){
	return n.content
}

func (n *Document)SetHtmlNode(h *HtmlNode){
	if n.content != nil {
		n.content.setParent(nil)
	}
	n.content = h
	if h != nil {
		h.setParent(n)
		if n.doctype != nil {
			h.setBefore(n.doctype)
		}
	}
}

func (n *Document)GetNodeList()(*NodeList){
	panic("Document node cannot operate children")
}

func (n *Document)AppendChild(o Node){
	panic("Document node cannot operate children")
}

func (n *Document)RemoveChild(o Node){
	panic("Document node cannot operate children")
}

func (n *Document)InsertChild(o Node, _t ...Node){
	panic("Document node cannot operate children")
}

func (n *Document)Find(pattern string)([]Node){
	return n.content.Find(pattern)
}

func (p *Document)String()(str string){
	str = ""
	if p.doctype != nil {
		str += p.doctype.String()
	}
	if p.content != nil {
		str += p.content.String()
	}
	return
}

type DocType struct{
	Node0
	typ string
	extras []string
}

func NewDocTypeNode(typ string, extras ...string)(n *DocType){
	n = &DocType{
		typ: strings.ToLower(typ),
		extras: extras,
	}
	n.Node0 = NewNode0(n)
	return
}

func (*DocType)IsSimple()(bool){
	return true
}

func (n *DocType)Name()(string){
	return "#doctype"
}

func (n *DocType)GetText()(string){
	return ""
}

func (n *DocType)SetText(t string){
}

func (n *DocType)GetValue()(string){
	return n.typ
}

func (n *DocType)SetValue(t string){
	n.typ = t
}

func (n *DocType)setParent(p Node){
	if p != nil {
		if _, ok := p.(*Document); !ok {
			panic("DocType Node's parent must be Document")
		}
	}
	n.Node0.setParent(p)
}

func (n *DocType)Extras()(*[]string){
	return &(n.extras)
}

func (n *DocType)String()(str string){
	str = "<!DOCTYPE " + n.typ
	for _, e := range n.extras {
		str += " " + e
	}
	str += ">"
	return
}

type HtmlNode struct{
	ParentNode0
}

func NewHtmlNode()(p *HtmlNode){
	p = &HtmlNode{}
	p.ParentNode0 = NewParentNode0(p)
	return
}

func (p *HtmlNode)Name()(string){
	return "html"
}

func (n *HtmlNode)setParent(p Node){
	if p != nil {
		if _, ok := p.(*Document); !ok {
			panic("HtmlNode's parent must be Document")
		}
	}
	n.ParentNode0.setParent(p)
}

func (p *HtmlNode)String()(str string){
	return "<html" + p.AttrString() + ">" + p.ContentString() + "</html>"
}

type HeadNode struct{
	ParentNode0
}

func NewHeadNode()(p *HeadNode){
	p = &HeadNode{}
	p.ParentNode0 = NewParentNode0(p)
	return
}

func (p *HeadNode)Name()(string){
	return "head"
}

func (n *HeadNode)setParent(p Node){
	if p != nil {
		if _, ok := p.(*HtmlNode); !ok {
			panic("HeadNode's parent must be html")
		}
	}
	n.ParentNode0.setParent(p)
}

func (n *HeadNode)setText(t string){
}

func (p *HeadNode)GetText()(string){
	return ""
}

func (p *HeadNode)String()(str string){
	return "<head" + p.AttrString() + ">" + p.ContentString() + "</head>"
}

type BodyNode struct{
	ParentNode0
}

func NewBodyNode()(p *BodyNode){
	p = &BodyNode{}
	p.ParentNode0 = NewParentNode0(p)
	return
}

func (p *BodyNode)Name()(string){
	return "body"
}

func (n *BodyNode)setParent(p Node){
	if p != nil {
		if _, ok := p.(*HtmlNode); !ok {
			panic("BodyNode's parent must be html")
		}
	}
	n.ParentNode0.setParent(p)
}

func (p *BodyNode)String()(str string){
	return "<body" + p.AttrString() + ">" + p.ContentString() + "</body>"
}

var (
	_ Node = (*Document)(nil)
	_ Node = (*DocType)(nil)
	_ Node = (*HtmlNode)(nil)
	_ Node = (*HeadNode)(nil)
	_ Node = (*BodyNode)(nil)
)
