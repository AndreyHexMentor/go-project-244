package code

import (
	"code/internal/diff"
	"code/internal/formatters"
	"code/internal/parser"
	"fmt"
)

func GenDiff(path1, path2, formatName string) (string, error) {
	data1, err := parser.Parse(path1)
	if err != nil {
		return "", fmt.Errorf("parse %s: %w", path1, err)
	}
	data2, err := parser.Parse(path2)
	if err != nil {
		return "", fmt.Errorf("parse %s: %w", path2, err)
	}

	tree := diff.BuildDiff(data1, data2)
	out, err := formatters.Format(tree, formatName)
	if err != nil {
		return "", err
	}

	return out, nil
}
