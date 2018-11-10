package rest

import (
	"net/http"
)

func (srv *Service) createProfile(w http.ResponseWriter, r *http.Request) {
	var pd NewProfileData
	if err := srv.readRequestBody(&pd, r); err != nil {
		srv.writeErrorWithJSON(w, err)
		return
	}

	if err := pd.Validate(); err != nil {
		srv.writeErrorWithJSON(w, err)
		return
	}

	if err := srv.pfRepo.Create(pd.Prepare()); err != nil {
		srv.writeErrorWithJSON(w, err)
		return
	}

	srv.writeOk(w)
}
