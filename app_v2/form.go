package main

// MessageForm 構造体は、input項目群です。
type MessageForm struct {
	InputName    StringField
	InputMessage TextAreaField
}

// Validate メソッドは、input項目の入力チェックを行います。
func (m *MessageForm) Validate() bool {
	nv := m.InputName.Validate()
	mv := m.InputMessage.Validate()
	return nv && mv
}
