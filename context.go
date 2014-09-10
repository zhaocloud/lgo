package lgo

type Context struct {
	Request *Request
	Response ResponseWriter
}

func (ctx *Context) PathParam(name string) string {
	return ctx.Request.PathParam(name)
}
