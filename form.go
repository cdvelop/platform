package platform

import "strconv"

func (Theme) FormTemplate(object_name, input_tags string) string {
	return `<form class="form-distributed-fields" id="` + object_name + `-form" name="` + object_name + `-form" autocomplete="off">
	` + input_tags + `
	</form>`
}

// object_name ej: user,staff,product
// field_name ej: id_user,name,address
// legend ej: Usuario, Personal, Productos
// html_name ej: text, checkbox, textarea
func (t Theme) InputTemplate(object_name, field_name, legend, html_name, input_tag string, index int) string {

	st_index := strconv.Itoa(index)

	id := t.InputIdTemplate(object_name, field_name, st_index)

	return `<fieldset data-name="` + field_name + `" tabindex="` + st_index + `"` + cssClass(html_name) + `>
	<legend><label for="` + id + `">` + legend + `</label></legend>
	` + input_tag + `
    </fieldset>`
}

func (Theme) InputIdTemplate(object_name, field_name, index string) string {
	return `form.` + object_name + `.` + field_name + `.` + index
}

func cssClass(html_name string) string {
	var css string

	switch html_name {

	case "checkbox":
		css = "min-width-100"

	case "textarea":
		css = "min-width-100"

	default:
		css = "max-height-100"

	}

	return ` class="` + css + `"`
}

func (t Theme) JsFormVariablesTemplate(object_name string) string {
	return `const form = module.querySelector('#` + object_name + `-form');
	const fieldset = module.querySelectorAll("#` + object_name + `-form fieldset");`
}

func (Theme) JsInputVariableTemplate(field_name, html_name string) string {

	query_type := `querySelector`

	switch html_name {
	case "radio":
		query_type = `querySelectorAll`
	case "checkbox":
		query_type = `querySelectorAll`
	}

	return `let input_` + field_name + ` = form.` + query_type + `("[name='` + field_name + `']");\n`

}
