package rest

import (
	"errors"
	"github.com/bombergame/common/errs"
	"github.com/bombergame/profiles-service/repositories"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (srv *Service) createProfile(w http.ResponseWriter, r *http.Request) {
	var pd NewProfileData
	if err := srv.readRequestBody(&pd, r); err != nil {
		srv.writeErrorWithBody(w, err)
		return
	}

	if err := pd.Validate(); err != nil {
		srv.writeErrorWithBody(w, err)
		return
	}

	if err := srv.config.ProfileRepository.Create(pd.Prepare()); err != nil {
		srv.writeErrorWithBody(w, err)
		return
	}

	srv.writeOk(w)
}

func (srv *Service) getProfiles(w http.ResponseWriter, r *http.Request) {
	pageIndex, err := srv.readPageIndex(r)
	if err != nil {
		srv.writeErrorWithBody(w, err)
		return
	}

	pageSize, err := srv.readPageSize(r)
	if err != nil {
		srv.writeErrorWithBody(w, err)
		return
	}

	pf, err := srv.config.ProfileRepository.GetAllPaginated(pageIndex, pageSize)
	if err != nil {
		srv.writeErrorWithBody(w, err)
		return
	}

	var pAll Profiles
	pAll.Prepare(pf)

	srv.writeOkWithBody(w, pAll)
}

func (srv *Service) getProfile(w http.ResponseWriter, r *http.Request) {
	id, err := srv.readProfileID(r)
	if err != nil {
		srv.writeErrorWithBody(w, err)
		return
	}

	p, err := srv.config.ProfileRepository.FindByID(id)
	if err != nil {
		srv.writeErrorWithBody(w, err)
		return
	}

	var pf Profile
	pf.Prepare(*p)

	srv.writeOkWithBody(w, pf)
}

func (srv *Service) updateProfile(w http.ResponseWriter, r *http.Request) {
	authID, err := srv.readAuthProfileID(r)
	if err != nil {
		srv.writeErrorWithBody(w, err)
		return
	}

	profileID, err := srv.readProfileID(r)
	if err != nil {
		srv.writeErrorWithBody(w, err)
		return
	}

	if authID != profileID {
		srv.writeError(w, http.StatusForbidden)
		return
	}

	var pd ProfileDataUpdate
	if err := srv.readRequestBody(&pd, r); err != nil {
		srv.writeErrorWithBody(w, err)
		return
	}

	if err := pd.Validate(); err != nil {
		srv.writeErrorWithBody(w, err)
		return
	}

	if err := srv.config.ProfileRepository.Update(profileID, pd.Prepare()); err != nil {
		srv.writeErrorWithBody(w, err)
		return
	}

	srv.writeOk(w)
}

func (srv *Service) deleteProfile(w http.ResponseWriter, r *http.Request) {
	authID, err := srv.readAuthProfileID(r)
	if err != nil {
		srv.writeErrorWithBody(w, err)
		return
	}

	profileID, err := srv.readProfileID(r)
	if err != nil {
		srv.writeErrorWithBody(w, err)
		return
	}

	if authID != profileID {
		srv.writeError(w, http.StatusForbidden)
		return
	}

	if err := srv.config.ProfileRepository.Delete(profileID); err != nil {
		srv.writeErrorWithBody(w, err)
		return
	}

	srv.writeOk(w)
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
