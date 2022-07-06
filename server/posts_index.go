package server

import (
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"time"

	"github.com/broothie/ghx/model"
	"github.com/broothie/ghx/v"
	"github.com/samber/lo"
)

func (s *Server) postsIndex(w http.ResponseWriter, r *http.Request) {
	var posts []model.Post
	if err := s.db.WithContext(r.Context()).Preload("Comments").Find(&posts).Limit(30).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sort.Slice(posts, func(i, j int) bool { return posts[i].Votes > posts[j].Votes })

	s.Render(w, r, http.StatusOK,
		v.Div(v.Attr{"class": "container mx-auto"},
			v.Div(nil,
				lo.Map(posts, func(post model.Post, i int) v.Node {
					return v.Div(v.Attr{"class": "d-flex flex-column border p-3"},
						v.Div(v.Attr{"class": "d-flex flex-row"},
							v.P(v.Attr{"class": "me-1"}, v.Text(fmt.Sprintf("%d.", i+1))),
							v.A(v.Attr{"href": post.Link, "class": "me-1"}, v.Text(post.Title)),
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
							v.A(v.Attr{"href": fmt.Sprintf("/posts/%s", post.ID)}, v.Text(fmt.Sprintf("%d comments", len(post.Comments)))),
						),
					)
				})...,
			),
		),
	)
}
