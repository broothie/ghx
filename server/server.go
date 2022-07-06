package server

import (
	"log"
	"net/http"
	"os"

	"github.com/broothie/ghx/hx"
	"github.com/broothie/ghx/model"
	"github.com/gorilla/schema"
	"github.com/pkg/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Server struct {
	schemaDecoder *schema.Decoder
	db            *gorm.DB
	hx.HX
}

func New() (*Server, error) {
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

	return &Server{
		schemaDecoder: schema.NewDecoder(),
		db:            db,
		HX:            hx.HX{Layout: vLayout},
	}, nil
}

func (s *Server) unmarshalForm(r *http.Request, dst any) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	if err := s.schemaDecoder.Decode(dst, r.PostForm); err != nil {
		return err
	}

	return nil
}
