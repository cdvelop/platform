//go:build !wasm

package list

import _ "embed"

//go:embed list.css
var listCss string

//go:embed search.css
var searchCss string

func (l *List) RenderCSS() string {
    return listCss + searchCss
}

func (l *List) IconSvg() map[string]string {
    return nil
}
