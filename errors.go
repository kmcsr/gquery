
package gquery

import (
	fmt "fmt"
)

type UnexpectTokenError struct{
	char rune
	expects []rune
	extra string
}

func NewUnexpectTokenError(char rune, expects ...rune)(*UnexpectTokenError){
	return &UnexpectTokenError{
		char: char,
		expects: expects,
		extra: "",
	}
}

func (e *UnexpectTokenError)SetExtra(extra string)(*UnexpectTokenError){
	e.extra = extra
	return e
}

func (e *UnexpectTokenError)Error()(str string){
	str = fmt.Sprintf("Unexpected token: '%c'(%U)", e.char, e.char)
	if len(e.expects) > 0 {
		str += " expect ["
		for _, c := range e.expects {
			str += fmt.Sprintf("'%c'(%U), ", c, c)
		}
		str = str[:len(str) - 2] + "]"
	}
	str += e.extra
	return
}


var (
	_ error = (*UnexpectTokenError)(nil)
)
