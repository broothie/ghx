package server

import (
	"github.com/broothie/ghx/hx"
	"github.com/broothie/ghx/v"
)

func vLayout(node v.Node) v.Node {
	return v.HTML(nil,
		v.Head(nil,
			v.Stylesheet("https://cdn.jsdelivr.net/npm/bootstrap@5.2.0-beta1/dist/css/bootstrap.min.css"),
			v.Title(nil, v.Text("ghx")),
		),
		v.Body(nil,
			v.Div(v.Attr{"class": "sticky-top container mx-auto d-flex flex-row p-3"},
				v.A(v.Attr{"href": "/", "class": "me-1"}, v.Text("home")),
				v.A(v.Attr{"href": "/posts/new"}, v.Text("new post")),
			),
			node,
			v.JS("https://cdn.jsdelivr.net/npm/bootstrap@5.2.0-beta1/dist/js/bootstrap.bundle.min.js"),
			hx.Script,
		),
	)
}
