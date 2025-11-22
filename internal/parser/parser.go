package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func Parse(path string) (map[string]interface{}, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	ext := filepath.Ext(path)
	switch ext {
	case ".json":
		var m map[string]interface{}
		if err := json.Unmarshal(body, &m); err != nil {
			return nil, fmt.Errorf("json unmarshal: %w", err)
		}
		return m, err
	case ".yml", ".yaml":
		var m map[string]interface{}
		if err = yaml.Unmarshal(body, &m); err != nil {
			return nil, fmt.Errorf("yaml unmarshal: %w", err)
		}
		return m, nil
	default:
		return nil, fmt.Errorf("unsupported file extension: %s", ext)
	}
}
