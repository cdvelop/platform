// let press_item_timer;
// let threshold = 3000; // (3 segundos)
// let clickThreshold = 200; // (milisegundos)
// let startTime; // Variable global

function startDeleteTimerOLD(target, callbackClickHandler, callbackDeleteHandler) {
    let deleteTab = target.parentNode.querySelector('.delete-tab');
    startTime = Date.now();
    let clickDetected = false;

    press_item_timer = setInterval(() => {
        const elapsed_time = Date.now() - startTime;
        let width = (elapsed_time / threshold) * target.offsetWidth;
        width = Math.round(width);
        width = width > 100 ? 100 : width;

        deleteTab.style.width = width + '%';

        if (elapsed_time >= threshold) {
            clearDeleteTimer(target, callbackClickHandler, callbackDeleteHandler);
        }

        if (!clickDetected && elapsed_time <= clickThreshold) {
            // Si se detecta un clic antes de 1 segundo, llama a callbackClickHandler
            clickDetected = true;
        }
    }, 100);
}

function clearDeleteTimerOLD(target, callbackClickHandler, callbackDeleteHandler) {
    if (press_item_timer) {
        clearInterval(press_item_timer);
    }

    const elapsed_time = Date.now() - startTime;

    if (elapsed_time <= clickThreshold && callbackClickHandler) {
        callbackClickHandler(target);
    } else if (callbackDeleteHandler) {
        callbackDeleteHandler(target);
    }

    if (target != undefined) {
        let deleteTab = target.parentNode.querySelector('.delete-tab');
        deleteTab.style.width = '0';
    }
}
