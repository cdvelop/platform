//go:build !wasm

package login

import _ "embed"

//go:embed login.css
var css string

func (l *Login) RenderCSS() string {
    return css
}

func (l *Login) IconSvg() map[string]string {
    return nil
}
