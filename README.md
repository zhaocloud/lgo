Lgo is a lightweight RESTful API framework in Golang. Reference to [Beego](http://github.com/astaxie/beego) & [go-json-rest](http://github.com/ant0ine/go-json-rest)

### Basics
Go code:
``` go
package main

import "github.com/litgh/lgo"

type MainController struct {
  lgo.Controller
}

func (this *MainController) Get() {
  this.Json(map[string]string{
    "Body":"Hello World!"
  })
}

func (this *MainController) Post() {
  id := this.PathParam(":id")
  ...
}

func main() {
  h := lgo.NewRequestHandler()
  h.Route("GET", "/", &MainController{})
  h.Route("POST", "/:id", &MainController{})
}
```
