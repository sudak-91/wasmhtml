package element

import (
	"github.com/sudak-91/wasmhtml"
)

type Body struct {
	CommonElement
}

func GetBody() *Body {
	var b Body
	b.Object = wasmhtml.QuerySelector(wasmhtml.Document, "body")
	return &b
}
