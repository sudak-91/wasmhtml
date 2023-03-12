# wasmhtml
Библиотека, предоставляющая инструментарий для работы с HTML элементами и дпльнейшей сборки в wasm файл

## Начало работы
Для начала работы требуется установить пакет
`go get github.com/sudak-91/wasmhtml`

## Создание страницы
```
	body := element.GetBody()     //Получение <BODY>
	leftdiv := body.AddDiv()      //Создание дочернего <DIV>
	leftdiv.Id = "lefyDiv"        //Добавление ID
	leftdiv.AddClass("container") //Добавление клааса
	body.Generatr() 			  //Генерация страницы
```
