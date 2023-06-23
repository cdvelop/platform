package platform

func (Theme) ModuleJsTemplate(module_name, functions, listener_add, listener_rem string) string {
	return `MODULES['` + module_name + `'] = (function () {
	let crud = new Object();
	const module = document.getElementById('` + module_name + `');
	` + functions + `
	crud.ListenerModuleON = function () {
	 ` + listener_add + `
	};
	crud.ListenerModuleOFF = function () {
	 ` + listener_rem + `
	};
	return crud;
})();`
}
