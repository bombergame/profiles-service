package rest

import (
	"net/http"
)

func (srv *Service) getProfile(w http.ResponseWriter, r *http.Request) {
	id, err := srv.readPathInt64(r, "profile_id")
	if err != nil {
		srv.writeErrorWithBody(w, err)
		return
	}

	p, err := srv.pfRepo.FindByID(id)
	if err != nil {
		srv.writeErrorWithBody(w, err)
		return
	}

	var pf Profile
	pf.Prepare(*p)

	srv.writeOkWithBody(w, pf)
}
