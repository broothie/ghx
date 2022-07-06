package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/broothie/ghx/model"
	"github.com/broothie/ghx/v"
	"github.com/go-chi/chi/v5"
)

func (s *Server) requireComment(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		comment := model.Comment{Base: model.ID(chi.URLParam(r, "commentID"))}
		if tx := s.db.WithContext(r.Context()).Find(&comment); tx.Error != nil {
			http.Error(w, tx.Error.Error(), http.StatusInternalServerError)
			return
		} else if tx.RowsAffected == 0 {
			http.Error(w, fmt.Sprintf("no comment with id %q", comment.ID), http.StatusInternalServerError)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "comment", comment)))
	})
}

func vComment(comment model.Comment) v.Node {
	return v.Div(v.Attr{"class": "border p-3"},
		v.P(nil, v.Text(comment.Text)),
		v.Div(v.Attr{"class": "d-flex flex-row text-muted"},
			v.Div(v.Attr{"class": "me-1"}, vCommentVotes(comment)),
			v.Text(fmt.Sprintf("%s ago", time.Since(comment.CreatedAt))),
		),
	)
}

func vCommentVotes(comment model.Comment) v.Node {
	return v.Div(v.Attr{
		"id":    fmt.Sprintf("post_%s_comment_%s_votes", comment.PostID, comment.ID),
		"class": "d-flex flex-row",
	},
		v.P(v.Attr{
			"class":     "me-1",
			"role":      "button",
			"hx-patch":  fmt.Sprintf("/posts/%s/comments/%s/upvote", comment.PostID, comment.ID),
			"hx-target": fmt.Sprintf("#post_%s_comment_%s_votes", comment.PostID, comment.ID),
		},
			v.Text("^"),
		),
		v.Text(fmt.Sprintf("%d votes", comment.Votes)),
	)
}

func vCommentForm(post model.Post, text string, message string) v.Node {
	return v.Form(v.Attr{
		"id":          fmt.Sprintf("post_%s_comments_form", post.ID),
		"class":       "border p-3",
		"hx-post":     fmt.Sprintf("/posts/%s/comments", post.ID),
		"hx-target":   fmt.Sprintf("#post_%s_comments", post.ID),
		"hx-swap":     "beforeend",
		"hx-swap-oob": true,
	},
		v.Div(v.Attr{"class": "mb-3"},
			v.Label(v.Attr{
				"for":   fmt.Sprintf("post_%s_comments_form_text", post.ID),
				"class": "form-label",
			}),
			v.TextArea(v.Attr{
				"id":          fmt.Sprintf("post_%s_comments_form_text", post.ID),
				"name":        "text",
				"placeholder": "comment a comment",
				"class":       v.Classes{"form-control": true, "is-invalid": message != ""},
				"required":    true,
			},
				v.Text(text),
			),
			v.If(message != "", v.Func(func() (v.Node, error) {
				return v.Div(v.Attr{"class": "invalid-feedback"}, v.Text(message)), nil
			})),
		),
		v.Div(v.Attr{"class": "d-flex flex-row justify-content-end"},
			v.Button(v.Attr{"type": "submit", "class": "btn btn-primary"}, v.Text("submit")),
		),
	)
}
