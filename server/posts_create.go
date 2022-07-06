package server

import (
	"fmt"
	"net/http"

	"github.com/broothie/ghx/hx"
	"github.com/broothie/ghx/model"
)

func (s *Server) postsCreate(w http.ResponseWriter, r *http.Request) {
	var post model.Post
	if err := s.unmarshalForm(r, &post); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.db.WithContext(r.Context()).Create(&post).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hx.PageRedirect(w, r, fmt.Sprintf("/posts/%s", post.ID), http.StatusSeeOther)
}
