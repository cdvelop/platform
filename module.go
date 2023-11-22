package platform

import "github.com/cdvelop/model"

func (Theme) ModuleClassName() string {
	return "slider_panel"
}

func (t Theme) ModuleTemplate(m *model.Module, form_object *model.Object, list model.ContainerViewAdapter) string {

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

	<div class="box-snap search-list-container">`

	if list != nil {
		out += list.BuildContainerView("", "", true)
	} else {
		out += `<div class="container-list-only"></div>`
	}

	out += `</div>

	<div class="box-snap edition-container">

		<div class="form-inputs-container">
			<div class="form-distributed-fields">`

	out += t.buildHtmlForm(form_object)

	out += `</div></div>`

	out += buttons(m.ModuleName, form_object.ObjectName)

	out += `</div>
</div>`

	return out
}

// `<div class="container-list-search">

// 		</div>

// 		<div id="device-search-form" class="search-container">
// 			<label for="filter-object">
// 				<svg aria-hidden="true" focusable="false" class="form-btn">
// 					<use xlink:href="#icon-search" />
// 				</svg>
// 			</label>
// 			<input type="search" id="filter-object" title="letras números - permitidos max 50 caracteres"
// 				pattern="^[A-Za-zÑñ 0-9-]{1,20}$">
// 		</div>`
