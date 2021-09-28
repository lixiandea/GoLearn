package common

type Message struct {
	user string
	msg  string
}

func (m *Message) ToString() string {
	return "[" + m.user + "]: " + m.msg
}
