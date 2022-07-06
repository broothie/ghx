package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/broothie/ghx/hx"
	"github.com/broothie/ghx/model"
	"github.com/broothie/ghx/v"
	"github.com/go-chi/chi/v5"
)

func (s *Server) requirePost(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		post := model.Post{Base: model.ID(chi.URLParam(r, "postID"))}
		if tx := s.db.WithContext(r.Context()).Find(&post); tx.Error != nil {
			http.Error(w, tx.Error.Error(), http.StatusInternalServerError)
			return
		} else if tx.RowsAffected == 0 {
			hx.PageRedirect(w, r, "/", http.StatusNotFound)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "post", post)))
	})
}

func vPostVotes(post model.Post) v.Node {
	return v.Div(v.Attr{"id": fmt.Sprintf("post_%s_item_votes", post.ID), "class": "d-flex flex-row"},
		v.P(v.Attr{
			"class":     "me-1",
			"role":      "button",
			"hx-patch":  fmt.Sprintf("/posts/%s/upvote", post.ID),
			"hx-target": fmt.Sprintf("#post_%s_item_votes", post.ID),
		},
			v.Text("^"),
		),
		v.Text(fmt.Sprintf("%d votes", post.Votes)),
	)
}
