package main

import (
	"fmt"
)

// Validator インターフェースは、検証メソッドを提供します。
type Validator interface {
	Validate(value string) error
}

// Length 構造体は、長さに関する情報を保持します。
type Length struct {
	Min     int
	Max     int
	Message string
}

// Validate メソッドは、valueの文字列長の範囲を検証します。
func (l Length) Validate(value string) error {
	ln := len(value)
	if l.Min > ln || ln > l.Max {
		// メッセージ文の取得元：
		// https://github.com/wtforms/wtforms/blob/master/wtforms/validators.py#L103
		return fmt.Errorf("Field must be between %d and %d characters long", l.Min, l.Max)
	}
	return nil
}
