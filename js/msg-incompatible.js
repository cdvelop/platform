// Obtener la ventana del mensaje de modo incompatible
var incompatibleMessage = document.getElementById('incompatible-message');

// Función que muestra u oculta el mensaje según el tipo de dispositivo
function showIncompatibleMessage() {
  // function isMobileDevice() {
    //   return /Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent);
    // }
    
    // Mostrar u ocultar el mensaje según el tipo de dispositivo
    if (isMobileDevice() && window.innerWidth >= 600 && window.innerWidth <= 1024 && window.orientation === 90) {
      console.log("ES MOVIL")
      incompatibleMessage.style.display = 'flex'; // Mostrar en dispositivos móviles en modo horizontal
    } else {
      incompatibleMessage.style.display = 'none'; // Ocultar en dispositivos móviles en modo vertical y en dispositivos de escritorio
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