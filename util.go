
package gquery

import (
	io "io"
	os "os"
	bytes "bytes"
	unicode "unicode"
	utf8 "unicode/utf8"
	sort "sort"
	errors "errors"
)

type RuneSeekScanner interface{
	io.ReadSeeker
	io.RuneScanner
}

type osFileWrapper struct{
	*os.File
	lastRuneLen int
}

func newOsFileWrapper(fd *os.File)(*osFileWrapper){
	return &osFileWrapper{
		File: fd,
		lastRuneLen: -1,
	}
}

func (fd *osFileWrapper)ReadRune()(r rune, i int, err error){
	fd.lastRuneLen = -1
	buf := make([]byte, utf8.UTFMax)
	i = 1
	_, err = fd.Read(buf[:i])
	if err != nil { return }
	if !utf8.RuneStart(buf[0]) {
		r = utf8.RuneError
		return
	}
	for ; !utf8.Valid(buf[:i]) ; i++{
		if i >= utf8.UTFMax {
			r = utf8.RuneError
			return
		}
		_, err = fd.Read(buf[i:i + 1])
		if err != nil { return }
	}
	r, _ = utf8.DecodeRune(buf)
	fd.lastRuneLen = i
	return
}

func (fd *osFileWrapper)UnreadRune()(err error){
	if fd.lastRuneLen < 0 {
		return errors.New("gquery.osFileWrapper.UnreadRune: previous operation was not ReadRune")
	}
	_, err = fd.Seek((int64)(-fd.lastRuneLen), io.SeekCurrent)
	return
}

var (
	_ RuneSeekScanner = (*bytes.Reader)(nil)
	_ RuneSeekScanner = (*osFileWrapper)(nil)
)

func openFileToSeekScanner(path string)(s RuneSeekScanner, err error){
	var (
		fd *os.File
		info os.FileInfo
	)
	fd, err = os.Open(path)
	if err != nil { return }
	defer func(){ if fd != nil { fd.Close() } }()
	info, err = fd.Stat()
	if err == nil && info.Size() < 256 * 1024 { // < 256KB
		var bts []byte
		bts, err = io.ReadAll(fd)
		if err != nil { return }
		s = bytes.NewReader(bts)
	}else{
		s = newOsFileWrapper(fd)
		fd = nil
	}
	return
}

func skipWhites(r io.RuneScanner)(err error){
	var b rune
	for {
		b, _, err = r.ReadRune()
		if err != nil { return }
		if !unicode.IsSpace(b) {
			err = r.UnreadRune()
			return
		}
	}
}

func readToByte(r io.RuneScanner, t rune)(buf []rune, err error){
	buf = make([]rune, 0)
	var (
		b rune
	)
	for {
		b, _, err = r.ReadRune()
		if err != nil { return }
		if b == t {
			return
		}
		buf = append(buf, b)
	}
}

func readToStr(r io.RuneScanner, t string)(buf []rune, err error){
	rl := len(([]rune)(t))
	buf = make([]rune, 0)
	var (
		rbf = make([]rune, rl)
	)
	for i, _ := range rbf {
		rbf[i], _, err = r.ReadRune()
		if err != nil { return }
	}
	for (string)(rbf) != t {
		buf = append(buf, rbf[0])
		copy(rbf, rbf[1:])
		rbf[rl - 1], _, err = r.ReadRune()
		if err != nil { return }
	}
	return
}

func isWordLetter(b rune)(bool){
	return unicode.IsLetter(b) || unicode.IsDigit(b) || b == '-' || b == '_'
}

func readWord(r io.RuneScanner)(_ string, b rune, err error){
	var (
		word []rune = make([]rune, 0)
	)
	for {
		b, _, err = r.ReadRune()
		if err != nil { return }
		if !isWordLetter(b) {
			break
		}
		word = append(word, b)
	}
	return (string)(word), b, nil
}

func readString(r io.RuneScanner)(_ string, err error){
	var (
		b rune
		str []rune = make([]rune, 0)
	)
	b, _, err = r.ReadRune()
	if err != nil { return }
	if b == '\'' || b == '"' {
		str, err = readToByte(r, b)
		if err != nil { return }
		return (string)(str), nil
	}
	for {
		b, _, err = r.ReadRune()
		if err != nil { return }
		if unicode.IsSpace(b) || b == '\'' || b == '"' || b == '<' || b == '>' {
			err = r.UnreadRune()
			break
		}
		str = append(str, b)
	}
	return (string)(str), nil
}

func emptyString(str string)(bool){
	for _, c := range str {
		if !unicode.IsSpace(c) {
			return false
		}
	}
	return true
}

func standardizedText(str string)(string){
	if emptyString(str) {
		return ""
	}
	return empty_re.ReplaceAllString(str, " ")
}

func sortStrings(list []string)([]string){
	sort.Strings(list)
	return list
}

func asSortStrings(list ...string)([]string){
	return sortStrings(list)
}

func strInList(str string, list []string)(bool){
	i := sort.SearchStrings(list, str)
	return -1 < i && i < len(list) && list[i] == str
}

var simpleNodeList = asSortStrings(
	"br", "hr",
	"meta", "link",
	"input", "img",
)

var blockNodeList = asSortStrings(
	"div", "p",
)

func IsSimpleNodeName(id string)(ok bool){
	return strInList(id, simpleNodeList)
}

func IsBlockNodeName(id string)(ok bool){
	return strInList(id, blockNodeList)
}
