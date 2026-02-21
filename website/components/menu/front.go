//go:build wasm

package menu

// import (
// 	"github.com/tinywasm/dom"
// )

func (m *Menu) OnMount() {
	// Listen for hash changes to update active link
	// dom.OnHashChange(func(hash string) {
	// 	m.updateActive(hash)
	// })

	// Set initial active state
	// m.updateActive(dom.GetHash())
}

func (m *Menu) updateActive(hash string) {
	// if hash == "" {
	// 	hash = "#home"
	// }

	// Remove .hash-selected from all
	// for _, item := range m.Items {
	// 	linkID := "menu-link-" + item.ModuleName
	// 	elem, ok := dom.Get(linkID)
	// 	if ok && elem != nil {
			// elem.ClassList().Remove("hash-selected")

			// targetHash := "#" + item.ModuleName
			// if targetHash == hash {
			// 	elem.ClassList().Add("hash-selected")
			// }
	// 	}
	// }
}
