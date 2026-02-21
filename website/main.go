package main

import (
	"github.com/tinywasm/dom"
	"github.com/cdvelop/platform/website/components/layout"
	"github.com/cdvelop/platform/website/components/menu"
	"github.com/cdvelop/platform/website/components/login"
	"github.com/cdvelop/platform/website/components/module"
	"github.com/cdvelop/platform/website/components/form"
)

func main() {
	// Create Form Styles component (for CSS injection)
	formStyles := &form.FormStyles{}

	// Create menu
	m := &menu.Menu{
		Items: []menu.Item{
			{ModuleName: "home", Title: "Inicio", IconID: "icon-home"},
			{ModuleName: "users", Title: "Usuarios", IconID: "icon-key"},
		},
	}

	// Create login component
	l, err := login.New()
	if err != nil {
		dom.Log("Error creating login:", err)
		return
	}

	// Create module wrapper for content (e.g. login form)
	// Include formStyles to ensure form CSS is loaded
	content := dom.Div(
		formStyles,
		l,
	)

	mod := &module.Module{
		Title:   "Identificación",
		Content: content,
	}

	// Create Layout
	app := &layout.Layout{
		UserName: "Invitado",
		UserArea: "Público",
		Message:  "",
		Menu:     m,
		Content:  mod,
	}

	dom.Render("body", app)

	// Keep the app running
	select {}
}
