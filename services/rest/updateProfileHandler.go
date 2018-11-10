package rest

import (
	"net/http"
)

func (srv *Service) updateProfile(w http.ResponseWriter, r *http.Request) {
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

	var pd ProfileDataUpdate
	if err := srv.readRequestBody(&pd, r); err != nil {
		srv.writeErrorWithBody(w, err)
		return
	}

	if err := pd.Validate(); err != nil {
		srv.writeErrorWithBody(w, err)
		return
	}

	if err := srv.pfRepo.Update(profileID, pd.Prepare()); err != nil {
		srv.writeErrorWithBody(w, err)
		return
	}

	srv.writeOk(w)
}
