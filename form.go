package platform

import (
	"strconv"

	"github.com/cdvelop/model"
)

func formTemplate(object_name, field_with_id, input_tags string) string {
	return `<form class="form-distributed-fields" name="` + object_name + `" data-field_id="` + field_with_id + `" autocomplete="off">
	` + input_tags + `
	</form>`
}

// object_name ej: user,staff,product
// field_name ej: id_user,name,address
// legend ej: Usuario, Personal, Productos
// html_name ej: text, checkbox, textarea
func (t Theme) inputTemplate(object_name, field_name, legend, html_name, input_tag string, index int) string {

	st_index := strconv.Itoa(index)

	id := inputIdTemplate(object_name, field_name, st_index)
	var out string
	if html_name == "file" {
		return input_tag
	}

	out = `<fieldset data-name="` + field_name + `" tabindex="` + st_index + `"` + cssClass(html_name) + ` onclick="internalFocus(this)">
	<legend class="basic-legend"><label for="` + id + `">` + legend + `</label></legend>
	` + input_tag + `
    </fieldset>`

	return out
}

func inputIdTemplate(object_name, field_name, index string) string {
	return object_name + `.` + field_name + `.` + index
}

func cssClass(html_name string) string {
	var class = `normal border`

	// fmt.Println("HTML NAME:", html_name)

	switch html_name {

	case "checkbox", "textarea":
		class += ` all-width`

	// case "file":
	// 	class = `file`

	default:
		class += ` width-auto`
	}

	return ` class="` + class + `"`
}

// func (t Theme) JsFormVariablesTemplate(object_name string) string {
// 	return `const form = module.querySelector('#` + object_name + `-form');
// 	const fieldset = module.querySelectorAll("#` + object_name + `-form fieldset");`
// }

// func (Theme) JsInputVariableTemplate(field_name, html_name string) string {

// 	query_type := `querySelector`

// 	switch html_name {
// 	case "radio":
// 		query_type = `querySelectorAll`
// 	case "checkbox":
// 		query_type = `querySelectorAll`
// 	}

// 	return `let input_` + field_name + ` = form.` + query_type + `("[name='` + field_name + `']");\n`

// }

func (t Theme) buildHtmlForm(o *model.Object) string {

	if o != nil {

		if o.Module != nil && len(o.Fields) != 0 {
			var input_tags string

			for index, f := range o.RenderFields() {

				id := inputIdTemplate(o.ObjectName, f.Name, strconv.Itoa(index))

				tag := f.Input.BuildContainerView(id, f.Name, f.SkipCompletionAllowed)

				if f.Input.HtmlName() != "hidden" {

					input_tags += t.inputTemplate(o.ObjectName, f.Name, f.Legend, f.Input.HtmlName(), tag, index) + "\n"

				} else {

					input_tags += tag
				}

			}

			return formTemplate(o.ObjectName, o.PrimaryKeyName(), input_tags)
		}

	}

	return ""
}
