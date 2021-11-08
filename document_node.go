
package gquery

import (
	io "io"
	strings "strings"
)

type Document struct{
	ParentNode0
}

func NewDocumentNode()(p *Document){
	p = &Document{}
	p.ParentNode0 = NewParentNode0(p)
	return
}

func (*Document)Name()(string){
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
	return n.GetHtmlNode().GetText()
}

func (n *Document)SetText(t string){
	n.GetHtmlNode().SetText(t)
}

func (n *Document)setParent(p Node){
	panic("Document node must be the root parent")
}

func (n *Document)GetDocType()(*DocType){
	if dt, ok := n.children.First().(*DocType); ok {
		return dt
	}
	return nil
}

func (n *Document)SetDocType(t *DocType){
	if dt := n.GetDocType(); dt != nil {
		n.RemoveChild(dt)
	}
	if t != nil {
		n.InsertChild(t, nil)
	}
}

func (n *Document)GetHtmlNode()(*HtmlNode){
	if ht, ok := n.children.Last().(*HtmlNode); ok {
		return ht
	}
	return nil
}

func (n *Document)SetHtmlNode(h *HtmlNode){
	if ht := n.GetHtmlNode(); ht != nil {
		n.RemoveChild(ht)
	}
	if h != nil {
		n.AppendChild(h)
	}
}

func (n *Document)WriteTo(w io.Writer)(written int64, err error){
	return n.ContentWriteTo(w)
}

func (n *Document)String()(str string){
	return n.ContentString()
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
			panic("DocType's parent must be Document")
		}
	}
	n.Node0.setParent(p)
}

func (n *DocType)setBefore(p Node){
	if p != nil {
		panic("DocType must be first of NodeList")
	}
	n.Node0.setParent(p)
}

func (n *DocType)Extras()(*[]string){
	return &(n.extras)
}

func (n *DocType)WriteTo(w io.Writer)(written int64, err error){
	var n0 int
	n0, err = w.Write(([]byte)(n.String()))
	written = (int64)(n0)
	if err != nil { return }
	return
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

func (n *HtmlNode)setAfter(p Node){
	if p != nil {
		panic("HtmlNode must be last of NodeList")
	}
	n.ParentNode0.setParent(p)
}

func (n *HtmlNode)WriteTo(w io.Writer)(written int64, err error){
	var (
		n0 int
		n1 int64
	)
	written = 0
	n0, err = w.Write(([]byte)("<html"))
	written += (int64)(n0)
	if err != nil { return }
	n0, err = w.Write(([]byte)(n.AttrString()))
	written += (int64)(n0)
	if err != nil { return }
	n0, err = w.Write(([]byte)(">"))
	written += (int64)(n0)
	if err != nil { return }
	n1, err = n.ContentWriteTo(w)
	written += n1
	if err != nil { return }
	n0, err = w.Write(([]byte)("</html>"))
	written += (int64)(n0)
	if err != nil { return }
	return
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

func (n *HeadNode)WriteTo(w io.Writer)(written int64, err error){
	var (
		n0 int
		n1 int64
	)
	written = 0
	n0, err = w.Write(([]byte)("<head"))
	written += (int64)(n0)
	if err != nil { return }
	n0, err = w.Write(([]byte)(n.AttrString()))
	written += (int64)(n0)
	if err != nil { return }
	n0, err = w.Write(([]byte)(">"))
	written += (int64)(n0)
	if err != nil { return }
	n1, err = n.ContentWriteTo(w)
	written += n1
	if err != nil { return }
	n0, err = w.Write(([]byte)("</head>"))
	written += (int64)(n0)
	if err != nil { return }
	return
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

func (n *BodyNode)WriteTo(w io.Writer)(written int64, err error){
	var (
		n0 int
		n1 int64
	)
	written = 0
	n0, err = w.Write(([]byte)("<body"))
	written += (int64)(n0)
	if err != nil { return }
	n0, err = w.Write(([]byte)(n.AttrString()))
	written += (int64)(n0)
	if err != nil { return }
	n0, err = w.Write(([]byte)(">"))
	written += (int64)(n0)
	if err != nil { return }
	n1, err = n.ContentWriteTo(w)
	written += n1
	if err != nil { return }
	n0, err = w.Write(([]byte)("</body>"))
	written += (int64)(n0)
	if err != nil { return }
	return
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
