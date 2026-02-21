//go:build !wasm

package menu

import _ "embed"

//go:embed menu.css
var css string

func (m *Menu) RenderCSS() string {
    return css
}

func (m *Menu) IconSvg() map[string]string {
    // Menu icons are usually provided by the caller or global.
    // But if Menu defines its own icons, put them here.
    // The existing code uses "icon-home" etc which are in index.html.
    // Layout component handles global icons.
    return nil
}
