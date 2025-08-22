package platform

func (Theme) ModuleClassName() string {
	return "slider_panel"
}

func (t Theme) ModuleTemplate(c *TemplateModuleConfig) string {

	if c.Module == nil {
		return "error al generar template modulo no ingresado"
	}

	// HEADER -->
	out := `<div class="header-module">
	<div class="module-title">
		<h1>` + c.Module.Title() + `:</h1>
	</div>`
	out += c.HeaderInputTarget
	out += `</div>`
	// <-- HEADER

	// CONTAINER -->
	out += `<div class="scroll-container">`

	if c.RenderAllSpaceCentered {

		out += t.loginTemplate(c)

	} else {

		out += `<div class="box-snap search-list-container">`

		if c.AsideList != nil {
			out += c.AsideList.BuildContainerView("", "", true)
		} else {
			out += `<div class="container-list-only"></div>`
		}

		out += `</div>`

		out += `<div class="box-snap edition-container">`

		if len(c.FormButtons) != 0 {
			out += `<div class="form-container foot-buttons-border">`
		} else {
			out += `<div class="form-container full-form-border">`
		}

		out += `<div class="form-distributed-fields">`

		out += formTemplate(c.Form.ObjectName(), c.Form.PrimaryKeyName(), t.renderFields(c.Form))

		out += `</div></div>`

		out += buttons(c)

		out += `</div>`
	}

	out += `</div>`
	// <-- CONTAINER

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
