package rest

import (
	"encoding/json"
	"errors"
	"github.com/bombergame/common/consts"
	"github.com/bombergame/common/errs"
	"net/http"
	"strconv"
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

func (srv *Service) readQueryInt32(r *http.Request, name string, defaultValue int32) (int32, error) {
	v := r.URL.Query().Get(name)
	if v == consts.EmptyString {
		return defaultValue, nil
	}

	iv64, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		return 0, errs.NewInvalidFormatError(name + " type mismatch")
	}

	return int32(iv64), nil
}

func (srv *Service) readHeaderString(r *http.Request, name string) (string, error) {
	v := r.Header.Get(name)
	if v == consts.EmptyString {
		err := errors.New(name + " header not set")
		return consts.EmptyString, errs.NewServiceError(err)
	}
	return v, nil
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

func (srv *Service) writeOkWithBody(w http.ResponseWriter, v interface{}) {
	srv.writeJSON(w, http.StatusOK, v)
}

func (srv *Service) writeError(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}

func (srv *Service) writeErrorWithBody(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError

	switch srvErr := err.(type) {
	case *errs.NotAuthorizedError:
		status = http.StatusUnauthorized

	case *errs.AccessDeniedError:
		status = http.StatusForbidden

	case *errs.InvalidFormatError:
		status = http.StatusUnprocessableEntity

	case *errs.DuplicateError:
		status = http.StatusConflict

	case *errs.NotFoundError:
		status = http.StatusNotFound

	case *errs.ServiceError:
		srv.config.Logger.Error(srvErr.InnerError())
		status = http.StatusInternalServerError

	default:
		srv.config.Logger.Error(err.Error())
		srv.writeText(w, http.StatusInternalServerError, errs.ServiceErrorMessage)
		return
	}

	srv.writeText(w, status, err.Error())
}

func (srv *Service) writeText(w http.ResponseWriter, status int, txt string) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(status)
	w.Write([]byte(txt))
}

func (srv *Service) writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
