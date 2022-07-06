package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) Handler() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", http.RedirectHandler("/posts", http.StatusPermanentRedirect).ServeHTTP)

	r.Route("/posts", func(r chi.Router) {
		r.Get("/", s.postsIndex)
		r.Post("/", s.postsCreate)
		r.Get("/new", s.postsNew)

		r.Route("/{postID}", func(r chi.Router) {
			r.Use(s.requirePost)

			r.Get("/", s.postShow)
			r.Patch("/upvote", s.postUpvote)

			r.Route("/comments", func(r chi.Router) {
				r.Get("/", s.commentsIndex)
				r.Post("/", s.commentsCreate)

				r.Route("/{commentID}", func(r chi.Router) {
					r.Use(s.requireComment)

					r.Patch("/upvote", s.commentUpvote)
				})
			})
		})
	})

	return r
}
