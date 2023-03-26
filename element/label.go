package element

import (
	"github.com/sudak-91/wasmhtml"
)

type Label struct {
	CommonElement
	For       string `htmlattr:"for"`
	InnerHtml string
}

func NewLabel() Label {
	var l Label
	l.Object = wasmhtml.CreateElement("label")
	return l
}

func (l *Label) Generate() {
	for _, v := range l.Child {
		if data, ok := v.(HtmlGenerator); ok {
			data.Generate()
			l.appendChild(data.GetObject())
		}

	}
	wasmhtml.InnerHtml(l.Object, l.InnerHtml)

}
