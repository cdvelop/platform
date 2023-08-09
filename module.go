package platform

func (Theme) ModuleHtmlTemplate(module_name, content string) string {
	// nombre del module html y el contenido
	return `<div id="` + module_name + `" class="slider_panel">` + content + `</div>`
}
