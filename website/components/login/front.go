//go:build wasm

package login

import (
	"github.com/tinywasm/dom"
)

func (l *Login) OnMount() {
	dom.Log("Login component mounted")
}
