package server

import (
	"net/http"

	"github.com/broothie/ghx/model"
	"gorm.io/gorm"
)

func (s *Server) commentUpvote(w http.ResponseWriter, r *http.Request) {
	comment := r.Context().Value("comment").(model.Comment)

	if err := s.db.WithContext(r.Context()).Transaction(func(tx *gorm.DB) error {
		if tx := tx.WithContext(r.Context()).Find(&comment); tx.Error != nil {
			return tx.Error
		}

		if tx := tx.WithContext(r.Context()).Model(&comment).Update("votes", comment.Votes+1); tx.Error != nil {
			return tx.Error
		}

		return nil
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.Render(w, r, http.StatusOK, vCommentVotes(comment))
}
