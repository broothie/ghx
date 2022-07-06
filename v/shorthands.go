package v

import (
	"fmt"
	"io"
	"strings"

	"github.com/samber/lo"
)

func If(condition bool, f NodeFunc) NodeFunc {
	return func(w io.Writer) (int, error) {
		if condition {
			return f.WriteHTML(w)
		}

		return 0, nil
	}
}

func Not(condition bool, f NodeFunc) NodeFunc {
	return If(!condition, f)
}

func IfErr(err error, f NodeFunc) NodeFunc {
	return If(err != nil, f)
}

type Classes map[string]any

func (c Classes) String() string {
	return strings.Join(lo.Keys(lo.PickBy(c, func(_ string, value any) bool { return filterFalsey(value) })), " ")
}

func JS(src string, attr ...Attr) Element {
	return Script(Attr{"src": src, "type": "application/javascript"}.Merge(attr...))
}

func Stylesheet(href string, attr ...Attr) Element {
	return Link(Attr{"href": href, "rel": "stylesheet"}.Merge(attr...))
}

func Datas(attr Attr) Attr {
	return PrefixAttr("data-", attr)
}

func PrefixAttr(prefix string, attr Attr) Attr {
	return lo.MapKeys(attr, func(_ any, key string) string {
		return fmt.Sprintf("%s%s", prefix, key)
	})
}
