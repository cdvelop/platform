function InputRight(input) {
    console.log("INPUT Right RECIBIDO: ", input);

    let fd = input.closest('fieldset');
    if (fd != null) {
        fd.classList.add("foka");
        fd.classList.remove("ferr");
    }
};

function InputWrong(input) {
    console.log("INPUT Wrong RECIBIDO: ", input);

    let fd = input.closest('fieldset');
    if (fd != null) {
        fd.classList.remove("foka");
        fd.classList.add("ferr");
    }
};