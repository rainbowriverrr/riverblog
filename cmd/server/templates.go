package main

import (
	"html/template"
	"io"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo/v5"
)

var functions = template.FuncMap{}

type TemplateRegistry struct {
	templates map[string]*template.Template
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Template not found: "+name)
	}
	return tmpl.ExecuteTemplate(w, "base", data)
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./web/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {

		name := filepath.Base(page)

		//Parse base template file and add to set
		//Registers functions in the template, allowing them to be called from the template
		tmpl, err := template.New(name).Funcs(functions).ParseFiles("./web/base.html")
		if err != nil {
			return nil, err
		}

		//Parse partials and add to set
		// tmpl, err = tmpl.ParseGlob("./web/partials/*.html")
		// if err != nil {
		// 	return nil, err
		// }

		//Parse page and add to set
		tmpl, err = tmpl.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = tmpl

	}
	return cache, nil
}
