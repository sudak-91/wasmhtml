package element

import (
	"log"
	"reflect"
	"syscall/js"

	"github.com/sudak-91/wasmhtml"
	"github.com/sudak-91/wasmhtml/attribute"
	"github.com/sudak-91/wasmhtml/event"
)

type HtmlGenerator interface {
	Generate()
}

type CommonElement struct {
	Object js.Value
	Child  []HtmlGenerator
	Class  []string
	attribute.BaseAttribute
	event.BaseEvent
}

func (c *CommonElement) AddDiv() *Div {
	var div Div
	div.Object = wasmhtml.CreateElement("div")
	wasmhtml.AppendChild(c.Object, div.Object)
	c.Child = append(c.Child, &div)
	return &div
}
func (c *CommonElement) AddClass(class string) {
	c.Class = append(c.Class, class)
}
func (c *CommonElement) Generate() {
	for _, class := range c.Class {
		wasmhtml.AddClass(c.Object, class)
	}

	reflectValue := reflect.ValueOf(*c)
	for i := 0; i < reflectValue.NumField(); i++ {
		if reflectValue.Field(i).Kind() == reflect.Struct {
			c.subValueGenerate(reflectValue.Field(i))
		}
	}
}

func (c *CommonElement) subValueGenerate(value reflect.Value) {
	subType := value.Type()
	for i := 0; i < subType.NumField(); i++ {
		tag, ok := subType.Field(i).Tag.Lookup("htmlattr")
		if ok {
			c.typeSelector(value.Field(i), tag)
		}
		event, ok := subType.Field(i).Tag.Lookup("htmlevnt")
		if ok {
			wasmhtml.SetAttribute(c.Object, event, value.Field(i).String())
		}
		continue
	}
}

func (c *CommonElement) typeSelector(value reflect.Value, tag string) {
	switch v := value.Kind(); v {
	case reflect.String:
		log.Println(tag)
		log.Println(value)
		if value.String() != "" {
			wasmhtml.Set(c.Object, tag, value.String())
		}
	case reflect.Bool:
		if value.Bool() {
			wasmhtml.Set(c.Object, tag, "")
		}
	default:
		log.Println("Invalid data type")
	}
}
