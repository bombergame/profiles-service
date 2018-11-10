package rest

import (
	"github.com/bombergame/profiles-service/repositories"
	"net/http"
)

func (srv *Service) getProfiles(w http.ResponseWriter, r *http.Request) {
	pageIndex, err := srv.readQueryInt32(r, "page_index", 1)
	if err != nil {
		srv.writeErrorWithBody(w, err)
		return
	}

	pageSize, err := srv.readQueryInt32(r, "page_size", repositories.DefaultPageSize)
	if err != nil {
		srv.writeErrorWithBody(w, err)
		return
	}

	pf, err := srv.pfRepo.GetAllPaginated(pageIndex, pageSize)
	if err != nil {
		srv.writeErrorWithBody(w, err)
		return
	}

	var pAll Profiles
	pAll.Prepare(pf)

	srv.writeOkWithBody(w, pAll)
}
