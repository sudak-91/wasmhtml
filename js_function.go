package wasmhtml

import (
	"fmt"
	"strconv"
	"syscall/js"
)

func QuerySelector(object js.Value, selector string) js.Value {
	return object.Call("querySelector", selector)
}

// Create element
func CreateElement(element string) js.Value {
	return Document.Call("createElement", element)
}

// Set HTML 'attribute' with  'value'
func Set(object js.Value, attribute string, value any) {
	object.Set(attribute, value)
}

// Set HTML attribute HIDDEN
func SetHidden(object js.Value) {
	Set(object, "hidden", "hidden")
}

// Set HTML attribute ID
func SetId(object js.Value, id string) {
	Set(object, "id", id)
}

// Set "class" attribute
func SetClass(object js.Value, class string) {
	Set(object, "class", class)
}

// Set HTML attribute TITLE
func SetTitle(object js.Value, title string) {
	Set(object, "title", title)
}

// Get HTML attribute TITLE
func GetTitile(object js.Value) string {
	return GetAttribute(object, "title").String()
}

// Set HTML attribute ALIGN
func SetAlign(object js.Value, align string) {
	Set(object, "align", align)
}

// Get HTML attribute ALIGN
func GetAlign(object js.Value) string {
	return GetAttribute(object, "align").String()
}

// Set HTML attribute ACCESSKEY
func SetAccesskey(object js.Value, key string) {
	SetAttribute(object, "accesskey", key)
}

// Get HTML attribute ACCESSKEY
func GetAccesskey(object js.Value) string {
	return GetAttribute(object, "accesskey").String()
}

// Remove HTML attribute ACCESSKEY
func RemoveAccesskey(object js.Value) {
	RemoveAttribute(object, "accesskey")
}

// Set HTML attribute CONTENTEDITABLE
func SetContentEditable(object js.Value, value string) {
	SetAttribute(object, "contenteditable", value)
}

// Get HTML attribute CONTENTEDITABLE
func GetContentEditable(object js.Value) (bool, error) {
	result := GetAttribute(object, "contenteditable").String()
	data, err := strconv.ParseBool(result)
	if err != nil {
		return false, err
	}
	return data, nil
}

// Remove HTML attribute CONTENTEDITABLE
func RemoveContentEditable(object js.Value) {
	RemoveAttribute(object, "contenteditable")
}

// Set HTML attribute DIR
func SetDir(object js.Value, dir string) {
	SetAttribute(object, "dir", dir)
}

// Get HTML attribute DIR
func GetDir(object js.Value) string {
	return GetAttribute(object, "dir").String()
}

// Remove HTML attribute DIR
func RemoveDir(object js.Value) {
	RemoveAttribute(object, "dir")
}

// Set HTML attribute LANG
func SetLang(object js.Value, langCode string) {
	SetAttribute(object, "lang", langCode)
}

// Get HTML attribute LANG
func GetLang(object js.Value) string {
	return GetAttribute(object, "lang").String()
}

// Remove HTML attribute LANG
func RemoveLang(object js.Value) {
	RemoveAttribute(object, "lang")
}

// Set HTML attribute SPELLCHECK
func SetSpellcheck(object js.Value, value string) {
	SetAttribute(object, "spellcheck", value)
}

func AddClickEventListenr(object js.Value, jsFunc js.Func) {
	AddEventListener(object, "click", jsFunc)
}

func RemoveClickEventListener(object js.Value) {
	RemoveEventListener(object, "click")
}

// AddEventListener
func AddEventListener(object js.Value, event string, jsFunc js.Func) {
	object.Call("addEventListener", event, jsFunc)
}

// RemoveEventListener
func RemoveEventListener(object js.Value, event string) {
	object.Call("removeEventListener", event)
}

// Get HTML attribute SPELLCHECK
func GetSpellcheck(object js.Value) (bool, error) {
	result := GetAttribute(object, "spellcheck")
	data, err := strconv.ParseBool(result.String())
	if err != nil {
		return false, err
	}
	return data, nil
}

// Remove HTML attribute SPELLCHECK
func RemoveSpellcheck(object js.Value) {
	RemoveAttribute(object, "spellcheck")
}

// Set HTML attribute STYLE
func SetStyle(object js.Value, value string) {
	SetAttribute(object, "style", value)
}

// Remove HTML attribute STYLE
func RemoveStyle(object js.Value) {
	RemoveAttribute(object, "style")
}

// Get HTML attribute STYLE
func GetStyle(object js.Value) string {
	return GetAttribute(object, "style").String()
}

// Set HTML attribute TABINDEX
func SetTabindex(object js.Value, index int32) {
	data := fmt.Sprintf("%d", index)
	SetAttribute(object, "tabindex", data)
}

// Get HTML attribute TABINDEX
func GetTabindex(object js.Value) int32 {
	result := GetAttribute(object, "tabindex").String()
	data, _ := strconv.ParseInt(result, 10, 32)
	return int32(data)
}

// Remove HTML attribute TABINDEX
func RemoveTabindex(object js.Value) {
	RemoveAttribute(object, "tabindex")
}

// Set HTML event ONBLUR
func SetOnblurEvent(object js.Value, funcName string) {
	SetAttribute(object, "onblur", funcName)
}

// Remove HTML event ONBLUR
func RemoveOnblurEvent(object js.Value) {
	RemoveAttribute(object, "onblur")
}

// Set HTML event ONCHANGE
func SetOnchangeEvent(object js.Value, funcName string) {
	SetAttribute(object, "onchange", funcName)
}

// Remove HTML event ONCHANGE
func RemoveOnchangeEvent(object js.Value) {
	RemoveAttribute(object, "onchange")
}

// Set HTML event ONCLICK
func SetOnclickEvent(object js.Value, funcName string) {
	SetAttribute(object, "onclick", funcName)
}

// Remove HTML event ONCLICK
func RemoveOnclickEvent(object js.Value) {
	RemoveAttribute(object, "onclick")
}

// Set HTML event ONDBCLICK
func SetOndbclickEvent(object js.Value, funcName string) {
	SetAttribute(object, "ondbclick", funcName)
}

// Remove HTML event ONDBCLICK
func RemoveOndbclickEvent(object js.Value) {
	RemoveAttribute(object, "ondbclick")
}

// Set HTML event ONFOCUS
func SetOnfocusEvent(object js.Value, funcName string) {
	SetAttribute(object, "onfocus", funcName)
}

// Remove HTML event ONFOCUS
func RemoveOnfocusEvent(object js.Value) {
	RemoveAttribute(object, "onfocus")
}

// Set HTML event ONKEYDOWN
func SetOnKeyDownEvent(object js.Value, funcName string) {
	SetAttribute(object, "onkeydown", funcName)
}

// Remove HTML event ONKEYDOWN
func RemoveOnKeyDownEvent(object js.Value) {
	RemoveAttribute(object, "onkeydown")
}

// Set HTML event ONKEYPRESS
func SetOnKeyPressEvent(object js.Value, funcName string) {
	SetAttribute(object, "onkeypress", funcName)
}

// Remove HTML event ONKEYPRESS
func RemoveOnKeyPressEvent(object js.Value) {
	RemoveAttribute(object, "onkeypress")
}

// Set HTML event onkeyup
func SetOnKeyUpEvent(object js.Value, funcName string) {
	SetAttribute(object, "onkeyup", funcName)
}

func RemoveOnKeyUpEvent(object js.Value) {
	RemoveAttribute(object, "onkeyup")
}

func RemoveChild(parent js.Value, child js.Value) {
	parent.Call("removeChild", child)
}

// HTML attribute VALUE
func GetValue(object js.Value) js.Value {
	return GetAttribute(object, "value")
}

func SetValue(object js.Value, value string) {
	SetAttribute(object, "value", value)
}

func InnerHtml(object js.Value, value string) {
	Set(object, "innerHTML", value)
}

func AppendChild(parent js.Value, child js.Value) {
	parent.Call("appendChild", child)
}
func AddClass(object js.Value, className string) {
	object.Get("classList").Call("add", className)
}
func GetAttribute(object js.Value, parametrName string) js.Value {
	return object.Get(parametrName)
}
func GetID(object js.Value) string {
	return GetAttribute(object, "id").String()
}

// Sets the value of an attribute on the specified element. If the attribute already exists, the value is updated; otherwise a new attribute is added with the specified name and value.
func SetAttribute(object js.Value, attributeName string, attributeValue any) {
	object.Call("setAttribute", attributeName, attributeValue)
}

func RemoveHidden(object js.Value) {
	RemoveAttribute(object, "hidden")
}

// Remove HTML attribute 'attributeName'
func RemoveAttribute(object js.Value, attributeName string) {
	object.Call("removeAttribute", attributeName)
}

func GetContext(object js.Value, contextType string) js.Value {
	context := object.Call("getContext", contextType)
	return context
}
