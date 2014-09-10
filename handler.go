package lgo

import (
	"net/http"
	"log"
)


type HandlerFunc func(ResponseWriter, *Request)

type RequestHandler struct {
	handlerFunc http.HandlerFunc
	internalRequestMapping *requestMapping
	Middlewares []Middleware
	isInit bool
}

func NewRequestHandler() *RequestHandler {
	h := &RequestHandler{}
	h.internalRequestMapping = &requestMapping{urlMap:map[string][]*controllerInfo{}}
	return h
}

func (h *RequestHandler) Init() {
	h.isInit = true
	middlewares := []Middleware{}
	middlewares = append(middlewares, &errorMiddleware{})
	middlewares = append(middlewares, h.Middlewares...)

	h.handlerFunc = h.adapter(WrapMiddlewares(middlewares, h.service()))
}

func (rh *RequestHandler) adapter(handler HandlerFunc) http.HandlerFunc {
	return func(origWriter http.ResponseWriter, origRequest *http.Request) {
		request := Request{
			origRequest,
			nil,
			map[string]interface{}{},
		}
		writer := responseWriter{
			origWriter,
			false,
		}
		handler(&writer, &request)
	}
}

func (rh *RequestHandler) service() HandlerFunc {
	return func(w ResponseWriter, r *Request) {
		info := rh.internalRequestMapping.GetHandler(r)
		if info == nil {
			Abort(w, "404 Not found", http.StatusNotFound)
			return
		}
		ctx := &Context{Request:r,Response: w}
		info.handler.inject(ctx)
		switch r.Method {
		case "GET":
			info.handler.Get()
		default:
			info.handler.Post()
		}

	}
}

func (rh *RequestHandler) Route(method, path string, controller RestController) {
	rh.internalRequestMapping.Add(method, path, controller);
}

func (rh *RequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rh.handlerFunc(w, r)
}

func Run(port string, h *RequestHandler) {
	if !h.isInit {
		h.Init()
	}
	log.Println("start..")
	log.Fatal(http.ListenAndServe(port, h))
}
