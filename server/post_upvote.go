package server

import (
	"net/http"

	"github.com/broothie/ghx/model"
	"gorm.io/gorm"
)

func (s *Server) postUpvote(w http.ResponseWriter, r *http.Request) {
	post := r.Context().Value("post").(model.Post)

	if err := s.db.WithContext(r.Context()).Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(r.Context()).Find(&post).Error; err != nil {
			return err
		}

		if err := tx.WithContext(r.Context()).Model(&post).Update("votes", post.Votes+1).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.Render(w, r, http.StatusOK, vPostVotes(post))
}
