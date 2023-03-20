package cookie

import (
	"fmt"
	"strings"
	"syscall/js"

	"github.com/sudak-91/wasmhtml"
)

type Cookie struct {
	object js.Value
	data   map[string]string
}

type Cookies interface {
	GetCookies() (map[string]string, error)
}

func NewCookie() *Cookie {
	var c Cookie
	c.data = make(map[string]string)
	c.object = wasmhtml.Document.Get("cookie")
	return &c

}

func (c *Cookie) GetCookies() (map[string]string, error) {
	c.parseCookie()
	if c.data == nil {
		return nil, fmt.Errorf("Empty map")
	}
	return c.data, nil
}

func (c *Cookie) AddCoookie(key string, value string) {
	var data string
	switch {
	case c.object.String() == "":
		data = fmt.Sprintf("%s=%s", key, value)
	default:
		data = fmt.Sprintf("%s;%s=%s", c.object.String(), key, value)
	}
	wasmhtml.Document.Set("cookie", data)

}

func (c *Cookie) GetValue(key string) (string, error) {
	c.parseCookie()
	value, ok := c.data[key]
	if !ok {
		return "", fmt.Errorf("%s", "Not Value")
	}
	return value, nil
}

func (c *Cookie) parseCookie() {
	l := strings.Split(c.object.String(), ";")
	for _, value := range l {
		pairs := strings.Split(value, "=")
		if len(pairs) != 2 {
			return
		}
		c.data[pairs[0]] = pairs[1]
	}

}
