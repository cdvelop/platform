function updateInputClass(input, classesToAdd, classesToRemove) {
    let fd = input.closest('fieldset');
    if (fd != null) {
        if (classesToAdd && classesToAdd.length > 0) {
            fd.classList.add(...classesToAdd);
        }
        if (classesToRemove && classesToRemove.length > 0) {
            fd.classList.remove(...classesToRemove);
        }
    }
}

function InputRight(input) {
    // console.log("INPUT Right RECIBIDO: ", input);
    updateInputClass(input, ["foka"], ["ferr"]);
}

function InputWrong(input) {
    // console.log("INPUT Wrong RECIBIDO: ", input);
    updateInputClass(input, ["ferr"], ["foka"]);
}

function InputNormal(input) {
    // console.log("INPUT normal RECIBIDO: ", input);
    updateInputClass(input, [], ["ferr", "foka"]);
}