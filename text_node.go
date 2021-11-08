
package gquery

import (
	io "io"
	strings "strings"
	html "html"
)

type Text struct{
	Node0
	content string
}

func NewTextNode(texts ...string)(n *Text){
	n = &Text{
		content: html.EscapeString(strings.Join(texts, " ")),
	}
	n.Node0 = NewNode0(n)
	return
}

func (*Text)IsSimple()(bool){
	return true
}

func (*Text)Name()(string){
	return "#text"
}

func (n *Text)GetText()(str string){
	str = standardizedText(html.UnescapeString(n.content))
	if n.After() != nil {
	}
	if len(str) > 0 && str[0] == ' ' {
		if pn := FindPrevNodeExcepts(n, "#common"); pn == nil || IsBlockNodeName(pn.Name()) {
			str = str[1:]
		}
	}
	if len(str) > 0 && str[len(str) - 1] == ' ' {
		if nn := FindNextNodeExcepts(n, "#common"); nn == nil || IsBlockNodeName(nn.Name()) {
			str = str[:len(str) - 1]
		}
	}
	return
}

func (n *Text)SetText(t string){
	n.content = html.EscapeString(t)
}

func (n *Text)GetValue()(string){
	return n.content
}

func (n *Text)SetValue(t string){
	n.content = t
}

func (n *Text)WriteTo(w io.Writer)(written int64, err error){
	var n0 int
	n0, err = w.Write(([]byte)(n.content))
	written = (int64)(n0)
	if err != nil { return }
	return
}

func (n *Text)String()(string){
	return n.content
}

type Br struct{
	AttrNode0
}

func NewBrNode()(n *Br){
	n = &Br{}
	n.AttrNode0 = NewAttrNode0(n)
	return
}

func (*Br)IsSimple()(bool){
	return true
}

func (*Br)Name()(string){
	return "br"
}

func (n *Br)GetText()(string){
	return "\n"
}

func (n *Br)SetText(t string){
}

func (n *Br)WriteTo(w io.Writer)(written int64, err error){
	var n0 int
	n0, err = w.Write(([]byte)("<br/>"))
	written = (int64)(n0)
	if err != nil { return }
	return
}

func (n *Br)String()(string){
	return "<br/>"
}

var (
	_ Node = (*Text)(nil)
)

