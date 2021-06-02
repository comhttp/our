package utl

import (
	"fmt"
	"net/http"
	"text/template"
)

type interceptResponseWriter struct {
	http.ResponseWriter
	errH func(http.ResponseWriter, int)
}

func (w *interceptResponseWriter) WriteHeader(status int) {
	if status >= http.StatusBadRequest {
		w.errH(w.ResponseWriter, status)
		w.errH = nil
	} else {
		w.ResponseWriter.WriteHeader(status)
	}
}

type ErrorHandler func(http.ResponseWriter, int)

func (w *interceptResponseWriter) Write(p []byte) (n int, err error) {
	if w.errH == nil {
		return len(p), nil
	}
	return w.ResponseWriter.Write(p)
}

func DefaultErrorHandler(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "text/html")
	template.Must(template.ParseFiles("tpl/sys/error.gohtml", "tpl/sys/base.gohtml")).ExecuteTemplate(w, "base", map[string]interface{}{"status": status})
}

func InterceptHandler(next http.Handler, errH ErrorHandler) http.Handler {
	if errH == nil {
		errH = DefaultErrorHandler
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(&interceptResponseWriter{w, errH}, r)
	})
}

func ErrorLog(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}
