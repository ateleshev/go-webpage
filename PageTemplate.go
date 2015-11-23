package webpage

import (
	"html/template"
	"net/http"
	"path"
)

const (
	DEFAULT_TEMPLATE_NAME = "default"

	// Extension
	TEMPLATE_EXT = ".tmpl"

	// Files
	TEMPLATE_ALL    = "*" + TEMPLATE_EXT
	TEMPLATE_LAYOUT = "layout" + TEMPLATE_EXT

	// Directory
	TEMPLATE_DIR_MAIN    = "template"
	TEMPLATE_DIR_VIEW    = "view"
	TEMPLATE_DIR_PARTIAL = "partial"
	TEMPLATE_DIR_WEB     = "web"
)

type PageTemplate struct {
	Name string
	Path string
}

func NewPageTemplate(name string, templatePath string) *PageTemplate { // {{{
	pageTemplate := &PageTemplate{Name: name, Path: templatePath}

	return pageTemplate
} // }}}

func (this *PageTemplate) init() (*template.Template, error) { // {{{
	var err error
	var tmpl *template.Template

	if tmpl, err = template.ParseFiles(path.Join(this.Path, TEMPLATE_DIR_MAIN, this.Name, TEMPLATE_LAYOUT)); err != nil {
		return nil, err
	}

	if tmpl, err = tmpl.ParseGlob(path.Join(this.Path, TEMPLATE_DIR_MAIN, this.Name, TEMPLATE_DIR_PARTIAL, TEMPLATE_ALL)); err != nil {
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

	if tmpl, err = tmpl.ParseGlob(path.Join(this.Path, TEMPLATE_DIR_MAIN, this.Name, TEMPLATE_DIR_VIEW, page.Name, TEMPLATE_ALL)); err != nil {
		return err
	}

	if err = tmpl.Execute(w, page); err != nil {
		return err
	}

	return nil
} // }}}
