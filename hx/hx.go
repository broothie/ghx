package hx

import (
	"net/http"

	"github.com/broothie/ghx/v"
)

var Script = v.JS("https://unpkg.com/htmx.org@1.7.0")

func RequestIsHX(r *http.Request) bool {
	return r.Header.Get("HX-Request") == "true"
}

func Frame(url string) v.Node {
	return v.Div(v.Attr{
		"hx-get":     url,
		"hx-trigger": "revealed",
		"hx-swap":    "outerHTML",
	})
}

func PageRedirect(w http.ResponseWriter, r *http.Request, url string, code int) {
	if RequestIsHX(r) {
		w.Header().Set("HX-Redirect", url)
		w.WriteHeader(code)
	} else {
		http.Redirect(w, r, url, code)
	}
}

type LayoutFunc func(v.Node) v.Node

type HX struct {
	Layout LayoutFunc
}

func (hx HX) Render(w http.ResponseWriter, r *http.Request, statusCode int, nodes ...v.Node) {
	var node v.Node = v.Nodes(nodes)
	if hx.Layout != nil && !RequestIsHX(r) {
		node = hx.Layout(node)
	}

	v.Render(w, statusCode, node)
}
