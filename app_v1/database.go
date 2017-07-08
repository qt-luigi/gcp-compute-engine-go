package main

import (
	"database/sql"
	"os"
)

const (
	// QueryDML 定数は、idの降順でソートしたmessageデータを全件取得するselect文です。
	QueryDML = "select id, timestamp, name, message, filename from message order by id desc"
	// AddDML 定数は、messageデータを１件登録するinsert文です。
	AddDML = "insert into message(id, timestamp, name, message, filename) values(?, ?, ?, ?, ?)"
)

// Database 構造体は、DBとテーブルの作成、データの取得および登録を行います。
type Database struct {
	DriverName     string
	DataSourceName string
}

// CreateAll メソッドは、DBとテーブルを再作成します。
func (d Database) CreateAll(ddl string) error {
	// DB削除
	os.Remove(d.DataSourceName)
	// DBオーブン＆作成
	db, err := sql.Open(d.DriverName, d.DataSourceName)
	if err != nil {
		return err
	}
	defer db.Close()
	// テーブル定義文を実行
	_, err = db.Exec(ddl)
	if err != nil {
		return err
	}
	return nil
}

// Query メソッドは、指定された件数までのmessageデータを取得します。
func (d Database) Query(limit int) (Messages, error) {
	// DBオーブン
	db, err := sql.Open(d.DriverName, d.DataSourceName)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	// select文を実行
	rows, err := db.Query(QueryDML)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// 指定された件数までのmessageデータを保持する配列
	var messages Messages
	// 指定された件数までのmessageデータを確保
	idx := 0
	for rows.Next() {
		// 指定された件数に達したらループ終了
		if idx >= limit {
			break
		}
		// DBデータをMessage構造体に変換
		m := Message{}
		if err = rows.Scan(&m.ID, &m.Timestamp, &m.Name, &m.Message, &m.Filename); err != nil {
			return nil, err
		}
		messages = append(messages, m)
		idx++
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return messages, nil
}

// gen 変数は、発番を保持します。（簡易実装のため重複発番する可能性あり）
var gen int

// Add メソッドは、messageデータを１件登録します。
func (d Database) Add(m Message) error {
	// DBオーブン
	db, err := sql.Open(d.DriverName, d.DataSourceName)
	if err != nil {
		return err
	}
	defer db.Close()
	// トランザクション開始
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	// SQL文を準備
	stmt, err := tx.Prepare(AddDML)
	if err != nil {
		return err
	}
	defer stmt.Close()
	// id発番（簡易実装のため重複発番する可能性あり）
	gen++
	m.ID = gen
	// SQL文を実行
	_, err = stmt.Exec(m.ID, m.Timestamp, m.Name, m.Message, m.Filename)
	if err != nil {
		return err
	}
	// コミット
	tx.Commit()
	return nil
}
