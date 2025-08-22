package platform

func (t Theme) loginTemplate(c *TemplateModuleConfig) (out string) {

	out += `<div class="all-space-centered">
<div class="cont-centered-btn">`
	// <div class="cont-centered-btn">`
	out += `<form name="` + c.Form.ObjectName() + `" autocomplete="off" spellcheck="false">`
	out += t.renderFields(c.Form)
	out += `</form>`
	out += `<div class="login-button">`

	for _, button := range c.FormButtons {
		out += BuildHtmlFormButton(button)
	}

	out += `</div>`

	out += `</div>
</div>`

	return
}
