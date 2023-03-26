package main

import (
	"log"
	"syscall/js"

	"github.com/sudak-91/wasmhtml"
	"github.com/sudak-91/wasmhtml/element"
)

type MainPage struct {
	MainDiv     *element.Div
	ButtonPanel *element.Div
	Page1Button *element.Div
	Page2Button *element.Div
	Page3Button *element.Div
	View        *element.Div
}

func NewMainPage() *MainPage {
	var mainPage MainPage
	body := element.GetBody()
	mainPage.MainDiv = body.AddDiv()
	mainPage.MainDiv.SetID("main")
	mainPage.ButtonPanel = mainPage.MainDiv.AddDiv()
	mainPage.ButtonPanel.SetID("panel")
	mainPage.Page1Button = mainPage.ButtonPanel.AddDiv()
	mainPage.Page1Button.SetID("window1")
	mainPage.Page1Button.SetInnerHtml("Первый экран")
	mainPage.Page2Button = mainPage.ButtonPanel.AddDiv()
	mainPage.Page2Button.SetID("window2")
	mainPage.Page2Button.SetInnerHtml("Второй экран")
	mainPage.Page3Button = mainPage.ButtonPanel.AddDiv()
	mainPage.Page3Button.SetID("window3")
	mainPage.Page3Button.SetInnerHtml("Третий экран")
	mainPage.View = mainPage.MainDiv.AddDiv()
	mainPage.View.SetID("view")
	return &mainPage
}

func (m *MainPage) Render() {
	element.GetBody().Generate()
}

func main() {
	MainPage := NewMainPage()
	MainPage.Render()
	wasmhtml.AddClickEventListenr(MainPage.Page1Button.Object, js.FuncOf(func(this js.Value, args []js.Value) any {
		log.Println(len(MainPage.View.Child))
		for _, v := range MainPage.View.Child {
			log.Println("cycle")
			if data, ok := v.(element.HtmlGenerator); ok {
				log.Println("Delete elementh")
				MainPage.View.Object.Call("removeChild", data.GetObject())
			}
		}
		MainPage.View.Child = nil
		NewView := MainPage.View.AddDiv()
		NewView.SetID("view1")
		NewView.SetInnerHtml("View 1")
		MainPage.View.Generate()

		return nil
	}))
	wasmhtml.AddClickEventListenr(MainPage.Page2Button.Object, js.FuncOf(func(this js.Value, args []js.Value) any {
		NewView := MainPage.View.AddDiv()
		NewView.SetID("view2")
		NewView.SetInnerHtml("View 2")
		MainPage.View.Generate()
		return nil
	}))
	//wasmhtml.Document.Call("addEventListener", "click", js.FuncOf(TestFunc))
	//cookie := cookie.NewCookie()
	//cookie.AddCoookie("test", "success")
	//cookie.AddCoookie("test2", "eahhh")
	//data, err := cookie.GetCookies()
	//if err != nil {
	//	panic(err)
	//}
	//for k, v := range data {
	//	log.Printf("%s is %s", k, v)
	//}
	c := make(chan bool)
	<-c

}
func TestFunc(this js.Value, args []js.Value) any {
	log.Println(args[0])
	log.Println("tada")
	return nil
}
