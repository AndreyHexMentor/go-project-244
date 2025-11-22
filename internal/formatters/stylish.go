package formatters

import (
	"code/internal/diff"
	"fmt"
	"sort"
	"strings"
)

const indentSize = 4

// Формирует строку из n пробелов
func indent(n int) string {
	return strings.Repeat(" ", n)
}

// Преобразует значение в строку с учётом вложенности
func stringify(value interface{}, depth int) string {
	switch v := value.(type) {
	case map[string]interface{}:
		return mapToString(v, depth+1)
	case nil:
		return "null"
	case string:
		return v
	default:
		return fmt.Sprintf("%v", v)
	}
}

func mapToString(m map[string]interface{}, depth int) string {
	if len(m) == 0 {
		return "{}"
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var sb strings.Builder
	sb.WriteString("{\n")
	for _, k := range keys {
		space := indent(depth * indentSize)
		sb.WriteString(fmt.Sprintf("%s%s: %s\n", space, k, stringify(m[k], depth)))
	}
	sb.WriteString(fmt.Sprintf("%s}", indent((depth-1)*indentSize)))
	return sb.String()
}

func formatStylish(tree []diff.Node) (string, error) {
	return "{\n" + render(tree, 1) + "}", nil
}

func render(nodes []diff.Node, depth int) string {
	var sb strings.Builder
	indentForMark := indent(depth*indentSize - 2)
	indentPlain := indent(depth * indentSize)

	for _, n := range nodes {
		switch n.Type {
		case diff.Nested:
			sb.WriteString(fmt.Sprintf("%s  %s: {\n", indentPlain[:len(indentPlain)-2], n.Key))
			sb.WriteString(render(n.Children, depth+1))
			sb.WriteString(fmt.Sprintf("%s  }\n", indentPlain[:len(indentPlain)-2]))
		case diff.Unchanged:
			sb.WriteString(fmt.Sprintf("%s  %s: %s\n", indentPlain[:len(indentPlain)-2], n.Key, stringify(n.Value, depth)))
		case diff.Added:
			sb.WriteString(fmt.Sprintf("%s+ %s: %s\n", indentForMark, n.Key, stringify(n.Value, depth)))
		case diff.Removed:
			sb.WriteString(fmt.Sprintf("%s- %s: %s\n", indentForMark, n.Key, stringify(n.Value, depth)))
		case diff.Changed:
			sb.WriteString(fmt.Sprintf("%s- %s: %s\n", indentForMark, n.Key, stringify(n.OldValue, depth)))
			sb.WriteString(fmt.Sprintf("%s+ %s: %s\n", indentForMark, n.Key, stringify(n.Value, depth)))
		}
	}
	return sb.String()
}
