package v

import "io"

type Node interface {
	WriteHTML(w io.Writer) (int, error)
}

type NodeFunc func(w io.Writer) (int, error)

func (f NodeFunc) WriteHTML(w io.Writer) (int, error) {
	return f(w)
}

type Nodes []Node

func (n Nodes) WriteHTML(w io.Writer) (int, error) {
	total := 0
	for _, node := range n {
		if node == nil {
			continue
		}

		if n, err := node.WriteHTML(w); err != nil {
			return total + n, err
		} else {
			total += n
		}
	}

	return total, nil
}

func Func(f func() (Node, error)) NodeFunc {
	return func(w io.Writer) (int, error) {
		node, err := f()
		if err != nil {
			return 0, err
		}

		return node.WriteHTML(w)
	}
}
