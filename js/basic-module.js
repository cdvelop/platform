
const __col_select = getComputedStyle(document.documentElement).getPropertyValue('--color-selection');

function clickButtonMenuByModuleName(module_name) {
	let menuButton = document.querySelector('.menu-container a[name="' + module_name + '"]');
	if (menuButton != undefined) {
		setTimeout(function () {
			menuButton.click();
		}, 500);
	};
};


