# go-webpage

```
go get -u github.com/ArtemTeleshev/go-webpage
```

```
import wp "github.com/ArtemTeleshev/go-webpage"
```

Usage:
```
func (this *HomeWebController) Execute(w http.ResponseWriter, r *http.Request) {
  /**
   * Some actions
   */

  page := wp.NewRootPage(this.PageName)
  page.Data().Set("Title", "Home")
  page.Data().Set("Description", "Home page")

  path, _ := os.Getwd()
  template := wp.NewPageTemplate(this.TemplateName, path)
  if err := template.Execute(w, page); err != nil {
    log.Printf("Cannot execute Home web controller: %v", err)
  }
}
```
