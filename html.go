package platform

func (Theme) MenuButtonTemplate(module_name, index, icon_id, title string) string {
	return `<li class="navbar-item"><a href="#` + module_name + `" tabindex="` + index + `" class="navbar-link" name="` + module_name + `">
	<svg aria-hidden="true" focusable="false" class="fa-primary"><use xlink:href="#` + icon_id + `" /></svg>
	<span class="link-text">` + title + `</span></a></li>`
}

func (Theme) ModuleHtmlTemplate(module_name, content string) string {
	// nombre del module html y el contenido
	return `<div id="` + module_name + `" class="slider_panel">` + content + `</div>`
}

func (Theme) IconModuleTemplate(id, view_box string, paths ...string) string {
	out := `<symbol id="` + id + `" viewBox="` + view_box + `">`
	for _, d := range paths {
		out += `<path fill="currentColor" d="` + d + `" />`
	}
	out += `</symbol>`
	return out
}
