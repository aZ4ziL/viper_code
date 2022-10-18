package renderer

import (
	"text/template"

	"github.com/gin-contrib/multitemplate"
)

func CreateMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	// Home template
	r.AddFromFilesFuncs(
		"index",
		template.FuncMap{},
		"views/home/base.html",
		"views/home/index.html",
	)

	return r
}
