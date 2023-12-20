package platform

func (Theme) MenuClassName() string {
	return ".menu-container"
}

func (Theme) MenuItemClass() string {
	return "navbar-item"
}

func (t Theme) MenuButtonTemplate(module_name, index, icon_id, title string) string {
	return `<li class="` + t.MenuItemClass() + `"><a href="#` + module_name + `" tabindex="` + index + `" class="navbar-link" name="` + module_name + `">
	<svg aria-hidden="true" focusable="false" class="fa-primary"><use xlink:href="#` + icon_id + `" /></svg>
	<span class="link-text">` + title + `</span></a></li>`
}

func (Theme) RouterJSFuncName() string {
	return "menuRouter"
}

func (Theme) MenuClassSelected() string {
	return "hash-selected"
}
