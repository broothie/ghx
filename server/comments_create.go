package server

import (
	"net/http"

	"github.com/broothie/ghx/model"
)

func (s *Server) commentsCreate(w http.ResponseWriter, r *http.Request) {
	post := r.Context().Value("post").(model.Post)
	comment := model.Comment{PostID: post.ID}
	if err := s.unmarshalForm(r, &comment); err != nil {
		s.Render(w, r, http.StatusBadRequest, vCommentForm(post, comment.Text, err.Error()))
		return
	}

	if tx := s.db.WithContext(r.Context()).Create(&comment); tx.Error != nil {
		s.Render(w, r, http.StatusOK, vCommentForm(post, comment.Text, tx.Error.Error()))
		return
	}

	s.Render(w, r, http.StatusCreated,
		vComment(comment),
		vCommentForm(post, "", ""),
	)
}
