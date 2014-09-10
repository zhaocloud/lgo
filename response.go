package lgo

import (
	"net/http"
	"encoding/json"
)

type ResponseWriter interface {
	Json(v interface{}) error
	Header() http.Header
	Status(int) ResponseWriter
}

type responseWriter struct {
	http.ResponseWriter
	indentJson bool
}

func Abort(w ResponseWriter, error string, code int) {
	w.Status(code)
	err := w.Json(map[string]string{"Error": error})
	if err != nil {
		panic(err)
	}
}

func (w *responseWriter) Json(v interface{}) error {
	var b []byte
	var err error
	if w.indentJson {
		b, err = json.MarshalIndent(v, "", "  ")
	} else {
		b, err = json.Marshal(v)
	}
	if err != nil {
		return err
	}

	_, err = w.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func (w *responseWriter) Status(code int) ResponseWriter {
	if w.Header().Get("Content-Type") == "" {
		w.Header().Set("Content-Type", "application/json")
	}
	w.ResponseWriter.WriteHeader(code)
	return w
}

func (w *responseWriter) Write(b []byte) (int, error) {
	return w.ResponseWriter.Write(b)
}

func (w *responseWriter) Flush() {
	f := w.ResponseWriter.(http.Flusher)
	f.Flush()
}

func (w *responseWriter) CloseNotify() <-chan bool {
	n := w.ResponseWriter.(http.CloseNotifier)
	return n.CloseNotify()
}
