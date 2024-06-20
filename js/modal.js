// Obtener la ventana del mensaje de modo incompatible
let modal_windows = document.getElementById('modal-window');

function modalHandler(e) {
  // console.log("modalHandler:", e.target)
  switch (e.pointerType) {
    case "touch":
      break;
    case "mouse":
      break;
  }

  ShowModal(false, e.target, closeCallbackFun)
}

let closeCallbackFun;
// close_callback_fun
function ShowModal(opt, content, close_callback_fun) {

  closeCallbackFun = close_callback_fun

  if (opt == true) {

    if (modal_windows.firstChild) {
      // Si hay un hijo, reemplázalo con el nuevo contenido
      modal_windows.replaceChild(content, modal_windows.firstChild);
    } else {
      // Si no hay hijos, simplemente agrega el nuevo contenido al contenedor de medios
      modal_windows.appendChild(content); // Agrega la imagen seleccionada al modal
    }

    modal_windows.style.display = 'flex'; // Mostrar en dispositivos móviles en modo horizontal
    // modal_windows.style.opacity = 1;
  } else {
    modal_windows.style.display = 'none'; // Ocultar en dispositivos móviles en modo vertical y en dispositivos de escritorio

    // console.log("closeCallbackFun", closeCallbackFun)
    if (closeCallbackFun != undefined) {
      closeCallbackFun(content)
    }
  }
}




// Función que muestra u oculta el mensaje según el tipo de dispositivo
function showIncompatibleMessage() {
  // function isMobileDevice() {
  //   return /Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent);
  // }

  // Mostrar u ocultar el mensaje según el tipo de dispositivo
  if (isMobileDevice() && window.innerWidth >= 600 && window.innerWidth <= 1024 && window.orientation === 90) {
    console.log("ES MOVIL")
    ShowModal(true, '<H4>MODO DE VISUALIZACIÓN INCOMPATIBLE</H4>')
  } else {
    ShowModal(false, "")
  }
}

// Llamar a la función cuando se carga la página y cuando se cambia el tamaño de la ventana
window.addEventListener('load', showIncompatibleMessage);
window.addEventListener('resize', showIncompatibleMessage);

// detectar si es móvil o no
function isMobileDevice() {
  const isMobile = /Mobi/.test(navigator.userAgent) || /iPhone/.test(navigator.userAgent);
  const isSafari = /Safari/.test(navigator.userAgent);
  const isNarrow = (window.innerWidth / window.innerHeight) < 1.5; // ajusta el valor a lo que consideres una pantalla de proporción alargada
  return isMobile && isSafari && !isNarrow;
}

// X no hagas esto, es muy costoso
// window.addEventListener('resize', () => {
// if (window.innerWidth > 720) i
//  mostrarElementoEnDesktop()
//}
//})

// usa la API de media queries
const desktopMediaQuery = window.matchMedia('(min-width: 720px)')
desktopMediaQuery.addEventlistener('change',(event) => {
if (event.matches) {
  mostrarELementoEnDesktop()
 }
})