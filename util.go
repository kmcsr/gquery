
package gquery

import (
	io "io"
	unicode "unicode"
	regexp "regexp"
	sort "sort"
)

type RuneSeekScanner interface{
	io.ReadSeeker
	io.RuneScanner
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

var empty_re = regexp.MustCompile(`\s+`)

func zipString(str string)(string){
	if len(str) == 0 {
		return str
	}
	return empty_re.ReplaceAllString(str, " ")
}

func strInStrlist(str string, list []string)(bool){
	sort.Strings(list)
	i := sort.SearchStrings(list, str)
	return list[i] == str
}
