package element

import (
	"log"
	"syscall/js"

	"github.com/sudak-91/wasmhtml"
)

type Input struct {
	CommonElement
	Accept         string `htmlattr:"accept"` //Set filter to files type
	Align          string `htmlattr:"align"`
	Alt            string `htmlattr:"alt"` //Alternative text for image ot button
	Autocomplete   string `htmlattr:"autocomplete"`
	Autofocus      bool   `htmlattr:"autofocus"`
	Border         string `htmlattr:"border"`
	Checked        bool   `htmlattr:"checked"`
	Disabled       bool   `htmlattr:"disabled"`
	Form           string `htmlattr:"form"`
	FormAction     string `htmlattr:"formaction"`
	FormEncType    string `htmlattr:"formenctype"`
	FormMethod     string `htmlattr:"formmethod"`
	FormNoValidate bool   `htmlattr:"formnovalidate"`
	FormTarget     string `htmlattr:"formtarget"`
	List           string `htmlattr:"list"`
	Max            string `htmlattr:"max"`
	MaxLength      string `htmlattr:"maxlength"`
	Min            string `htmlattr:"min"`
	Multiple       bool   `htmlattr:"multiple"`
	Name           string `htmlattr:"name"`
	Pattern        string `htmlattr:"pattern"`
	Placeholder    string `htmlattr:"placeholder"`
	Readonly       bool   `htmlattr:"readonly"`
	Required       bool   `htmlattr:"required"`
	Size           string `htmlattr:"size"`
	Src            string `htmlattr:"src"`
	Step           string `htmlattr:"step"` //incremental step
	Tabindex       string `htmlattr:"tabindex"`
	Type           string `htmlattr:"type"`
}

func (i *Input) GetValue() js.Value {
	return wasmhtml.GetValue(i.Object)
}

func (i *Input) SetValue(value string) {
	wasmhtml.SetValue(i.Object, value)
}
func NewInput() Input {
	var input Input
	input.Object = wasmhtml.CreateElement("input")
	return input
}
func (i *Input) Generate() {
	log.Printf("start generate %s", i.Object)
	for _, v := range i.Child {
		if data, ok := v.(HtmlGenerator); ok {
			data.Generate()
			i.appendChild(data.GetObject())
		}

	}

}

const (
	AutocompleteON  = "on"
	AutocompleteOFF = "off"
)
