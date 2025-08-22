let press_item_timer;
let threshold = 2000; //  (2 segundos)
let elapsed_time;
let finished_time = false;
let current_element = Object();


function targetHandler(target, callbackClickHandler, callbackDeleteHandler) {
    cancelDeleteTab()// resetear anterior

    // console.log("targetHandler:", target)

    let startTime = Date.now();
    current_element.target = target;
    current_element.click = callbackClickHandler;
    current_element.delete = callbackDeleteHandler;


    const delete_tab = target.querySelector('.delete-tab');

    if (!delete_tab) {
        console.log("ERROR div delete_tab no encontrado", delete_tab,target)
        return
    }


    current_element.deleteTab = delete_tab

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
    current_element.target.addEventListener('touchend', mouseUpTargetHandler)

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
        element.target.removeEventListener('touchend', mouseUpTargetHandler)

        element.deleteTab.removeEventListener('click', cancelDeleteTab)
    }
}

function mouseUpTargetHandler() {
    // console.log("elapsed_time", elapsed_time)
    if (!finished_time) {

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




