
package gquery

import (
	io "io"
	strings "strings"
	regexp "regexp"
)

var empty_re = regexp.MustCompile(`\s+`)

type Node interface{
	Name()(string)
	IsSimple()(bool)
	HasAttr(k string)(bool)
	GetAttrDefault(k string, d string)(v string)
	GetAttr(k string)(v string)
	SetAttr(k string, v string)
	DelAttr(k string)(bool)
	Attrs()(map[string]string)
	GetText()(string)
	SetText(string)
	GetValue()(string)
	SetValue(string)
	HasChildren()(bool)
	GetNodeList()(*NodeList)
	AppendChild(Node)
	RemoveChild(Node)
	FindFunc(func(Node)(bool))([]Node)
	Find(string)([]Node)
	ForEachAllChildren(func(n Node))
	setParent(Node)
	Parent()(Node)
	setBefore(Node)
	Before()(Node)
	setAfter(Node)
	After()(Node)
	WriteTo(w io.Writer)(written int64, err error)
	String()(string)
}

type NodeList struct{
	head Node
	tail Node
	ins Node
}

func (l *NodeList)Clear(){
	l.head = nil
	l.tail = nil
}

func (l *NodeList)Parent()(Node){
	if l.head == nil {
		return nil
	}
	return l.head.Parent()
}

func (l *NodeList)First()(Node){
	return l.head
}

func (l *NodeList)Last()(Node){
	return l.tail
}

func (l *NodeList)Append(n Node){
	if p := n.Parent(); p != nil {
		p.RemoveChild(n)
	}
	n.setParent(l.ins)
	if l.tail == nil {
		l.head = n
		l.tail = n
	}else{
		l.tail.setAfter(n)
		l.tail = n
	}
}

func (l *NodeList)Remove(n Node){
	if n.Parent() != l.Parent(){
		panic("The node not in this node list")
	}
	if n == l.tail {
		l.tail = n.Before()
		n.setParent(nil)
		if l.tail == nil {
			l.head = nil
		}else{
			l.tail.setAfter(nil)
		}
		return
	}
	c := l.head
	for c != nil {
		if c == n {
			c.Before().setAfter(c.After())
			n.setParent(nil)
			return
		}
		c = c.After()
	}
}

func (l *NodeList)Insert(n Node, t Node){ // after t
	if p := n.Parent(); p != nil {
		p.RemoveChild(n)
	}
	if t == nil {
		if l.head == nil {
			l.tail = n
		}else{
			n.setAfter(l.head)
		}
		l.head = n
		return
	}
	if t.Parent() != l.Parent(){
		panic("The target node not in this node list")
	}
	n.setParent(l.Parent())
	n.setAfter(t.After())
	t.setAfter(n)
	if n.After() == nil {
		l.tail = n
	}
}

func (l *NodeList)IndexOf(n Node)(i int){
	c := l.head
	i = 0
	for c != nil {
		if c == n {
			return
		}
		c = c.After()
		i++
	}
	return -1
}

func (l *NodeList)ForEach(call func(n Node, i int)){
	c, i := l.head, 0
	for c != nil {
		call(c, i)
		c = c.After()
		i++
	}
}

func (l *NodeList)Iter()(func()(n Node, ok bool)){
	c := l.head
	end := c == nil
	return func()(n Node, _ bool){
		if end {
			return nil, false
		}
		n = c
		c = c.After()
		if c == nil {
			end = true
		}
		return n, true
	}
}

func (l *NodeList)AsList()(s []Node){
	s = make([]Node, 0)
	c := l.head
	for c != nil {
		s = append(s, c)
		c = c.After()
	}
	return
}

type Node0 struct{
	ins Node
	parent Node
	before Node
	after Node
}

func NewNode0(ins Node)(Node0){
	return Node0{
		ins: ins,
	}
}

func (*Node0)IsSimple()(bool){
	return false
}

func (*Node0)HasAttr(k string)(bool){
	return false
}

func (*Node0)GetAttrDefault(k string, d string)(v string){
	panic("Simple node don't have any attribute")
}

func (*Node0)GetAttr(k string)(v string){
	panic("Simple node don't have any attribute")
}

func (*Node0)SetAttr(k string, v string){
	panic("Simple node don't have any attribute")
}

func (*Node0)DelAttr(k string)(bool){
	panic("Simple node don't have any attribute")
}

func (*Node0)Attrs()(map[string]string){
	panic("Simple node don't have any attribute")
}

func (*Node0)GetValue()(string){
	return ""
}

func (*Node0)SetValue(t string){
}

func (*Node0)HasChildren()(bool){
	return false
}

func (*Node0)GetNodeList()(*NodeList){
	panic("Simple node don't have any children")
}

func (*Node0)AppendChild(Node){
	panic("Simple node don't have any children")
}

func (*Node0)RemoveChild(Node){
	panic("Simple node don't have any children")
}

func (*Node0)FindFunc(func(c Node)(bool))([]Node){
	panic("Simple node don't have any children")
}

func (*Node0)Find(string)([]Node){
	panic("Simple node don't have any children")
}

func (n *Node0)ForEachAllChildren(call func(n Node)){
	if n.ins.HasChildren() {
		n.ins.GetNodeList().ForEach(func(c Node, _ int){
			call(c)
			if !c.IsSimple() {
				c.ForEachAllChildren(call)
			}
		})
	}
}

func (n *Node0)setParent(o Node){
	n.parent = o
	n.after = nil
	n.before = nil
}

func (n *Node0)Parent()(Node){
	return n.parent
}

func (n *Node0)setBefore(o Node){
	if n.before != o {
		old := n.before
		n.before = o
		if old != nil {
			old.setAfter(nil)
		}
		if o != nil {
			o.setAfter(n.ins)
		}
	}
}

func (n *Node0)Before()(Node){
	return n.before
}

func (n *Node0)setAfter(o Node){
	if n.after != o {
		old := n.after
		n.after = o
		if old != nil {
			old.setBefore(nil)
		}
		if o != nil {
			o.setBefore(n.ins)
		}
	}
}

func (n *Node0)After()(Node){
	return n.after
}

func (*Node0)String()(string){
	return ""
}

type AttrNode0 struct{
	Node0
	attrs map[string]string
}

func NewAttrNode0(ins Node)(n AttrNode0){
	n = AttrNode0{
		attrs: make(map[string]string),
	}
	n.Node0 = NewNode0(ins)
	return
}

func (n *AttrNode0)HasAttr(k string)(ok bool){
	_, ok = n.attrs[k]
	return
}

func (n *AttrNode0)GetAttrDefault(k string, d string)(v string){
	var ok bool
	v, ok = n.attrs[k]
	if !ok {
		return d
	}
	return
}

func (n *AttrNode0)GetAttr(k string)(string){
	return n.GetAttrDefault(k, "")
}

func (n *AttrNode0)SetAttr(k string, v string){
	n.attrs[k] = v
}

func (n *AttrNode0)DelAttr(k string)(bool){
	if n.HasAttr(k) {
		delete(n.attrs, k)
	}
	return false
}

func (n *AttrNode0)Attrs()(map[string]string){
	return n.attrs
}

func (n *AttrNode0)AttrString()(str string){
	str = ""
	for k, v := range n.attrs {
		if strings.IndexByte(v, '"') > -1 {
			str += " " + k + "='" + v + "'"
		}else{
			str += " " + k + "=\"" + v + "\""
		}
	}
	return
}

type ParentNode0 struct{
	AttrNode0
	children *NodeList
}

func NewParentNode0(ins Node)(ParentNode0){
	return ParentNode0{
		AttrNode0: NewAttrNode0(ins),
		children: &NodeList{
			ins: ins,
		},
	}
}

func (n *ParentNode0)GetText()(str string){
	str = ""
	n.children.ForEach(func(n Node, _ int){
		str += n.GetText()
	})
	if IsBlockNodeName(n.ins.Name()) && len(str) > 0 {
		if str[0] != '\r' && str[0] != '\n' {
			str = "\n" + str
		}
		if str[len(str) - 1] != '\r' && str[len(str) - 1] != '\n' {
			str += "\n"
		}
	}
	return
}

func (n *ParentNode0)SetText(t string){
	n.children.Clear()
	n.AppendChild(NewTextNode(t))
}

func (*ParentNode0)HasChildren()(bool){
	return true
}

func (n *ParentNode0)GetNodeList()(*NodeList){
	return n.children
}

func (n *ParentNode0)AppendChild(o Node){
	p := n.ins
	for p != nil {
		if p == o {
			panic("Can not append parent to children")
		}
		p = p.Parent()
	}
	n.children.Append(o)
}

func (n *ParentNode0)RemoveChild(o Node){
	n.children.Remove(o)
}

func (n *ParentNode0)InsertChild(o Node, _t ...Node){
	var t Node = nil
	if len(_t) > 0 {
		t = _t[0]
	}
	n.children.Insert(o, t)
}

func (n *ParentNode0)FindFunc(call func(c Node)(bool))(children []Node){
	children = make([]Node, 0)
	n.children.ForEach(func(n Node, _ int){
		if n.Name()[0] != '#' && call(n) {
			children = append(children, n)
		}
		if !n.IsSimple() {
			children = append(children, n.FindFunc(call)...)
		}
	})
	return
}

func (n *ParentNode0)Find(pattern_ string)(children []Node){
	call := func(c Node)(bool){
		return true
	}
	patterns := strings.Split(pattern_, ",")
	for _, ptn := range patterns {
		c0 := call
		ptn = strings.TrimSpace(ptn)
		switch{
		case ptn[0] == '#':
			idpt := ptn[1:]
			call = func(c Node)(bool){
				return c.HasAttr("id") && strInList(idpt, sortStrings(empty_re.Split(c.GetAttr("id"), -1))) && c0(c)
			}
		case ptn[0] == '.':
			clspt := ptn[1:]
			call = func(c Node)(bool){
				return c.HasAttr("class") && strInList(clspt,sortStrings( empty_re.Split(c.GetAttr("class"), -1))) && c0(c)
			}
		case ptn[0] == '*':
			//
		default:
			npt := ptn[1:]
			call = func(c Node)(bool){
				return c.Name() == npt && c0(c)
			}
		}
	}
	return n.FindFunc(call)
}

func (n *ParentNode0)ContentWriteTo(w io.Writer)(written int64, err error){
	var n0 int64
	written = 0
	iter := n.children.Iter()
	for {
		nd, ok := iter()
		if !ok { break }
		n0, err = nd.WriteTo(w)
		written += n0
		if err != nil { return }
	}
	return
}

func (n *ParentNode0)ContentString()(str string){
	str = ""
	n.children.ForEach(func(n Node, _ int){
		str += n.String()
	})
	return
}

func FindPrevNodeFunc(n Node, call func(c Node)(bool))(c Node){
	c = n.Before()
	if c == nil {
		return
	}
	for c != nil && !call(c) {
		c = c.Before()
	}
	return
}

func FindNextNodeFunc(n Node, call func(c Node)(bool))(c Node){
	c = n.After()
	if c == nil {
		return
	}
	for c != nil && !call(c) {
		c = c.After()
	}
	return
}

func FindPrevNodeExcepts(n Node, xs ...string)(c Node){
	sortStrings(xs)
	return FindPrevNodeFunc(n, func(c Node)(bool){
		return !strInList(c.Name(), xs)
	})
}

func FindNextNodeExcepts(n Node, xs ...string)(c Node){
	sortStrings(xs)
	return FindNextNodeFunc(n, func(c Node)(bool){
		return !strInList(c.Name(), xs)
	})
}

