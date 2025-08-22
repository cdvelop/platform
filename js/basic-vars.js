// Obtén el valor de la variable CSS utilizando getComputedStyle
const root_styles = getComputedStyle(document.documentElement);

// tamaño por defecto campos formulario
const field_height = parseFloat(root_styles.getPropertyValue('--field-height'));