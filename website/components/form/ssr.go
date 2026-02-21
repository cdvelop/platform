//go:build !wasm

package form

import _ "embed"

//go:embed fieldset.css
var fieldsetCss string

//go:embed basic_input.css
var basicInputCss string

//go:embed radio_check.css
var radioCheckCss string

//go:embed add_basic.css
var addBasicCss string

func (f *FormStyles) RenderCSS() string {
	return fieldsetCss + basicInputCss + radioCheckCss + addBasicCss
}

func (f *FormStyles) IconSvg() map[string]string {
	return nil
}
