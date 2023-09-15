package platform

import "github.com/cdvelop/model"

func (Theme) ModuleClassName() string {
	return "slider_panel"
}

func (t Theme) ModuleTemplate(m *model.Module, form_object *model.Object, options ...string) string {

	if m == nil {
		return "error al generar template modulo no ingresado"
	}

	out := `<div class="header-module">
	<div class="module-title">
		<h1>` + m.Title + `:</h1>
	</div>`

	out += m.HeaderInputTarget

	out += `</div>`

	out += `<div class="scroll-container">

	<div class="box-snap search-list-container">

		<div class="container-list-search">
		


		</div>

		<div id="device-search-form" class="search-container">
			<label for="search_device">
				<svg aria-hidden="true" focusable="false" class="form-btn">
					<use xlink:href="#icon-search" />
				</svg>
			</label>
			<input type="search" id="search-footer-device" title="letras numeros - permitidos max 50 caracteres"
				pattern="^[A-Za-zÑñ 0-9-]{1,20}$">
		</div>

	</div>

	<div class="box-snap edition-container">

		<div class="form-inputs-container">
			<div class="form-distributed-fields">`

	out += t.buildHtmlForm(form_object)

	out += `</div>
		</div>`

	out += buttons()

	out += `</div>
</div>`

	return out
}
