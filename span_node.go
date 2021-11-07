
package gquery

import (
	strings "strings"
)

type Span struct{
	ParentNode0
	id string
}

func NewSpanNode(id string)(p *Span){
	p = &Span{
		id: strings.ToLower(id),
	}
	p.ParentNode0 = NewParentNode0(p)
	return
}

func (p *Span)Name()(string){
	return p.id
}

func (p *Span)String()(str string){
	return "<" + p.id + p.AttrString() + ">" + p.ContentString() + "</" + p.id + ">"
}

type SimpleSpan struct{
	AttrNode0
	id string
}

func NewSimpleSpanNode(id string)(p *SimpleSpan){
	p = &SimpleSpan{
		id: strings.ToLower(id),
	}
	p.AttrNode0 = NewAttrNode0(p)
	return
}

func (p *SimpleSpan)IsSimple()(bool){
	return true
}

func (p *SimpleSpan)Name()(string){
	return p.id
}

func (n *SimpleSpan)GetText()(string){
	return ""
}

func (n *SimpleSpan)SetText(t string){
}

func (p *SimpleSpan)String()(str string){
	return "<" + p.id + p.AttrString() + "/>"
}

var (
	_ Node = (*Span)(nil)
	_ Node = (*SimpleSpan)(nil)
)
