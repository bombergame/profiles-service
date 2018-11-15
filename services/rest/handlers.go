package rest

import (
	"github.com/bombergame/profiles-service/clients/auth-service/grpc"
	"net/http"
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

	if pd.Password != "" {
		go func() {
			err := srv.config.AuthGrpc.DeleteAllSessions(
				&authgrpc.ProfileID{
					Value: profileID,
				},
			)
			if err != nil {
				srv.config.Logger.Error(err)
			}
		}()
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
