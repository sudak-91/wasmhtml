package element

import (
	"syscall/js"

	"github.com/sudak-91/wasmhtml"
)

type Input struct {
	CommonElement
}

func (i *Input) GetValue() js.Value {
	return wasmhtml.GetValue(i.Object)
}

func (i *Input) SetValue(value string) {
	wasmhtml.SetValue(i.Object, value)
}
func NewInput() *Input {
	var input Input
	input.CommonElement.Object = wasmhtml.CreateElement("input")
	return &input
}
func (i *Input) SetType(inputType string) {
	wasmhtml.Set(i.Object, "type", inputType)
}

const (
	AutocompleteON  = "on"
	AutocompleteOFF = "off"
)
