package server

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/broothie/ghx/model"
	"github.com/broothie/ghx/v"
	"github.com/samber/lo"
)

func (s *Server) commentsIndex(w http.ResponseWriter, r *http.Request) {
	post := r.Context().Value("post").(model.Post)
	if err := s.db.WithContext(r.Context()).Find(&post.Comments, "post_id = ?", post.ID).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sort.Slice(post.Comments, func(i, j int) bool { return post.Comments[i].Votes > post.Comments[j].Votes })

	s.Render(w, r, http.StatusOK,
		v.Div(v.Attr{"id": fmt.Sprintf("post_%s_comments", post.ID)},
			lo.Map(post.Comments, func(comment model.Comment, _ int) v.Node { return vComment(comment) })...,
		),
	)
}
