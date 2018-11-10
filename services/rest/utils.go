package rest

import (
	"net/http"
)

type LoggingResponseWriter struct {
	status int
	writer http.ResponseWriter
}

func (w *LoggingResponseWriter) Header() http.Header {
	return w.writer.Header()
}

func (w *LoggingResponseWriter) Write(b []byte) (int, error) {
	return w.writer.Write(b)
}

func (w *LoggingResponseWriter) WriteHeader(status int) {
	w.status = status
	w.writer.WriteHeader(status)
}

func (srv *Service) writeOk(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

func (srv *Service) writeError(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}

func (srv *Service) writeErrorWithJSON(w http.ResponseWriter, err error) {

}
