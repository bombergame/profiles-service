package rest

import (
	"encoding/json"
	"errors"
	"github.com/bombergame/common/consts"
	"github.com/bombergame/common/errs"
	"github.com/bombergame/profiles-service/clients/auth-service/grpc"
	"github.com/bombergame/profiles-service/repositories"
	"github.com/gorilla/mux"
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

func (srv *Service) readProfileID(r *http.Request) (int64, error) {
	v, ok := mux.Vars(r)["profile_id"]
	if !ok {
		err := errors.New("profile_id cannot be parsed")
		return 0, errs.NewServiceError(err)
	}

	iv64, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, errs.NewServiceError(err)
	}

	return iv64, nil
}

func (srv *Service) readAuthProfileID(r *http.Request) (int64, error) {
	v, err := srv.readHeaderString(r, "X-Profile-ID")
	if err != nil {
		return 0, err
	}

	iv64, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, errs.NewServiceError(err)
	}

	return iv64, nil
}

func (srv *Service) readPageIndex(r *http.Request) (int32, error) {
	return srv.readQueryInt32(r, "page_index", 1)
}

func (srv *Service) readPageSize(r *http.Request) (int32, error) {
	return srv.readQueryInt32(r, "page_size", repositories.DefaultPageSize)
}

func (srv *Service) readUserAgent(r *http.Request) (string, error) {
	return srv.readHeaderString(r, "User-Agent")
}

func (srv *Service) readAuthToken(r *http.Request) (string, error) {
	bearer, err := srv.readHeaderString(r, "Authorization")
	if err != nil {
		return consts.EmptyString, err
	}

	n := len("Bearer ")
	if len(bearer) <= n {
		return consts.EmptyString, errs.NewInvalidFormatError("wrong authorization token")
	}

	return bearer[n:], nil
}

func (srv *Service) setAuthProfileID(r *http.Request, id int64) {
	r.Header.Set("X-Profile-ID", strconv.FormatInt(id, 10))
}

func (srv *Service) deleteSessions(profileID int64) {
	err := srv.config.AuthGrpc.DeleteAllSessions(
		&authgrpc.ProfileID{
			Value: profileID,
		},
	)
	if err != nil {
		srv.config.Logger.Error(err)
	}
}
