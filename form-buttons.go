package platform

func buttons(c *TemplateModuleConfig) (out string) {

	if c != nil && len(c.FormButtons) != 0 {
		out = `<div class="crud-buttons-container"><div class="contebuton">`

		// out += BuildHtmlFormButton(ButtonForm{
		// 	ModuleName: c.Module.ModuleName,
		// 	ButtonName: "btn_cancel",
		// 	Title:      "Comenzar Nuevamente",
		// 	IconID:     "icon-btn-cancel",
		// 	OnclickFun: "resetModule(this)",
		// })

		for _, btn := range c.FormButtons {

			out += BuildHtmlFormButton(btn)

		}

		// 	out += `<button type="button" name="btn_save" title="Tomar Fotografía" disabled>
		// 	<svg aria-hidden="true" focusable="false" class="form-btn">
		// 		<use xlink:href="#icon-camera" />
		// 	</svg>
		// </button>`

		out += `</div></div>`

	}

	return out
}

func BuildHtmlFormButton(b *ButtonForm) string {

	var disabled string
	if b.Disabled {
		disabled = ` disabled`
	}

	var module string
	if b.ModuleName != "" {
		module = ` data-module="` + b.ModuleName + `"`
	}

	var object string
	if b.ObjectName != "" {
		object = ` data-object_name="` + b.ObjectName + `"`
	}

	return `<button type="button"` + module + object + ` name="` + b.ButtonName + `" title="` + b.Title + `" onclick="` + b.OnclickFun + `"` + disabled + `>
		<svg aria-hidden="true" focusable="false" class="form-btn">
		<use xlink:href="#` + b.IconID + `" />
		</svg>
	</button>`
}
