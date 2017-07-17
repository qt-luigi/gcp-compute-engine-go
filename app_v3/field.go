package main

import (
	"mime/multipart"
	"net/http"
)

// Field 構造体は、formのHTMLタグに関する情報を保持します。
type Field struct {
	Name       string
	Label      string
	Value      string
	Validators []Validator
	Errors     []string
}

// Validate メソッドは、Validators内に登録されている検証を実行してエラーがあればメッセージをErrorsに登録します。
func (f *Field) Validate() bool {
	for _, v := range f.Validators {
		if err := v.Validate(f.Value); err != nil {
			f.Errors = append(f.Errors, err.Error())
		}
	}
	return len(f.Errors) == 0
}

// StringField 構造体は、textタグに関する情報を保持します。
type StringField struct {
	Field
	Type string
}

// TextAreaField 構造体は、textareaタグに関する情報を保持します。
type TextAreaField struct {
	Field
}

// UploadField 構造体は、uploadタグに関する情報を保持します。
type UploadField struct {
	Field
	Type string

	File        multipart.File
	FileHandler *multipart.FileHeader
	FileError   error
}

// NewStringField 関数は、指定された値で初期化されたStringField構造体を生成します。
func NewStringField(r *http.Request, name, label string, validators []Validator) StringField {
	f := Field{Name: name, Label: label, Value: r.FormValue(name), Validators: validators}
	return StringField{Field: f, Type: "text"}
}

// NewTextAreaField 関数は、指定された値で初期化されたTextAreaField構造体を生成します。
func NewTextAreaField(r *http.Request, name, label string, validators []Validator) TextAreaField {
	f := Field{Name: name, Label: label, Value: r.FormValue(name), Validators: validators}
	return TextAreaField{Field: f}
}

// NewUploadField 関数は、指定された値で初期化されたUploadField構造体を生成します。
func NewUploadField(r *http.Request, name, label string, validators []Validator) UploadField {
	f := Field{Name: name, Label: label, Value: r.FormValue(name), Validators: validators}
	file, handler, err := r.FormFile(name)
	return UploadField{Field: f, Type: "file", File: file, FileHandler: handler, FileError: err}
}
