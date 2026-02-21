//go:build wasm

package button

import (
	"github.com/tinywasm/dom"
)

func (b *Button) OnMount() {
	dom.Log("Button component mounted")
}
