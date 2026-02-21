package module

import (
	"github.com/tinywasm/dom"
)

type Module struct {
	*dom.Element
	Title   string
	Content dom.Component
}

func (m *Module) Render() *dom.Element {
	if m.Element == nil {
		m.Element = &dom.Element{}
	}

	// Wrapper with class for layout grid area
	wrapper := dom.Div().Class("slider_panel").ID("module-content")

	// Header
	header := dom.Div().Class("header-module")
	titleDiv := dom.Div().Class("module-title")
	titleDiv.Add(dom.H1(m.Title))
	header.Add(titleDiv)

	wrapper.Add(header)

	// Content area
	scrollContainer := dom.Div().Class("scroll-container")
	if m.Content != nil {
		scrollContainer.Add(m.Content)
	}

	wrapper.Add(scrollContainer)

	return wrapper
}
