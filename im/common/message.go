package common

type Message struct {
	User string
	Msg  string
}

func (m *Message) ToString() string {
	return "[" + m.User + "]: " + m.Msg
}
