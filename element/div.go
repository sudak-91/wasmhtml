package element

import "github.com/sudal-91/wasmhtml"

type Div struct {
	CommonElement
	Align     string `htmlattr:"align"`
	Title     string `htmlattr:"title"`
	InnerHtml string
}

func NewDiv() *Div {
	var div Div
	div.Object = wasmhtml.CreateElement("div")
	return &div
}

func (d *Div) Generate() {
	for _, v := range d.Child {
		v.Generate()
	}
	d.CommonElement.Generate()
	wasmhtml.Set(d.Object, "align", d.Align)
	wasmhtml.Set(d.Object, "title", d.Title)
	wasmhtml.InnerHtml(d.Object, d.InnerHtml)

}
