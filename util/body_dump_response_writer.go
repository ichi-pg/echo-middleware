package util

import (
	"bufio"
	"bytes"
	"io"
	"net"
	"net/http"
)

type bodyDumpResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

// BodyDumpResponseWriter は bodyBuf にレスポンスボディを複写します。
func BodyDumpResponseWriter(bodyBuf *bytes.Buffer, res http.ResponseWriter) http.ResponseWriter {
	return &bodyDumpResponseWriter{
		Writer:         io.MultiWriter(res, bodyBuf),
		ResponseWriter: res,
	}
}

func (w *bodyDumpResponseWriter) WriteHeader(code int) {
	w.ResponseWriter.WriteHeader(code)
}

func (w *bodyDumpResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func (w *bodyDumpResponseWriter) Flush() {
	w.ResponseWriter.(http.Flusher).Flush()
}

func (w *bodyDumpResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return w.ResponseWriter.(http.Hijacker).Hijack()
}
