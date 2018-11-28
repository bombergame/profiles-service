package profilesrest

import (
	"errors"
	"fmt"
	"github.com/bombergame/auth-service/services/grpc"
	"github.com/bombergame/common/consts"
	"github.com/bombergame/common/errs"
	"github.com/bombergame/profiles-service/repositories"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

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

func (srv *Service) readPageIndex(r *http.Request) (int32, error) {
	return srv.readQueryInt32(r, "page_index", 1)
}

func (srv *Service) readPageSize(r *http.Request) (int32, error) {
	return srv.readQueryInt32(r, "page_size", repositories.DefaultPageSize)
}

func (srv *Service) deleteSessions(profileID int64) {
	err := srv.components.AuthClient.DeleteAllSessions(
		authgrpc.ProfileID{
			Value: profileID,
		},
	)
	if err != nil {
		srv.Logger().Error(fmt.Sprintf("Error on DELETE_ALL_SESSIONS: %s", err))
	}
}
