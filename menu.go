package platform

func (Theme) MenuButtonTemplate(module_name, index, icon_id, title string) string {
	return `<li class="navbar-item"><a href="#` + module_name + `" tabindex="` + index + `" class="navbar-link" name="` + module_name + `">
	<svg aria-hidden="true" focusable="false" class="fa-primary"><use xlink:href="#` + icon_id + `" /></svg>
	<span class="link-text">` + title + `</span></a></li>`
}
