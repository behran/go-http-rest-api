package apiserver

import "net/http"

type responseWriter struct {
	http.ResponseWriter
	code int
}

// переопределяем метод
func (w *responseWriter) WriteHeader(statusCode int) {
	w.code = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
