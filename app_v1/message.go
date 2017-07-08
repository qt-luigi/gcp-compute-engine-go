package main

// Message 構造体はmessageデータのモデルです。
type Message struct {
	ID        int
	Timestamp string
	Name      string
	Message   string
	Filename  string
}

// Messages はMessage構造体のスライスです。
type Messages []Message

// Reverse メソッドはMessagesスライスの要素の並びを反転します。
func (m Messages) Reverse() {
	for i, j := 0, len(m)-1; i < j; i, j = i+1, j-1 {
		m[i], m[j] = m[j], m[i]
	}
}
