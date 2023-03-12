package main

import (
	"github.com/sudal-91/wasmhtml/element"
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

}
