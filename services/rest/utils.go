package rest

import (
	"encoding/json"
	"github.com/bombergame/common/errs"
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

func (srv *Service) readRequestBody(v interface{}, r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		return errs.NewInvalidFormatError("invalid request body")
	}
	return nil
}

func (srv *Service) writeOk(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

func (srv *Service) writeError(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}

func (srv *Service) writeErrorWithBody(w http.ResponseWriter, err error) {
	switch srvErr := err.(type) {
	case *errs.NotAuthorizedError:
		w.WriteHeader(http.StatusUnauthorized)

	case *errs.AccessDeniedError:
		w.WriteHeader(http.StatusForbidden)

	case *errs.InvalidFormatError:
		w.WriteHeader(http.StatusUnprocessableEntity)

	case *errs.DuplicateError:
		w.WriteHeader(http.StatusConflict)

	case *errs.NotFoundError:
		w.WriteHeader(http.StatusNotFound)

	case *errs.ServiceError:
		srv.logger.Error(srvErr.InnerError())
		w.WriteHeader(http.StatusInternalServerError)

	default:
		srv.logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		srv.writeText(w, errs.ServiceErrorMessage)
		return
	}

	srv.writeText(w, err.Error())
}

func (srv *Service) writeText(w http.ResponseWriter, txt string) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(txt))
}

func (srv *Service) writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
