package button

import (
	"github.com/tinywasm/dom"
)

type Button struct {
	*dom.Element
	Name     string
	Title    string
	IconID   string
	Type     string
	Disabled bool
	OnClick  func(dom.Event)
	FormID   string
}

func (b *Button) Render() *dom.Element {
	btn := dom.Button()

	if b.Type != "" {
		btn.Attr("type", b.Type)
	} else {
		btn.Attr("type", "button")
	}

	if b.Name != "" {
		btn.Attr("name", b.Name)
	}

	if b.Title != "" {
		btn.Title(b.Title)
	}

	if b.FormID != "" {
		btn.Attr("form", b.FormID)
	}

	if b.Disabled {
		btn.Attr("disabled", "")
	}

	if b.OnClick != nil {
		btn.OnClick(b.OnClick)
	}

	svg := dom.Svg().Class("form-btn").
		Attr("aria-hidden", "true").
		Attr("focusable", "false")

	if b.IconID != "" {
		use := dom.Use().Href("#" + b.IconID)
		svg.Add(use)
	}

	btn.Add(svg)

	b.Element = btn
	return b.Element
}
