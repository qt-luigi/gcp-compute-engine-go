package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// path 変数は、インストールパスです。
// ローカル実行の際は""に変更してください。
const installPath = "/opt/dengonban/v2"

// #### Edit Here
const (
	dbuser     = "appuser"
	dbpass     = "pas4appuser"
	dbinstance = "<Your_Instance_Connection_Name>"
	dbname     = "message_db"
)

// db 変数は、データベース情報を保持します。
var db Database

// ddl 定数は、DBに作成するテーブルの定義文です。
const ddl = `
	create table if not exist message (
		id integer primary key,
		timestamp varchar(19),
		name varchar(16),
		message varchar(1024),
		filename varchar(128)
	);`

// 本init 関数では、データベースとテーブルを再作成します。
func init() {
	dsn := fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s", dbuser, dbpass, dbinstance, dbname)
	db = Database{DriverName: "mysql", DataSourceName: dsn}
	if err := db.CreateAll(ddl); err != nil {
		log.Fatal(err)
	}
}

// templateFuncs 変数はテンプレートで使用する独自関数を定義します。
var templateFuncs = template.FuncMap{
	"add_br": func(text string) template.HTML {
		return template.HTML(strings.Replace(template.HTMLEscapeString(text), "\n", "<br>", -1))
	},
	"safe": func(text string) template.HTML {
		return template.HTML(text)
	},
}

// NewMessageForm 関数は、必要な値が初期設定されたMessageForm構造体を生成します。
func NewMessageForm(r *http.Request) MessageForm {
	form := MessageForm{}
	form.InputName = NewStringField(r, "input_name", "お名前", []Validator{Length{Min: 1, Max: 16}})
	form.InputMessage = NewTextAreaField(r, "input_message", "メッセージ", []Validator{Length{Min: 1, Max: 1024}})
	return form
}

// index 関数は、トップページ画面のハンドラーです。
func index(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "index.html", nil)
}

// messages 関数は、投稿されたメッセージを入力および表示する画面のハンドラーです。
func messages(w http.ResponseWriter, r *http.Request) {
	form := NewMessageForm(r)
	lastMessages, err := db.Query(5)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	lastMessages.Reverse()
	data := struct {
		Form     MessageForm
		Messages Messages
	}{
		Form:     form,
		Messages: lastMessages,
	}
	RenderTemplate(w, "messages.html", data)
}

// post 関数は、投稿されたメッセージを受け取る画面のハンドラーです。
func post(w http.ResponseWriter, r *http.Request) {
	form := NewMessageForm(r)
	if r.Method == "POST" && form.Validate() {
		jst, err := time.LoadLocation("Asia/Tokyo")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		timestamp := time.Now().UTC().In(jst).Format("2006/01/02 15:04:05")
		name := r.FormValue("input_name")
		message := r.FormValue("input_message")
		if err := db.Add(Message{Timestamp: timestamp, Name: name, Message: message}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		data := struct {
			Name      string
			Timestamp string
		}{
			Name:      name,
			Timestamp: timestamp,
		}
		RenderTemplate(w, "post.html", data)
	} else {
		http.Redirect(w, r, "messages", http.StatusFound)
	}
}

// 本init 関数では、ルーティングとハンドラー関数を設定します。
func init() {
	// 静的ファイルパスを設定
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(filepath.Join(installPath, "static")))))
	// 動的ページのパスとハンドラーを設定
	http.HandleFunc("/", index)
	http.HandleFunc("/messages", messages)
	http.HandleFunc("/post", post)
}

// main 関数は、エントリーポイントです。
func main() {
	// ローカル実行の際は":8080"に変更して「http://localhost:8080」でアクセスしてください。
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
