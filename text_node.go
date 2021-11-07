
package gquery

import (
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

func (n *Text)GetText()(string){
	return html.UnescapeString(n.content)
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

func (n *Text)String()(string){
	return n.content
}

var (
	_ Node = (*Text)(nil)
)

