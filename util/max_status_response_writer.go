package util

import (
	"bufio"
	"math"
	"net"
	"net/http"
)

type maxStatusResponseWriter struct {
	code  int
	wrote bool
	http.ResponseWriter
}

// MaxStatusResponseWriter は最大のステータスコードを書き込みます。
func MaxStatusResponseWriter(res http.ResponseWriter) http.ResponseWriter {
	return &maxStatusResponseWriter{
		code:           http.StatusOK,
		ResponseWriter: res,
	}
}

func (w *maxStatusResponseWriter) WriteHeader(code int) {
	w.code = int(math.Max(float64(w.code), float64(code)))
}

func (w *maxStatusResponseWriter) Write(b []byte) (int, error) {
	if !w.wrote {
		w.ResponseWriter.WriteHeader(w.code)
		w.wrote = true
	}
	return w.ResponseWriter.Write(b)
}

func (w *maxStatusResponseWriter) Flush() {
	w.ResponseWriter.(http.Flusher).Flush()
}

func (w *maxStatusResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return w.ResponseWriter.(http.Hijacker).Hijack()
}
