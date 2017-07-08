package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

// templates 変数は、テンプレートファイルの整合性をチェックします。
var templates = template.Must(template.New("app").Funcs(templateFuncs).ParseFiles(filepath.Join(installPath, "templates/_formhelpers.html"),
	filepath.Join(installPath, "templates/index.html"),
	filepath.Join(installPath, "templates/messages.html"),
	filepath.Join(installPath, "templates/post.html")))

// RenderTemplate 関数は、テンプレートにdataの値を設定してレスポンスに出力します。
func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
