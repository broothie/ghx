package server

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/broothie/ghx/hx"
	"github.com/broothie/ghx/model"
	"github.com/broothie/ghx/v"
)

func (s *Server) postShow(w http.ResponseWriter, r *http.Request) {
	post := r.Context().Value("post").(model.Post)

	s.Render(w, r, http.StatusOK,
		v.Div(v.Attr{"class": "container mx-auto p-3"},
			v.Div(v.Attr{"class": "border p-3"},
				v.Div(v.Attr{"class": "d-flex flex-row"},
					v.A(v.Attr{"href": post.Link}, v.Text(post.Title)),
					v.P(v.Attr{"class": "text-muted"}, v.Func(func() (v.Node, error) {
						link, err := url.Parse(post.Link)
						if err != nil {
							return nil, err
						}

						return v.Text(fmt.Sprintf("(%s)", link.Hostname())), nil
					})),
				),
				v.Div(v.Attr{"class": "d-flex flex-row text-muted"},
					v.Div(v.Attr{"class": "me-1"}, vPostVotes(post)),
					v.P(v.Attr{"class": "me-1"}, v.Text(fmt.Sprintf("%v ago |", time.Since(post.CreatedAt)))),
					v.P(nil, v.Text(fmt.Sprintf("%d comments", len(post.Comments)))),
				),
			),
			hx.Frame(fmt.Sprintf("/posts/%s/comments", post.ID)),
			vCommentForm(post, "", ""),
		),
	)
}
