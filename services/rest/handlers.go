package profilesrest

import (
	"github.com/bombergame/common/consts"
	"github.com/bombergame/common/errs"
	"github.com/bombergame/profiles-service/domains"
	"net/http"
)

func (srv *Service) createProfile(w http.ResponseWriter, r *http.Request) {
	var pd NewProfileData
	if err := srv.ReadRequestBody(&pd, r); err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	if err := pd.Validate(); err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	if err := srv.components.ProfileRepository.Create(pd.Prepare()); err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	srv.WriteOk(w)
}

func (srv *Service) getProfiles(w http.ResponseWriter, r *http.Request) {
	pageIndex, err := srv.readPageIndex(r)
	if err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	pageSize, err := srv.readPageSize(r)
	if err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	list, err := srv.components.ProfileRepository.GetAllPaginated(pageIndex, pageSize)
	if err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	srv.WriteOkWithBody(w, domains.Profiles(list))
}

func (srv *Service) getProfile(w http.ResponseWriter, r *http.Request) {
	id, err := srv.readProfileID(r)
	if err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	p, err := srv.components.ProfileRepository.FindByID(id)
	if err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	srv.WriteOkWithBody(w, p)
}

func (srv *Service) updateProfile(w http.ResponseWriter, r *http.Request) {
	authID, err := srv.ReadAuthProfileID(r)
	if err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	profileID, err := srv.readProfileID(r)
	if err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	if authID != profileID {
		err := errs.NewAccessDeniedError()
		srv.WriteErrorWithBody(w, err)
		return
	}

	var pd ProfileDataUpdate
	if err := srv.ReadRequestBody(&pd, r); err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	if err := pd.Validate(); err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	if err := srv.components.ProfileRepository.Update(profileID, pd.Prepare()); err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	if pd.Password != consts.EmptyString {
		go srv.deleteSessions(profileID)
	}

	srv.WriteOk(w)
}

func (srv *Service) deleteProfile(w http.ResponseWriter, r *http.Request) {
	authID, err := srv.ReadAuthProfileID(r)
	if err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	profileID, err := srv.readProfileID(r)
	if err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	if authID != profileID {
		err := errs.NewAccessDeniedError()
		srv.WriteErrorWithBody(w, err)
		return
	}

	if err := srv.components.ProfileRepository.Delete(profileID); err != nil {
		srv.WriteErrorWithBody(w, err)
		return
	}

	go srv.deleteSessions(profileID)

	srv.WriteOk(w)
}
