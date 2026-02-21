package login

import (
	"github.com/tinywasm/dom"
	"github.com/tinywasm/form"
	"github.com/cdvelop/platform/website/models"
)

type Login struct {
	*dom.Element
	Form *form.Form
}

func New() (*Login, error) {
	// Create form with specific parent ID.
	f, err := form.New("login-area", &models.Login{})
	if err != nil {
		return nil, err
	}

	l := &Login{
		Form: f,
	}

	f.OnSubmit(func(data any) error {
		// Example login logic
		dom.Log("Submitting login form...")
		return nil
	})

	return l, nil
}

func (l *Login) Render() *dom.Element {
	if l.Element != nil {
		return l.Element
	}

	container := dom.Div().Class("all-space-centered")

	center := dom.Div().Class("cont-centered-btn")

	center.Add(l.Form)

	btnDiv := dom.Div().Class("login-button")

	formID := l.Form.GetID()

	btn := dom.Button().
		Attr("type", "submit").
		Attr("form", formID).
		Attr("name", "btn_login").
		Attr("title", "Ingresar")

	svg := dom.Svg().Class("form-btn").
		Attr("aria-hidden", "true").
		Attr("focusable", "false")

	use := dom.Use().Attr("xlink:href", "#icon-key")
	svg.Add(use)

	btn.Add(svg)
	btnDiv.Add(btn)

	center.Add(btnDiv)

	container.Add(center)

	l.Element = container
	return l.Element
}
