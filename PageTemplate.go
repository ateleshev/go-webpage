package webpage

import (
	"html/template"
	"net/http"
	"path"
)

const (
	// Extension
	EXT = ".tmpl"

	// Files
	ALL    = "*" + EXT
	LAYOUT = "layout" + EXT

	// Directory
	DIR_MAIN    = "template"
	DIR_VIEW    = "view"
	DIR_PARTIAL = "partial"
)

type PageTemplate struct {
	Name string
	Path string
}

func NewPageTemplate(name string, path string) *PageTemplate { // {{{
	pageTemplate := &PageTemplate{Name: name, Path: path}

	return pageTemplate
} // }}}

func (this *PageTemplate) init() (*template.Template, error) { // {{{
	var err error
	var tmpl *template.Template

	if tmpl, err = template.ParseFiles(path.Join(this.Path, DIR_MAIN, this.Name, LAYOUT)); err != nil {
		return nil, err
	}

	if tmpl, err = tmpl.ParseGlob(path.Join(this.Path, DIR_MAIN, this.Name, DIR_PARTIAL, ALL)); err != nil {
		return nil, err
	}

	return tmpl, nil
} // }}}

func (this *PageTemplate) Execute(w http.ResponseWriter, page *Page) error { // {{{
	var err error
	var tmpl *template.Template

	if tmpl, err = this.init(); err != nil {
		return err
	}

	if tmpl, err = tmpl.ParseGlob(path.Join(this.Path, DIR_MAIN, this.Name, DIR_VIEW, page.Name, ALL)); err != nil {
		return err
	}

	if err = tmpl.Execute(w, page); err != nil {
		return err
	}

	return nil
} // }}}
