//go:build wasm

package list

import (
	"github.com/tinywasm/dom"
)

func (l *List) OnMount() {
	dom.Log("List component mounted")
}
