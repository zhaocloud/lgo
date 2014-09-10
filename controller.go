package lgo


type RestController interface{
	Get()
	Post()
	Put()
	Delete()
	inject(obj ...interface{})
}

type controllerInfo struct {
	method string
	pattern string
	handler RestController
}

type Controller struct {
	Ctx *Context
}

func (c *Controller) Get() {}
func (c *Controller) Post() {}
func (c *Controller) Put() {}
func (c *Controller) Delete() {}
func (c *Controller) inject(obj ...interface{}) {
	for _, o := range obj {
		switch o.(type) {
		case *Context:
			c.Ctx = o.(*Context)
		}
	}
}

func (c *Controller) Json(v interface{}) error {
	return c.Ctx.Response.Json(v)
}

func (c *Controller) Abort(msg string, code int) {
	Abort(c.Ctx.Response, msg, code)
}
