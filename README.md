# go-webpage

```
go get -u github.com/ateleshev/go-webpage
```

```
import (
  "github.com/ateleshev/go-webpage"
  "github.com/ateleshev/go-webcontext"
)
```

Usage:
```
func (this *HomeWebController) Execute(w http.ResponseWriter, r *http.Request) {
  /**
   * Some actions
   */

  page := webpage.NewRootPage(webcontext.NewContext(), "home")
  page.Data().Set("Title", "Home")
  page.Data().Set("Description", "Home page")

  path, _ := os.Getwd()
  template := webpage.NewPageTemplate(this.TemplateName, path)
  if err := template.Execute(w, page); err != nil {
    log.Printf("Cannot execute Home web controller: %v", err)
  }
}
```
