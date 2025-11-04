package formatters

import (
	"code/internal/diff"
	"encoding/json"
)

func formatJSON(tree []diff.Node) (string, error) {
	body, err := json.MarshalIndent(tree, "", " ")
	if err != nil {
		return "", err
	}
	return string(body), nil
}
