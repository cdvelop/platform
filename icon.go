package platform

func (Theme) SpriteIconTemplate(id, view_box string, paths ...string) string {
	out := `<symbol id="` + id + `" viewBox="` + view_box + `">`
	for _, d := range paths {
		out += `<path fill="currentColor" d="` + d + `" />`
	}
	out += `</symbol>`
	return out
}
