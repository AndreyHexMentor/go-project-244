package formatters

import (
	"code/internal/diff"
	"fmt"
)

func Format(tree []diff.Node, formatName string) (string, error) {
	switch formatName {
	case "stylish", "":
		return formatStylish(tree)
	case "plain":
		return formatPlain(tree)
	case "json":
		return formatJSON(tree)
	default:
		return "", fmt.Errorf("unknown format: %s", formatName)
	}
}
