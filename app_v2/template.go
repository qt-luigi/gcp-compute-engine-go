package main

import (
	"html/template"
	"net/http"
)

// templates 変数は、テンプレートファイルの整合性をチェックします。
var templates = template.Must(template.New("app").Funcs(templateFuncs).ParseFiles("templates/_formhelpers.html",
	"templates/index.html",
	"templates/messages.html",
	"templates/post.html"))

// RenderTemplate 関数は、テンプレートにdataの値を設定してレスポンスに出力します。
func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
