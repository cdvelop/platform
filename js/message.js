const user_desktop_messages = document.getElementById('user-desktop-messages');
let user_mobile_messages = document.getElementById('user-mobile-messages');


function ShowMessageType(data) {
	let seconds = "" + HowManyWords(data.Message);
	if (data.Message != "") {

		let tipo;
		switch (data.Type) {
			case "error":
				tipo = "err"
				seconds = "240";
				break;

			case "del":
				tipo = "del";
				break;

			default:
				tipo = "ok"
				break;
		}

		// console.log("tamaño mensaje", data.Message.length, "PALABRAS: ", seconds)
		document.documentElement.style.setProperty('--time-read-waiting', seconds + 's');

		const message_out = '<H4 class="' + tipo + '">' + data.Message + '</H4>';

		if (screen.width <= 600) {
			user_mobile_messages.innerHTML = message_out;
		} else {
			user_desktop_messages.innerHTML = message_out;
		}

		if (tipo == "err") {
			console.log(data);
		}
	};
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