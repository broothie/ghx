package v

func Tag(name string, attr Attr, nodes ...Node) Element {
	return Element{Name: name, Attributes: attr, Nodes: nodes}
}
