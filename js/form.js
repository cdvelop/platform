function updateInputClass(input, classesToAdd, classesToRemove,message) {
    let fd = input.closest('fieldset');

    // console.log("FIELDSET SELECCIONADO: ",fd," classesToAdd: ",classesToAdd," classesToRemove: ",classesToRemove)
    if (fd != null) {
        fd.dataset.err = message;
        if (classesToAdd && classesToAdd.length > 0) {
            fd.classList.add(...classesToAdd);
        }
        if (classesToRemove && classesToRemove.length > 0) {
            fd.classList.remove(...classesToRemove);
        }
    }
}

function inputRight(input) {
    // console.log("INPUT Right RECIBIDO: ", input);
    updateInputClass(input, ["foka"], ["ferr"],"");
}

function inputWrong(input,message) {
    // console.log("INPUT Wrong:", message,input);
    updateInputClass(input, ["ferr"], ["foka"],message);
}

function inputNormal(input) {
    // console.log("INPUT normal RECIBIDO: ", input);
    updateInputClass(input, [], ["ferr", "foka"],"");
}


function focusNextInput(input) {
    const cont = input.parentNode.parentNode;
    const index_value = parseInt(cont.getAttribute('tabindex'), 10); // Obtener el valor actual y convertirlo a número

    if (index_value != null && index_value >= 0) {
        const form = input.form;
        const next_number = index_value + 1
        // console.log("next_number: ", next_number);
        var next_cont = form.querySelector('[tabindex="' + next_number + '"]');
        if (next_cont != null) {
            // console.log("next_cont: ", next_cont);
            internalFocus(next_cont)
        }
    }
}

function internalFocus(container) {
    const formElements = container.querySelectorAll('input, select, textarea');
    if (formElements.length > 0) {
        formElements[0].focus();
    }
}


function checkMultipleInputs(input) {
    // const father = input.parentNode.parentNode
	// const checkboxes = father.querySelectorAll('input[type="'+input.type+'"]');
	// let all_selected = true;

	// for (const cb of checkboxes) {
	// 	if (cb.checked) {
	// 		all_selected = false;
	// 		break;
	// 	}
	// }

	// if (all_selected) {
	// 	if (!input.hasAttribute('required')) {
    //         inputNormal(input);
	// 	} else {
	// 		// inputWrong(input);
	// 	}
	// } else {
	// 	// inputRight(input);
	// }

    userTyping(input)
}