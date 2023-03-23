package element

import (
	"log"
	"reflect"
	"syscall/js"

	"github.com/sudak-91/wasmhtml"
	"github.com/sudak-91/wasmhtml/event"
)

type HtmlGenerator interface {
	Generate()
	GetObject() js.Value
}

type CommonElement struct {
	Object js.Value
	Child  []interface{}
	Class  []string
	event.BaseEvent
}

// Set HTML attribute ID
func (c *CommonElement) SetID(id string) {
	wasmhtml.SetId(c.Object, id)
}

// Get HTML attribute ID
func (c *CommonElement) GetID() string {
	return wasmhtml.GetID(c.Object)
}

// Set HTML attribute HIDDEN
func (c *CommonElement) SetHidden() {
	wasmhtml.SetHidden(c.Object)
}

// Remove HTML attribute HIDDEN
func (c *CommonElement) RemoveHidden() {
	wasmhtml.RemoveHidden(c.Object)
}

// Get HTML attribute HIDDEN
func (c *CommonElement) IsHidden() bool {
	return wasmhtml.GetAttribute(c.Object, "hidden").Bool()
}

// Set HTML attribute TITLE
func (c *CommonElement) SetTitle(title string) {
	wasmhtml.SetTitle(c.Object, title)
}

// Get HTML attribute TITLE
func (c *CommonElement) GetTitle() string {
	return wasmhtml.GetTitile(c.Object)
}

// Set HTML attribute ACCESSKEY
func (c *CommonElement) SetAccesskey(key string) {
	wasmhtml.SetAccesskey(c.Object, key)
}

// Get HTML attribute ACCESSKEY
func (c *CommonElement) GetAccesskey() string {
	return wasmhtml.GetAccesskey(c.Object)
}

// Remove HTML attribute ACCESSKEY
func (c *CommonElement) RemoveAccesskey() {
	wasmhtml.RemoveAccesskey(c.Object)
}

// Set HTML attribute CONTENTEDITABLE 'true'
func (c *CommonElement) EnableEdit() {
	wasmhtml.SetContentEditable(c.Object, "true")
}

// Set HTML attribute CONTENTEDITABLE 'false'
func (c *CommonElement) DisableEdit() {
	wasmhtml.SetContentEditable(c.Object, "false")
}

// Get HTML attribute CONTENTEDITABLE
// TRUE - Edit Enable
// FALSE - Edit Disable
func (c *CommonElement) IsEditable() (bool, error) {
	return wasmhtml.GetContentEditable(c.Object)
}

// Remove HTML attribute CONTENEDITABLE
func (c *CommonElement) RemoveContenEditable() {
	wasmhtml.RemoveContentEditable(c.Object)
}

// Set HTML attribute DIR left to right
func (c *CommonElement) SetLTRDir() {
	wasmhtml.SetDir(c.Object, "ltr")
}

// Set HTML attribute DIR right to left
func (c *CommonElement) SetRTLDir() {
	wasmhtml.SetDir(c.Object, "rtl")
}

// Get HTML attribute DIR
// ltr - left to right
// rtl - right to left
func (c *CommonElement) GetDir() string {
	return wasmhtml.GetDir(c.Object)
}

// Set HTML attribute LANG lang code
func (c *CommonElement) SetLang(langCode string) {
	wasmhtml.SetLang(c.Object, langCode)
}

// Get HTML attribute LANG
func (c *CommonElement) GetLang() string {
	return wasmhtml.GetLang(c.Object)
}

// Remove HTML attribure LANG
func (c *CommonElement) RemoveLang() {
	wasmhtml.RemoveLang(c.Object)
}

// Remove HTML attribute DIR
func (c *CommonElement) DefaultDir() {
	wasmhtml.RemoveDir(c.Object)
}

// Enable SPELLCHECK attribute
func (c *CommonElement) EnableSpellcheck() {
	wasmhtml.SetSpellcheck(c.Object, "true")
}

// Disable SPELLCHECK attribute
func (c *CommonElement) DisableSpellcheck() {
	wasmhtml.SetSpellcheck(c.Object, "false")
}

// Remove SPELLCHECK attribute
func (c *CommonElement) DefualtSpellcheck() {
	wasmhtml.RemoveSpellcheck(c.Object)
}

// Get SPELLCHECK attribute value
func (c *CommonElement) IsSpellcheck() (bool, error) {
	return wasmhtml.GetSpellcheck(c.Object)
}

// Set STYLE attribute
func (c *CommonElement) SetStyle(style string) {
	wasmhtml.SetStyle(c.Object, style)
}

// Remove STYLE attribute
func (c *CommonElement) RemoveStyle() {
	wasmhtml.RemoveStyle(c.Object)
}

// Get STYLE attribute
func (c *CommonElement) GetStyle() string {
	return wasmhtml.GetStyle(c.Object)
}

// Set TABINDEX attribute
func (c *CommonElement) SetTabindex(index int32) {
	wasmhtml.SetTabindex(c.Object, index)
}

// Get TABINDEX attribute
func (c *CommonElement) GetTabindex() int32 {
	return wasmhtml.GetTabindex(c.Object)
}

// Remove TABINDEX attribute
func (c *CommonElement) RemoveTabindex() {
	wasmhtml.RemoveTabindex(c.Object)
}

func (c *CommonElement) AddDiv() *Div {
	result := NewDiv()
	c.appendChild(result.Object)
	return result
}
func (c *CommonElement) appendChild(child js.Value) {
	c.Object.Call("appendChild", child)
}

func (c *CommonElement) AddInput() *Input {
	var input Input
	input.Object = wasmhtml.CreateElement("input")
	c.Child = append(c.Child, &input.CommonElement)
	return &input
}

func (c *CommonElement) AddLabel() *Label {
	var label Label
	label.Object = wasmhtml.CreateElement("label")
	c.Child = append(c.Child, label.CommonElement)
	return &label
}
func (c *CommonElement) AddClass(class string) {
	c.Class = append(c.Class, class)
}

func (c *CommonElement) Generate() {

	for _, class := range c.Class {
		wasmhtml.AddClass(c.Object, class)
	}
	for _, v := range c.Child {
		if data, ok := v.(HtmlGenerator); ok {
			data.Generate()
			c.appendChild(data.GetObject())
		}
	}

	/*reflectValue := reflect.ValueOf(*c)
	for i := 0; i < reflectValue.NumField(); i++ {
		if reflectValue.Field(i).Kind() == reflect.Struct {
			c.subValueGenerate(reflectValue.Field(i))
		}
	}*/
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
			c.eventSelector(value.Field(i), event)
			//wasmhtml.SetAttribute(c.Object, event, value.Field(i).String())
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

func (c *CommonElement) eventSelector(value reflect.Value, event string) {
	switch v := value.Kind(); v {
	case reflect.String:
		if value.String() != "" {
			wasmhtml.SetAttribute(c.Object, event, value.String())
		}
	default:
		log.Println("Invalid type")
	}
}
