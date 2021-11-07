
package gquery

type Comment struct{
	Node0
	content string
}

func NewCommentNode(text string)(n *Comment){
	n = &Comment{
		content: text,
	}
	n.Node0 = NewNode0(n)
	return
}

func (*Comment)IsSimple()(bool){
	return true
}

func (*Comment)Name()(string){
	return "#comment"
}

func (n *Comment)GetText()(string){
	return ""
}

func (n *Comment)SetText(t string){
}

func (n *Comment)GetValue()(string){
	return n.content
}

func (n *Comment)SetValue(t string){
	n.content = t
}

func (n *Comment)String()(string){
	return "<!--" + n.content + "-->"
}

var (
	_ Node = (*Comment)(nil)
)

