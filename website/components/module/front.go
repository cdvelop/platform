//go:build wasm

package module

import (
	"github.com/tinywasm/dom"
)

func (m *Module) OnMount() {
	dom.Log("Module component mounted")
}
