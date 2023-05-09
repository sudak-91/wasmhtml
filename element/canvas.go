package element

import (
	"syscall/js"

	"github.com/sudak-91/wasmhtml"
)

type Canvas struct {
	CommonElement
	Context js.Value
	Height  string `htmlattr:"height"`
	Width   string `htmlattr:"width"`
}

func NewCanvas() *Canvas {
	var can Canvas
	can.CommonElement.Object = wasmhtml.CreateElement("canvas")
	return &can
}

func (c *Canvas) SetHeight(value string) {
	wasmhtml.Set(c.Object, "height", value)
}

func (c *Canvas) SetWidth(value string) {
	wasmhtml.Set(c.Object, "width", value)
}

func (c *Canvas) Get2DContext() {
	c.Context = wasmhtml.GetContext(c.Object, "2d")
}
