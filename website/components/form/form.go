package form

import "github.com/tinywasm/dom"

type FormStyles struct {
	*dom.Element
}

func (f *FormStyles) Render() *dom.Element {
	if f.Element == nil {
		f.Element = dom.Span().Attr("style", "display: none;")
	}
	return f.Element
}
