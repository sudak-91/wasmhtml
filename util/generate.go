package util

import (
	"log"
	"reflect"
	"syscall/js"

	"github.com/sudak-91/wasmhtml"
)

func Generate(reflectValue reflect.Value, d js.Value) {
	for i := 0; i < reflectValue.NumField(); i++ {
		log.Printf("%+v", reflectValue.Field(i).Kind())
		if reflectValue.Field(i).Kind() == reflect.Struct {
			SubValueGenerate(d, reflectValue.Field(i))
			continue
		}
	}
}

func SubValueGenerate(c js.Value, value reflect.Value) {
	subType := value.Type()
	for i := 0; i < subType.NumField(); i++ {
		tag, ok := subType.Field(i).Tag.Lookup("htmlattr")
		if ok {
			typeSelector(c, value.Field(i), tag)
		}
		event, ok := subType.Field(i).Tag.Lookup("htmlevnt")
		if ok {
			eventSelector(c, value.Field(i), event)
		}
		continue
	}
}

func typeSelector(c js.Value, value reflect.Value, tag string) {
	switch v := value.Kind(); v {
	case reflect.String:
		log.Println(tag)
		log.Println(value.String())
		if value.String() != "" {
			wasmhtml.Set(c, tag, value.String())
		}
	case reflect.Bool:
		if value.Bool() {
			wasmhtml.Set(c, tag, "")
		}
	default:
		log.Println("Invalid data type")
	}
}

func eventSelector(c js.Value, value reflect.Value, event string) {
	switch v := value.Kind(); v {
	case reflect.String:
		if value.String() != "" {
			wasmhtml.SetAttribute(c, event, value.String())
		}
	default:
		log.Println("Invalid type")
	}
}
