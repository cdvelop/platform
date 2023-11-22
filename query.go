package platform

func (t Theme) QuerySelectorMenuModule(module_name string) string {
	return t.MenuClassName() + " a[name='" + module_name + "']"
}

func (Theme) QuerySelectorModule(module_name string) string {
	return "div#" + module_name
}

func (Theme) QuerySelectorObject(module_name, object_name string) string {
	return "div#" + module_name + " [data-id='" + object_name + "']"
}
