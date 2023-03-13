package element

import (
	"log"

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
func (b *Body) Generate() {
	log.Println("Body")
	for _, v := range b.Child {
		v.Generate()
	}
}
