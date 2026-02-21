//go:build wasm

package form

import (
	"github.com/tinywasm/dom"
)

func (f *FormStyles) OnMount() {
	dom.Log("Form Styles loaded")
}
