package menu

import (
	"github.com/tinywasm/dom"
	"strconv"
)

type Item struct {
	ModuleName string
	Title      string
	IconID     string
}

type Menu struct {
	*dom.Element
	Items []Item
}

func (m *Menu) Render() *dom.Element {
	if m.Element == nil {
		m.Element = &dom.Element{}
	}

	ul := dom.Ul().Class("navbar-container")

	for i, item := range m.Items {
		index := strconv.Itoa(i)

		linkID := "menu-link-" + item.ModuleName

		link := dom.A("").
			ID(linkID).
			Attr("href", "#" + item.ModuleName).
			Attr("tabindex", index).
			Class("navbar-link").
			Attr("name", item.ModuleName)

		svg := dom.Svg().
			Attr("aria-hidden", "true").
			Attr("focusable", "false").
			Class("fa-primary")

		use := dom.Use().Attr("xlink:href", "#" + item.IconID)
		svg.Add(use)

		span := dom.Span(item.Title).Class("link-text")

		link.Add(svg, span)

		li := dom.Li().Class("navbar-item").Add(link)
		ul.Add(li)
	}

	nav := dom.Nav().Class("menu-container").Add(ul)

	return nav
}
