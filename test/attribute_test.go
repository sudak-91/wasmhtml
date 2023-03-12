package common_attribute_test

import (
	"log"
	"reflect"
	"testing"

	"github.com/sudal-91/wasmhtml/attribute"
)

type InjectBaseAttrinute struct {
	base attribute.BaseAttribute
}

func TestCommonAttrinute(t *testing.T) {
	var data attribute.BaseAttribute
	getReflectData(data)
}

func TestInjectBaseAttribute(t *testing.T) {
	var data InjectBaseAttrinute
	data.base.Class = "SisskY"
	getReflectData(data)
}

func reflectGenerator(reftype reflect.Type, refValue reflect.Value) {
	for i := 0; i < reftype.NumField(); i++ {
		field := reftype.Field(i)
		if refValue.Field(i).Kind() == reflect.Struct {
			refSubtype := refValue.Field(i).Type()
			for j := 0; j < refSubtype.NumField(); j++ {
				f := refSubtype.Field(j).Tag
				log.Println(f)
				log.Println(refValue.Field(i).Field(j).String())
			}
		}
		log.Println(field)
	}

}

func getReflectData(data any) {
	dataType := reflect.TypeOf(data)
	dataValue := reflect.ValueOf(data)
	reflectGenerator(dataType, dataValue)
}
