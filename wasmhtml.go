package wasmhtml

import (
	"syscall/js"
)

var Document = js.Global().Get("document")

type Elements interface {
	NewDiv(id string) js.Value
	NewButton(id string) js.Value
	NewInput(id string) js.Value
}
