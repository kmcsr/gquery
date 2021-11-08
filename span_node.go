
package gquery

import (
	io "io"
	strings "strings"
)

type Span struct{
	ParentNode0
	id string
}

func NewSpanNode(id string)(n *Span){
	n = &Span{
		id: strings.ToLower(id),
	}
	n.ParentNode0 = NewParentNode0(n)
	return
}

func (n *Span)Name()(string){
	return n.id
}

func (n *Span)WriteTo(w io.Writer)(written int64, err error){
	var (
		n0 int
		n1 int64
	)
	written = 0
	n0, err = w.Write(([]byte)("<" + n.id))
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
	n0, err = w.Write(([]byte)("</" + n.id + ">"))
	written += (int64)(n0)
	if err != nil { return }
	return
}

func (n *Span)String()(str string){
	return "<" + n.id + n.AttrString() + ">" + n.ContentString() + "</" + n.id + ">"
}

type SimpleSpan struct{
	AttrNode0
	id string
}

func NewSimpleSpanNode(id string)(n *SimpleSpan){
	n = &SimpleSpan{
		id: strings.ToLower(id),
	}
	n.AttrNode0 = NewAttrNode0(n)
	return
}

func (*SimpleSpan)IsSimple()(bool){
	return true
}

func (n *SimpleSpan)Name()(string){
	return n.id
}

func (n *SimpleSpan)GetText()(string){
	return ""
}

func (n *SimpleSpan)SetText(t string){
}

func (n *SimpleSpan)WriteTo(w io.Writer)(written int64, err error){
	var n0 int
	written = 0
	n0, err = w.Write(([]byte)("<" + n.id))
	written += (int64)(n0)
	if err != nil { return }
	n0, err = w.Write(([]byte)(n.AttrString()))
	written += (int64)(n0)
	if err != nil { return }
	n0, err = w.Write(([]byte)("/>"))
	written += (int64)(n0)
	if err != nil { return }
	return
}

func (n *SimpleSpan)String()(str string){
	return "<" + n.id + n.AttrString() + "/>"
}

var (
	_ Node = (*Span)(nil)
	_ Node = (*SimpleSpan)(nil)
)
