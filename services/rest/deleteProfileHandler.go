package rest

import (
	"net/http"
)

func (srv *Service) deleteProfile(w http.ResponseWriter, r *http.Request) {
	authID, err := srv.readHeaderInt64(r, "X-Profile-ID")
	if err != nil {
		srv.writeErrorWithBody(w, err)
		return
	}

	profileID, err := srv.readPathInt64(r, "profile_id")
	if err != nil {
		srv.writeErrorWithBody(w, err)
		return
	}

	if authID != profileID {
		srv.writeError(w, http.StatusForbidden)
		return
	}

	if err := srv.pfRepo.Delete(profileID); err != nil {
		srv.writeErrorWithBody(w, err)
		return
	}

	srv.writeOk(w)
}
