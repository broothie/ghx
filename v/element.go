package v

import (
	"fmt"
	"io"
)

type Element struct {
	Name       string
	Attributes Attr
	Nodes      Nodes
}

func (e Element) WriteHTML(w io.Writer) (int, error) {
	total := 0

	if n, err := fmt.Fprintf(w, "<%s", e.Name); err != nil {
		return total + n, err
	} else {
		total += n
	}

	if len(e.Attributes) > 0 {
		if n, err := fmt.Fprint(w, " "); err != nil {
			return total + n, err
		} else {
			total += n
		}

		if n, err := e.Attributes.writeHTML(w); err != nil {
			return total + n, err
		} else {
			total += n
		}
	}

	if n, err := fmt.Fprint(w, ">"); err != nil {
		return total + n, err
	} else {
		total += n
	}

	if n, err := e.Nodes.WriteHTML(w); err != nil {
		return total + n, err
	} else {
		total += n
	}

	if n, err := fmt.Fprintf(w, "</%s>", e.Name); err != nil {
		return total + n, err
	} else {
		total += n
	}

	return total, nil
}
