package formatters

import (
	"code/internal/diff"
	"fmt"
	"strings"
)

func valueToPlain(v interface{}) string {
	switch vv := v.(type) {
	case map[string]interface{}:
		return "[complex value]"
	case string:
		return fmt.Sprintf("'%s'", vv)
	case nil:
		return "null"
	default:
		return fmt.Sprintf("%v", vv)
	}
}

func buildPlain(nodes []diff.Node, path string, lines *[]string) {
	for _, n := range nodes {
		property := n.Key
		if path != "" {
			property = path + "." + n.Key
		}
		switch n.Type {
		case diff.Nested:
			buildPlain(n.Children, property, lines)
		case diff.Added:
			*lines = append(*lines, fmt.Sprintf("Property '%s' was added with value: %s", property, valueToPlain(n.Value)))
		case diff.Removed:
			*lines = append(*lines, fmt.Sprintf("Property '%s' was removed", property))
		case diff.Changed:
			*lines = append(*lines, fmt.Sprintf("Property '%s' was updated. From %s to %s", property, valueToPlain(n.OldValue), valueToPlain(n.Value)))
		case diff.Unchanged:
			// skip
		}
	}
}

func formatPlain(tree []diff.Node) (string, error) {
	var lines []string
	buildPlain(tree, "", &lines)
	return strings.Join(lines, "\n"), nil
}
