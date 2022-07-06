package app

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/broothie/ghx/hx"
	"github.com/broothie/ghx/model"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/schema"
	"github.com/pkg/errors"
	"github.com/unrolled/render"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type App struct {
	schemaDecoder *schema.Decoder
	db            *gorm.DB
	render        *render.Render
}

func New() (*App, error) {
	db, err := gorm.Open(sqlite.Open("ghx.sqlite.db"), &gorm.Config{Logger: logger.New(log.New(os.Stdout, "", 0), logger.Config{LogLevel: logger.Info})})
	if err != nil {
		return nil, errors.Wrap(err, "failed to open sqlite db")
	}

	if err := db.AutoMigrate(
		new(model.Post),
		new(model.Comment),
	); err != nil {
		return nil, errors.Wrap(err, "failed to auto migrate")
	}

	return &App{
		schemaDecoder: schema.NewDecoder(),
		db:            db,
		render: render.New(render.Options{
			Extensions:                  []string{".gohtml"},
			IsDevelopment:               true,
			RenderPartialsWithoutPrefix: true,
		}),
	}, nil
}

func (a *App) Handler() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/posts", func(r chi.Router) {
		r.Route("/{postID}", func(r chi.Router) {
			r.Use(a.requirePost)

			r.Get("/", a.postsShow)
		})
	})

	return r
}

func (a *App) requirePost(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var post model.Post
		if tx := a.db.WithContext(r.Context()).Find(&post, "id = ?", chi.URLParam(r, "postID")); tx.Error != nil {
			http.Error(w, tx.Error.Error(), http.StatusInternalServerError)
			return
		} else if tx.RowsAffected == 0 {
			http.Error(w, "not found", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "post", post)))
	})
}

func (a *App) postsShow(w http.ResponseWriter, r *http.Request) {
	a.renderHTML(w, r, http.StatusOK, "posts/show", map[string]any{"post": r.Context().Value("post")})
}

func (a *App) renderHTML(w http.ResponseWriter, r *http.Request, status int, name string, binding any, htmlOpt ...render.HTMLOptions) {
	if !hx.RequestIsHX(r) {
		htmlOpt = append(htmlOpt, render.HTMLOptions{Layout: "layout"})
	}

	a.render.HTML(w, status, name, binding, htmlOpt...)
}
