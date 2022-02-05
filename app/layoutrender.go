package app

import (
	"html/template"
	"net/http"
)

type LayoutRender struct {
	TemplateDir string
	Layout      string
	Page        string
	Data        interface{}
	Delims      Delims
	Ext         string
}

type Delims struct {
	Left  string
	Right string
}

func NewLayoutRender(templateDir, layout string) LayoutRender {
	return LayoutRender{
		TemplateDir: templateDir,
		Layout:      layout,
		Delims: Delims{
			Left:  "{{",
			Right: "}}",
		},
		Ext: ".html",
	}
}

func (r LayoutRender) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)

	var files []string = []string{
		r.TemplateDir + "/layouts/" + r.Layout + r.Ext,
		r.TemplateDir + "/pages/" + r.Page + r.Ext,
	}

	tpl := template.Must(
		template.New(r.Page).
			Delims(r.Delims.Left, r.Delims.Right).
			Funcs(template.FuncMap{}).
			ParseFiles(files...),
	)

	return tpl.ExecuteTemplate(w, r.Layout+r.Ext, r.Data)
}

func (r LayoutRender) WriteContentType(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
}
