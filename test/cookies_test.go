package common_attribute_test

import (
	"log"
	"strings"
	"testing"
)

func TestCookieParse(t *testing.T) {
	k := "a=111;b=222;c=333"
	var cookies map[string]string
	cookies = make(map[string]string)
	l := strings.Split(k, ";")
	for _, value := range l {
		pairs := strings.Split(value, "=")
		cookies[pairs[0]] = pairs[1]
	}
	if cookies["a"] != "111" {
		t.Fail()
	}
	log.Println(cookies)
}
