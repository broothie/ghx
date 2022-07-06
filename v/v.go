package v

import (
	"io"
	"net/http"
)

func Render(w http.ResponseWriter, statusCode int, nodes ...Node) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(statusCode)
	if _, err := WriteHTML(w, nodes...); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func WriteHTML(w io.Writer, nodes ...Node) (int, error) {
	return Nodes(nodes).WriteHTML(w)
}
