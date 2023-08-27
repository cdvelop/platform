const user_message_update = document.getElementById('user-desktop-messages');

function ShowMessageToUser(message, ...options) {
	// console.log("MENSAJE: ", message, " OPCIONES: ", options)

	const keywords = ["del", "perm", "stop"];

	if (message != "") {
		let seconds = "" + HowManyWords(message);
		let tipo = "ok"; // Valor predeterminado

		for (let i = 0; i < options.length; i++) {

			if (options[i].includes("err")) {
				tipo = "err";
				seconds = "240";
				break; 
			} else if (keywords.some(keyword => options[i].includes(keyword)))  {
				tipo = "permanent";
				break;
			}
		}

		document.documentElement.style.setProperty('--time-read-waiting', seconds + 's');

		const message_out = '<H4 class="' + tipo + '">' + message + '</H4>';

		user_message_update.innerHTML = message_out;

		// Resto del código...
		if (tipo == "err") {
			console.log(message);
		}
	}
};

function HowManyWords(text) {
	let words = text.split(" ");
	// if (words.length > 10) {
	// 	return words.length - 5
	// }
	return words.length;
}

if (screen.width <= 600) {
	user_mobile_messages.addEventListener('click', closeMobileMessage)
}

function closeMobileMessage() {
	// console.log("CLICK EN MENSAJE",e.target);
	document.documentElement.style.setProperty('--time-read-waiting', '0s');
}

