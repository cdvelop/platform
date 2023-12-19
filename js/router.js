let HASH_OLD;
let HASH_NOW;

// console.log("ROUTER")

document.querySelector(".menu-container").addEventListener("click", function (e) {
	let hash_target;
	// console.log("REQUEST_ROUTER: ", e.target, " TAG: ", e.target.tagName);

	switch (e.target.tagName) {
		case "svg":
			// console.log("SVG ",hash_target.parentNode)
			hash_target = e.target.parentNode;
			break;

		case "path":
			hash_target = e.target.parentNode.parentNode;
			break;

		case "SPAN":
			hash_target = e.target.parentNode;
			// console.log("SPAN ",hash_target)
			break;

		case "use":
			hash_target = e.target.parentNode.parentNode;

			break;
		default:
			hash_target = e.target
			break;
	}

	// console.log("TARGET FINAL: ",hash_target);

	// moduleClickedUI(hash_target.name);

	OnOffHash(hash_target);
});

// prendo apago link botones rutas
function OnOffHash(hash_now) {
	if (HASH_OLD === undefined) {
		HASH_OLD = hash_now;

		HASH_OLD.classList.toggle("hash-selected");

	} else {

		// remover addEventListener module anterior
		if (MODULES[HASH_OLD.name] != undefined && MODULES[HASH_OLD.name].hasOwnProperty('ListenerModuleOFF')) {
			// apago listener module anterior
			MODULES[HASH_OLD.name].ListenerModuleOFF();
		}
		HASH_OLD.classList.toggle("hash-selected");
		hash_now.classList.toggle("hash-selected");
		// changeClass(HASH_OLD, "svg", "btn-url");
		// changeClass(hash_now, "svg", "btn-selected");
		HASH_OLD = hash_now;
	}

	// console.log("RUTA ACTUAL: ", hash_now.name)

	if (MODULES[hash_now.name] != undefined && MODULES[hash_now.name].hasOwnProperty('ListenerModuleON')) {
		// iniciar addEventListener module actual
		MODULES[hash_now.name].ListenerModuleON();
	}

};