//go:build !wasm

package button

import _ "embed"

//go:embed crud-buttons.css
var crudButtonsCss string

//go:embed delete_button.css
var deleteButtonCss string

func (b *Button) RenderCSS() string {
	return crudButtonsCss + deleteButtonCss
}

func (b *Button) IconSvg() map[string]string {
	return nil
}
