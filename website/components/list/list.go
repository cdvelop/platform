package list

import (
	"github.com/tinywasm/dom"
)

type List struct {
	*dom.Element
}

func (l *List) Render() *dom.Element {
	if l.Element != nil {
		return l.Element
	}

	container := dom.Div().Class("container-list-only")

	// Search input usually goes here (from search.css)
	// <div id="device-search-form" class="search-container">
	//     <label for="filter-object"><svg ...></svg></label>
	//     <input type="search" ...>
	// </div>

	searchForm := dom.Div().ID("device-search-form").Class("search-container")

	label := dom.Label().Attr("for", "filter-object")
	svg := dom.Svg().Class("form-btn").
		Attr("aria-hidden", "true").
		Attr("focusable", "false")
	use := dom.Use().Href("#icon-search")
	svg.Add(use)
	label.Add(svg)

	input := dom.Input("search").
		ID("filter-object").
		Title("letras números - permitidos max 50 caracteres").
		Attr("pattern", "^[A-Za-zÑñ 0-9-]{1,20}$")

	searchForm.Add(label, input)

	container.Add(searchForm)

	// List items would be added here

	l.Element = container
	return l.Element
}
