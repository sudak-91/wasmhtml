package element

import "github.com/sudak-91/wasmhtml"

type Li struct {
	CommonElement
	Type  string `htmlattr:"type"`
	Value string `htmlattr:"Value"`
}

func NewLi() *Li {
	var li Li
	li.CommonElement.Object = wasmhtml.CreateElement("li")
	return &li
}

func (l *Li) SetInnerHtml(value string) {
	wasmhtml.InnerHtml(l.Object, value)
}

func (l *Li) SetType(value string) {
	wasmhtml.Set(l.Object, "type", value)
}

func (l *Li) SetValue(value string) {
	wasmhtml.Set(l.Object, "value", value)
}
