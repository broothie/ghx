package server

import (
	"net/http"

	"github.com/broothie/ghx/v"
)

func (s *Server) postsNew(w http.ResponseWriter, r *http.Request) {
	s.Render(w, r, http.StatusOK,
		v.Form(v.Attr{
			"action": "/posts",
			"method": http.MethodPost,
			"class":  "container mx-auto d-flex flex-column p-5",
		},
			v.H3(nil, v.Text("create a post")),
			v.Div(v.Attr{"class": "mb-3"},
				v.Input(v.Attr{
					"type":        "text",
					"name":        "title",
					"placeholder": "title",
					"required":    true,
					"class":       "form-control",
				}),
			),
			v.Div(v.Attr{"class": "mb-3"},
				v.Input(v.Attr{
					"type":        "text",
					"name":        "link",
					"placeholder": "link",
					"required":    true,
					"class":       "form-control",
				}),
			),
			v.Div(v.Attr{"class": "d-flex flex-row justify-content-end"},
				v.A(v.Attr{"href": "/", "class": "btn btn-secondary me-3"}, v.Text("back")),
				v.Button(v.Attr{"type": "submit", "class": "btn btn-primary"}, v.Text("submit")),
			),
		),
	)
}
