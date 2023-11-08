package platform

func buttons(module_name, object_name string) string {

	out := `<div class="crud-buttons-container">
	<div class="contebuton">`

	out += `<button type="button" data-module="` + module_name + `" name="btn_cancel" title="Comenzar Nuevamente" onclick="resetModule(this)">
	<svg aria-hidden="true" focusable="false" class="form-btn">
		<use xlink:href="#icon-btn-cancel" />
	</svg>
	</button>`

	out += `  <button type="button" data-object_name="` + object_name + `" name="btn_printer" title="Imprimir" onclick="printForm(this)">
	<svg aria-hidden="true" focusable="false" class="form-btn">
		<use xlink:href="#icon-printer" />
	</svg>
</button>`

	// 	out += `<button type="button" name="btn_save" title="Tomar Fotografía" disabled>
	// 	<svg aria-hidden="true" focusable="false" class="form-btn">
	// 		<use xlink:href="#icon-camera" />
	// 	</svg>
	// </button>`

	out += `</div>
</div>`

	return out
}
