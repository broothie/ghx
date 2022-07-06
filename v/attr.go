package v

import (
	"fmt"
	"html"
	"io"
	"strings"

	"github.com/samber/lo"
)

type Attr map[string]any

func (a Attr) writeHTML(w io.Writer) (int, error) {
	filtered := lo.PickBy(a, func(_ string, value any) bool { return filterFalsey(value) })

	strs := lo.Map(lo.Entries(filtered), func(entry lo.Entry[string, any], _ int) string {
		return fmt.Sprintf("%s=%q", html.EscapeString(entry.Key), html.EscapeString(fmt.Sprint(entry.Value)))
	})

	return w.Write([]byte(strings.Join(strs, " ")))
}

func Attrs(attr ...Attr) Attr {
	return lo.Assign(lo.Map(attr, func(attr Attr, i int) map[string]any { return attr })...)
}

func (a Attr) Merge(attr ...Attr) Attr {
	return Attrs(append([]Attr{a}, attr...)...)
}

func filterFalsey(value any) bool {
	if value == nil {
		return false
	}

	b, ok := value.(bool)
	if ok {
		return b
	}

	return true
}
