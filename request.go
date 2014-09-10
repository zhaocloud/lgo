package lgo

import "net/http"

type Request struct {
	*http.Request
	PathParams map[string]string
	Chain map[string]interface{}
}

func (req *Request) PathParam(name string) string {
	return req.PathParams[name]
}
