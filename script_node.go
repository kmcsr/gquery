
package gquery

import (
	io "io"
	strings "strings"
)

type Script struct{
	AttrNode0
	content string
}

func NewScriptNode(contents ...string)(n *Script){
	n = &Script{
		content: strings.Join(contents, "\n"),
	}
	n.AttrNode0 = NewAttrNode0(n)
	return
}

func (*Script)IsSimple()(bool){
	return false
}

func (*Script)Name()(string){
	return "script"
}

func (n *Script)GetText()(string){
	return ""
}

func (n *Script)SetText(t string){
}

func (n *Script)GetValue()(string){
	return n.content
}

func (n *Script)SetValue(t string){
	n.content = t
}

func (n *Script)WriteTo(w io.Writer)(written int64, err error){
	var n0 int
	written = 0
	n0, err = w.Write(([]byte)("<script"))
	written += (int64)(n0)
	if err != nil { return }
	n0, err = w.Write(([]byte)(n.AttrString()))
	written += (int64)(n0)
	if err != nil { return }
	n0, err = w.Write(([]byte)(">"))
	written += (int64)(n0)
	if err != nil { return }
	n0, err = w.Write(([]byte)(n.content))
	written += (int64)(n0)
	if err != nil { return }
	n0, err = w.Write(([]byte)("</script>"))
	written += (int64)(n0)
	if err != nil { return }
	return
}

func (n *Script)String()(string){
	return "<script" + n.AttrString() + ">" + n.content + "</script>"
}

type Style struct{
	AttrNode0
	content string
}

func NewStyleNode(contents ...string)(n *Style){
	n = &Style{
		content: strings.Join(contents, ""),
	}
	n.AttrNode0 = NewAttrNode0(n)
	return
}

func (*Style)IsSimple()(bool){
	return false
}

func (*Style)Name()(string){
	return "style"
}

func (n *Style)GetText()(string){
	return ""
}

func (n *Style)SetText(t string){
}

func (n *Style)GetValue()(string){
	return n.content
}

func (n *Style)SetValue(t string){
	n.content = t
}

func (n *Style)WriteTo(w io.Writer)(written int64, err error){
	var n0 int
	written = 0
	n0, err = w.Write(([]byte)("<style"))
	written += (int64)(n0)
	if err != nil { return }
	n0, err = w.Write(([]byte)(n.AttrString()))
	written += (int64)(n0)
	if err != nil { return }
	n0, err = w.Write(([]byte)(">"))
	written += (int64)(n0)
	if err != nil { return }
	n0, err = w.Write(([]byte)(n.content))
	written += (int64)(n0)
	if err != nil { return }
	n0, err = w.Write(([]byte)("</style>"))
	written += (int64)(n0)
	if err != nil { return }
	return
}

func (n *Style)String()(string){
	return "<style" + n.AttrString() + ">" + n.content + "</style>"
}

var (
	_ Node = (*Script)(nil)
	_ Node = (*Style)(nil)
)
