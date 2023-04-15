package element

import "github.com/sudak-91/wasmhtml"

type Ul struct {
	CommonElement
}

func NewUl() *Ul {
	var u Ul
	u.Object = wasmhtml.CreateElement("ul")
	return &u
}

func (u *Ul) SetType(ulType string) {
	wasmhtml.Set(u.Object, "type", ulType)
}
