//go:build !wasm

package module

import _ "embed"

//go:embed slider-panel.css
var sliderPanelCss string

//go:embed module.css
var moduleCss string

//go:embed scroll-snap.css
var scrollSnapCss string

func (m *Module) RenderCSS() string {
    return sliderPanelCss + moduleCss + scrollSnapCss
}

func (m *Module) IconSvg() map[string]string {
    return nil
}
