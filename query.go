package platform

func (Theme) QuerySelectorObject(module_name, object_name string) (query string) {
	return "div#" + module_name + " [data-id='" + object_name + "']"
}
