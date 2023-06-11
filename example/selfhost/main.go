package main

import (
	"fmt"
	"log"
	"syscall/js"

	"github.com/sudak-91/wasmhtml"
	"github.com/sudak-91/wasmhtml/element"
)

type MainPage struct {
	Body          *element.Body
	MainDiv       *element.Div
	Navigation    *element.Div
	Container     *element.Div
	Canvas        *element.Div
	Elem          js.Value
	ActiveElement string
}

func NewMainPage() *MainPage {
	var mainPage MainPage
	mainPage.Body = element.GetBody()
	mainPage.MainDiv = mainPage.Body.AddDiv()
	mainPage.MainDiv.SetID("main")
	mainPage.MainDiv.SetStyle(`
		display: flex;
		flex-direction: row;
		align-items: flex-start;
		padding: 10px;
		gap: 10px;
		position: relative;
		width: 1306px;
		height: 836px;
		background: #DFDFDF;
	`)
	mainPage.Navigation = mainPage.MainDiv.AddDiv()
	mainPage.Navigation.AddClass("nav")
	mainPage.Navigation.SetStyle(`
	display: flex;
	flex-direction: column;
	align-items: flex-start;
	padding: 10px;
	gap: 10px;
	
	width: 320px;
	height: 816px;
	
	background: #D0CFCF;
	
	/* Inside auto layout */
	
	flex: none;
	order: 0;
	align-self: stretch;
	flex-grow: 0;
	`)
	CircleButton := mainPage.Navigation.AddDiv()
	CircleButton.AddClass("btn")
	CircleButton.SetID("circle")
	CircleButton.SetStyle(`
		box-sizing: border-box;
		width: 100px;
		height: 100px;
		border: 1px solid #000000;
		border-radius: 10px;
		flex: none;
		order: 0;
		flex-grow: 0;
		cursor:move;
	`)
	wasmhtml.SetAttribute(CircleButton.Object, "draggable", "true")
	wasmhtml.AddClickEventListenr(CircleButton.Object, js.FuncOf(mainPage.SelectElement))
	wasmhtml.AddEventListener(CircleButton.Object, "dragstart", js.FuncOf(mainPage.DragStart))
	//wasmhtml.AddEventListener(CircleButton.Object, "dragend", js.FuncOf(mainPage.DragStop))
	mainPage.Container = mainPage.MainDiv.AddDiv()
	mainPage.Container.AddClass("cont")
	mainPage.Container.SetStyle(`
		background: #D2D2D2;
		flex: none;
		order: 1;
		align-self: stretch;
		flex-grow: 1;
		width: 952px;
		height: 816px;
`)
	mainPage.Canvas = mainPage.Container.AddDiv()
	mainPage.Canvas.SetStyle(`
	width: 500px;
	height: 500px;
	background: white;
	`)
	mainPage.Canvas.AddClass("frame")
	mainPage.Canvas.SetID("canvasFrame")
	//wasmhtml.AddClickEventListenr(mainPage.Canvas.Object, js.FuncOf(mainPage.CreateElement))
	wasmhtml.AddEventListener(mainPage.Canvas.Object, "dragover", js.FuncOf(mainPage.DragOver))
	wasmhtml.AddEventListener(mainPage.Canvas.Object, "drop", js.FuncOf(mainPage.Drop))

	return &mainPage
}

func (m *MainPage) Render() {
	m.Body.Generate()
}

func main() {
	MainPage := NewMainPage()
	MainPage.Render()
	/*wasmhtml.AddClickEventListenr(MainPage.Page1Button.Object, js.FuncOf(func(this js.Value, args []js.Value) any {
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
	//}*/
	c := make(chan bool)
	<-c

}
func TestFunc(this js.Value, args []js.Value) any {
	log.Println(args[0])
	log.Println("tada")
	return nil
}

func (m *MainPage) SelectElement(this js.Value, args []js.Value) any {
	m.ActiveElement = wasmhtml.GetID(this)
	log.Println(m.ActiveElement)
	return nil
}

func (m *MainPage) DragStart(this js.Value, args []js.Value) any {
	log.Println("drag start")
	m.Elem = this
	return nil
}

func (m *MainPage) DragStop(this js.Value, args []js.Value) any {
	log.Println("drag stop")
	return nil
}

func (m *MainPage) DragOver(this js.Value, args []js.Value) any {
	args[0].Call("preventDefault", "")
	log.Println("drag over")
	return nil
}

func (m *MainPage) Drop(this js.Value, args []js.Value) any {
	//args[0].Call("preventDefault", "")
	l := this.Get("id").String()
	log.Println(l)

	el := m.Elem.Call("cloneNode", "true")
	x := args[0].Get("clientX").Int()
	y := args[0].Get("clientY").Int()
	Leftpos := fmt.Sprintf(`position: absolute;
	 left: %d;
	  top:%d;
	  box-sizing: border-box;
		width: 100px;
		height: 100px;
		border: 1px solid #000000;
		border-radius: 10px;
		flex: none;
		order: 0;
		flex-grow: 0;
		cursor:move;
	  `, x, y)
	wasmhtml.SetStyle(el, Leftpos)
	wasmhtml.AppendChild(this, el)
	m.Elem = js.Value{}
	return nil
}

func (m MainPage) CreateElement(this js.Value, args []js.Value) any {
	for k, v := range args {
		log.Printf("Key: %d, Value: %s", k, v.String())
	}
	log.Println(args[0].Get("clientX"))
	log.Println(args[0].Get("clientY"))
	log.Println(this.Get("offsetLeft"))
	log.Println(this.Call("getBoundingClientRect", "left"))
	return nil
}
