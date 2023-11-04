let press_item_timer;
let threshold = 3000; //  (3 segundos)
let elapsed_time;
let finished_time = false;
let current_element = Object();


function targetHandler(target, callbackClickHandler, callbackDeleteHandler) {
    cancelDeleteTab()// resetear anterior

    let startTime = Date.now();
    current_element.target = target;
    current_element.click = callbackClickHandler;
    current_element.delete = callbackDeleteHandler;
    current_element.deleteTab = target.parentNode.querySelector('.delete-tab');

    resetMouseUpPress(current_element);// resetear actual

    press_item_timer = setInterval(() => {
        elapsed_time = Date.now() - startTime;
        let width = (elapsed_time / threshold) * target.offsetWidth;
        width = Math.round(width);
        width = width > 100 ? 100 : width; // Asegura que el ancho no sea mayor al 100%

        current_element.deleteTab.style.width = width + '%';

        if (elapsed_time >= threshold) {
            finished_time = true;
            current_element.delete(current_element.target)
            resetTimer();
        }
    }, 100); // Actualiza la barra cada 100 ms


    current_element.target.addEventListener('mouseup', mouseUpTargetHandler)

    current_element.deleteTab.addEventListener('click', cancelDeleteTab)

}


function cancelDeleteTab() {
    // console.log("cancelDeleteTab cancelar")
    updateDeleteTab(current_element)
    resetTimer()
    resetMouseUpPress(current_element)
}

function resetTimer() {
    if (press_item_timer) {
        clearInterval(press_item_timer);
        // console.log("clearInterval")
    }
}


function resetMouseUpPress(element) {
    // console.log("resetMouseUpPress")
    if (element.hasOwnProperty("target")) {
        element.target.removeEventListener('mouseup', mouseUpTargetHandler)
        element.deleteTab.removeEventListener('click', cancelDeleteTab)
    }
}

function mouseUpTargetHandler() {
    if (!finished_time) {
        // console.log("mouseUpTargetHandler", current_element)

        if (elapsed_time <= 300) {
            current_element.click(current_element.target)
        }

    }
    resetTimer()
    finished_time = false;
    updateDeleteTab(current_element)
}


function updateDeleteTab(element) {
    if (element.hasOwnProperty("deleteTab")) {
        // let deleteTab = element.parentNode.querySelector('.delete-tab');
        current_element.deleteTab.style.width = '0';
    }
}




