package model

type Hello struct {
	Name string
}

func (h *Hello) SayHello() string {
	return "Hello" + h.Name
}
