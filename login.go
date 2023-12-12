package platform

import "github.com/cdvelop/model"

func (t Theme) loginTemplate(o *model.Object) (out string) {

	out += `<div class="all-space-centered">
<div class="cont-centered-btn">`
	// <div class="cont-centered-btn">`
	out += `<form name="` + o.ObjectName + `" autocomplete="off" spellcheck="false">`
	out += t.renderFields(o)
	out += `</form>`
	out += `<div class="login-button">`

	out += BuildHtmlFormButton(ButtonForm{
		ObjectName: o.ObjectName,
		ButtonName: "btn_" + o.ObjectName,
		Title:      o.Title,
		IconID:     "icon-key",
		OnclickFun: "submitLoginForm(this)",
	})
	out += `</div>`

	out += `</div>
</div>`

	return
}
