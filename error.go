package main

import (
	"fmt"
)

type DivideError struct {
	devidee int
	devided int
}

func (de *DivideError) Error() string {
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
    `
	return fmt.Sprintf(strFormat, de.dividee)
}

func Devide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{}
	}
}
