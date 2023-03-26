package element

import (
	"syscall/js"

	"github.com/sudak-91/wasmhtml"
)

type Div struct {
	CommonElement
}

func NewDiv() *Div {
	var div Div
	div.CommonElement.Object = wasmhtml.CreateElement("div")
	return &div
}

func (d *Div) SetAlign(align string) {
	wasmhtml.SetAlign(d.Object, align)
}

func (d *Div) GetAlign() string {
	return wasmhtml.GetAlign(d.Object)
}

func (d *Div) GetObject() js.Value {
	return d.CommonElement.Object
}

func (d *Div) SetInnerHtml(value string) {
	wasmhtml.InnerHtml(d.Object, value)
}
