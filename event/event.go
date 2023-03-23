package event

type BaseEvent struct {
	//OnBlur      string `htmlevnt:"onblur"`
	//OnChange    string `htmlevnt:"onchange"`
	//OnClick     string `htmlevnt:"onclick"`
	//OnDbClick   string `htmlevnt:"ondbclick"`
	OnFocus     string `htmlevnt:"onfocus"`
	OnKeyDown   string `htmlevnt:"onkeydown"`
	OnKeyPress  string `htmlevnt:"onkeypress"`
	OnKeyUp     string `htmlevnt:"onkeyup"`
	OnLoad      string `htmlevnt:"onload"`
	OnMouseDown string `htmlevnt:"onmousedown"`
	OnMouseMove string `htmlevnt:"onmousemove"`
	OnMouseOut  string `htmlevnt:"onmouseout"`
	OnMouseOver string `htmlevnt:"onmouseover"`
	OnMouseUp   string `htmlevnt:"onmouseup"`
	OnReset     string `htmlevnt:"onreset"`
	OnSelect    string `htmlevnt:"onselect"`
	OnSubmit    string `htmlevnt:"onsubmit"`
	OnUnload    string `htmlevnt:"onunload"`
}
