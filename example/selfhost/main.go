package main

import (
	"log"

	"github.com/sudak-91/wasmhtml/cookie"
	"github.com/sudak-91/wasmhtml/element"
)

func main() {
	body := element.GetBody()     //Получение <BODY>
	leftdiv := body.AddDiv()      //Создание дочернего <DIV>
	leftdiv.Id = "lefyDiv"        //Добавление ID
	leftdiv.AddClass("container") //Добавление клааса
	rightdiv := body.AddDiv()
	rightdiv.Id = "rightDiv"
	rightdiv.InnerHtml = "Click me"
	rightdiv.OnClick = "alert(\"tada\")"
	rightdiv.AddClass("container")
	body.Generate() //Генерация старницы
	cookie := cookie.NewCookie()
	cookie.AddCoookie("test", "success")
	cookie.AddCoookie("test2", "eahhh")
	data, err := cookie.GetCookies()
	if err != nil {
		panic(err)
	}
	for k, v := range data {
		log.Printf("%s is %s", k, v)
	}
	c := make(chan bool)
	<-c
}
