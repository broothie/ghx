package v

import (
	"html"
	"io"
)

func Text(text string) NodeFunc {
	return func(w io.Writer) (int, error) {
		return w.Write([]byte(html.EscapeString(text)))
	}
}
