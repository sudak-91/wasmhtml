package attribute

type BaseAttribute struct {
	Accesskey       string `htmlattr:"accesskey"`
	Contentediteble string `htmlattr:"contetediteble"`
	Dir             string `htmlattr:"dir"`
	Hidden          bool   `htmlattr:"hidden"`
	Id              string `htmlattr:"id"`
	Lang            string `htmlattr:"lang"`
	Spellcheck      string `htmlattr:"spellcheck"`
	Style           string `htmlattr:"style"`
	Tabindex        string `htmlattr:"tabindex"`
	Title           string `htmlattr:"title"`
	XmlLang         string `htmlattr:"xml:langd"`
}
