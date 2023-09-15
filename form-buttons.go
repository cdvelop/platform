package platform

func buttons() string {

	out := `<div class="crud-buttons-container">
	<div class="contebuton">`

	out += `  <button name="btn_delete" title="Eliminar" disabled="true">
	<svg aria-hidden="true" focusable="false" class="form-btn">
		<use xlink:href="#icon-btn-delete" />
	</svg>
</button>`

	out += `<button name="btn_cancel" title="Cancelar" disabled="true">
	<svg aria-hidden="true" focusable="false" class="form-btn">
		<use xlink:href="#icon-btn-cancel" />
	</svg>
</button>`

	out += `<button name="btn_add" title="Iniciar Captura de Imagen">
	<svg aria-hidden="true" focusable="false" class="form-btn">
	<use xlink:href="#icon-btn-new" />
	</svg>
	</button>`

	out += `<!-- <button name="btn_save" title="Tomar Fotografia" disabled>
	<svg aria-hidden="true" focusable="false" class="form-btn">
		<use xlink:href="#icon-camera" />
	</svg>
</button> -->`

	out += `</div>
</div>`

	return out
}
