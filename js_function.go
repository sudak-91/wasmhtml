package wasmhtml

import "syscall/js"

func QuerySelector(object js.Value, selector string) js.Value {
	return object.Call("querySelector", selector)
}

// Create element
func CreateElement(element string) js.Value {
	return Document.Call("createElement", element)
}

func Set(object js.Value, attribute string, value any) {
	object.Set(attribute, value)
}

// Set Id attribute
func SetId(object js.Value, id string) {
	Set(object, "id", id)
}

// Set "class" attribute
func SetClass(object js.Value, class string) {
	Set(object, "class", class)
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

// Sets the value of an attribute on the specified element. If the attribute already exists, the value is updated; otherwise a new attribute is added with the specified name and value.
func SetAttribute(object js.Value, attributeName string, attributeValue string) {
	object.Call("setAttribute", attributeName, attributeValue)
}
