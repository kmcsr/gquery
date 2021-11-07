
package gquery

import (
	io "io"
	bytes "bytes"
	unicode "unicode"
	strings "strings"
	fmt "fmt"
)

type EndNode struct{
	Node0
	id string
}

func NewEndNode(id string)(*EndNode){
	return &EndNode{
		Node0: NewNode0(nil),
		id: id,
	}
}

func (p *EndNode)Name()(string){
	return p.id
}

func (n *EndNode)GetText()(string){
	return ""
}

func (n *EndNode)SetText(t string){
}

func (p *EndNode)String()(str string){
	return ""
}

type unexpectEndNode struct{
	node *EndNode
}

func newUnexpectEndNode(node *EndNode)(*unexpectEndNode){
	return &unexpectEndNode{
		node: node,
	}
}

func (e *unexpectEndNode)Error()(string){
	return fmt.Sprintf("Unexpected end node: %s", e.node.String())
}

func (e *unexpectEndNode)GetNode()(*EndNode){
	return e.node
}


var _ RuneSeekScanner = (*bytes.Reader)(nil)

func DecodeDocString(str string)(doc *Document, err error){
	return DecodeDoc(strings.NewReader(str))
}

func DecodeDocBytes(bts []byte)(doc *Document, err error){
	return DecodeDoc(bytes.NewReader(bts))
}

func DecodeDoc(r RuneSeekScanner)(doc *Document, err error){
	doc = NewDocumentNode()
	var (
		n Node
		ok bool = true
	)
	for ok {
		n, err = DecodeNode(r)
		if err != nil { return }
		ok = n.Name() == "#comment" || n.Name() == "#text"
	}
	if dtn, ok := n.(*DocType); ok {
		doc.SetDocType(dtn)
		for ok {
			n, err = DecodeNode(r)
			if err != nil { return }
			ok = n.Name() == "#comment" || n.Name() == "#text"
		}
	}
	if hn, ok := n.(*HtmlNode); ok {
		doc.SetHtmlNode(hn)
		return
	}
	return nil, fmt.Errorf("Unexpected node: (%s)%s", n.Name(), n.String())
}

func DecodeNode(r RuneSeekScanner)(nd Node, err error){
	var (
		b rune
		buf []byte
		rbf []rune
		sbf string
		ok bool
	)
	b, _, err = r.ReadRune()
	if err != nil { return }
	if b == '<' {
		b, _, err = r.ReadRune()
		if err != nil { return }
		if b == '/' {
			rbf, b, err = decodeTagPart(r)
			if b != '>' {
				if unicode.IsSpace(b) {
					err = skipWhites(r)
					if err != nil { return }
					b, _, err = r.ReadRune()
				}
			}
			if b != '>' {
				return nil, NewUnexpectTokenError(b, '>').SetExtra(" for </" + (string)(rbf))
			}
			return NewEndNode(strings.ToLower((string)(rbf))), nil
		}
		if b == '!' {
			buf = make([]byte, 2)
			_, err = r.Read(buf)
			if err != nil { return }
			if (string)(buf) == "--" {
				rbf, err = readToStr(r, "-->")
				return NewCommentNode((string)(rbf)), nil
			}
			_, err = r.Seek(-2, io.SeekCurrent)
			if err != nil { return }
			sbf, b, err = readWord(r)
			if !unicode.IsSpace(b) {
				return nil, NewUnexpectTokenError(b).SetExtra(" expect spaces")
			}
			err = skipWhites(r)
			if err != nil { return }
			id := strings.ToLower(sbf)
			switch id {
			case "doctype":
				rbf, b, err = decodeTagPart(r)
				if err != nil { return }
				if len(rbf) == 0 {
					return nil, NewUnexpectTokenError(b).SetExtra(" expect a doc type id for <!DOCTYPE")
				}
				typ, ext := (string)(rbf), make([]string, 0)
				for unicode.IsSpace(b) {
					rbf, b, err = decodeTagPart(r)
					if err != nil { return }
					ext = append(ext, (string)(rbf))
				}
				if b != '>' {
					return nil, NewUnexpectTokenError(b, '>').SetExtra(" for <!DOCTYPE")
				}
				return NewDocTypeNode(typ, ext...), nil
			default:
				return nil, fmt.Errorf("Unknown node: <!%s>", id)
			}
			return
		}
		err = r.UnreadRune()
		if err != nil { return }
		sbf, _, err = readWord(r)
		if err != nil { return }
		err = r.UnreadRune()
		if err != nil { return }
		err = skipWhites(r)
		if err != nil { return }
		bsp := false
		id := strings.ToLower(sbf)
		switch id {
		case "html":
			nd = NewHtmlNode()
		case "head":
			nd = NewHeadNode()
		case "style":
			nd = NewStyleNode()
			ok, err = readAttributes(r, nd)
			if err != nil { return }
			if !ok {
				rbf, err = readToStr(r, "</style>")
				if err != nil { return }
				nd.SetValue((string)(rbf))
			}
			return
		case "script":
			nd = NewScriptNode()
			ok, err = readAttributes(r, nd)
			if err != nil { return }
			if !ok {
				rbf, err = readToStr(r, "</script>")
				if err != nil { return }
				nd.SetValue((string)(rbf))
			}
			return
		default:
			if isSimpleNodeName(id) {
				nd = NewSimpleSpanNode(id)
				_, err = readAttributes(r, nd)
				if err != nil { return }
				return
			}else{
				nd = NewSpanNode(id)
				bsp = true
			}
		}
		ok, err = readAttributes(r, nd)
		if err != nil { return }
		if ok && bsp {
			nd2 := NewSimpleSpanNode(id)
			nd2.AttrNode0.attrs = nd.Attrs()
			nd = nd2
		}else{
			var nd2 Node
			for {
				nd2, err = DecodeNode(r)
				if err != nil { return }
				if nd2a, ok := nd2.(*EndNode); ok {
					if !isSimpleNodeName(nd2a.Name()) {
						if nd2a.Name() != id {
							err = newUnexpectEndNode(nd2a)
						}
						return
					}
				}else{
					nd.AppendChild(nd2)
				}
			}
		}
		return
	}
	err = r.UnreadRune()
	if err != nil { return }
	rbf, err = readToByte(r, '<')
	if err != nil { return }
	err = r.UnreadRune()
	if err != nil { return }
	nd = NewTextNode()
	nd.SetValue((string)(rbf))
	return
}

func decodeTagPart(r io.RuneScanner)(buf []rune, b rune, err error){
	var (
		b2 []rune
	)
	buf = make([]rune, 0)
	err = skipWhites(r)
	if err != nil { return }
	for {
		b, _, err = r.ReadRune()
		if err != nil { return }
		if unicode.IsSpace(b) || b == '>' || b == '/' {
			return
		}
		buf = append(buf, b)
		if b == '"' || b == '\'' {
			b2, err = readToByte(r, b)
			if err != nil { return }
			buf = append(append(buf, b2...), b)
		}
	}
}

func readAttributes(r io.RuneScanner, nd Node)(_ bool, err error){
	var (
		b rune
		key string
		value string
	)
	err = skipWhites(r)
	if err != nil { return }
	for {
		key, b, err = readWord(r)
		if err != nil { return }
		if unicode.IsSpace(b) {
			err = skipWhites(r)
			if err != nil { return }
			b, _, err = r.ReadRune()
		}
		switch b {
		case '=':
			err = skipWhites(r)
			if err != nil { return }
			value, err = readString(r)
			if err != nil { return }
			err = skipWhites(r)
			if err != nil { return }
		case '/':
			err = skipWhites(r)
			if err != nil { return }
			b, _, err = r.ReadRune()
			if err != nil { return }
			if b != '>' {
				return false, NewUnexpectTokenError(b).SetExtra(" for simple node end")
			}
			return true, nil
		case '>':
			return false, nil
		default:
			err = r.UnreadRune()
			if err != nil { return }
			value = key
		}
		if nd != nil {
			nd.SetAttr(key, value)
		}
	}
}

var simpleNodeMap = map[string]struct{}{
	"br": struct{}{},
	"hr": struct{}{},
	"meta": struct{}{},
	"link": struct{}{},
	"input": struct{}{},
	"img": struct{}{},
}

func isSimpleNodeName(id string)(ok bool){
	_, ok = simpleNodeMap[id]
	return
}

