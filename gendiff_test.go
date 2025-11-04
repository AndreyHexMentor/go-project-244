package gendiff

import (
	"code/code"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func setupFixtureFiles(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()

	// JSON файлы
	file1 := `{
  "host": "hexlet.io",
  "timeout": 50,
  "proxy": "123.234.53.22",
  "follow": false
}`
	file2 := `{
  "timeout": 20,
  "verbose": true,
  "host": "hexlet.io"
}`
	require.NoError(t, os.WriteFile(filepath.Join(dir, "file1.json"), []byte(file1), 0600))
	require.NoError(t, os.WriteFile(filepath.Join(dir, "file2.json"), []byte(file2), 0600))

	// YAML файлы
	yaml1 := `host: hexlet.io
timeout: 50
proxy: 123.234.53.22
follow: false`
	yaml2 := `timeout: 20
verbose: true
host: hexlet.io`
	require.NoError(t, os.WriteFile(filepath.Join(dir, "file1.yml"), []byte(yaml1), 0600))
	require.NoError(t, os.WriteFile(filepath.Join(dir, "file2.yml"), []byte(yaml2), 0600))

	return dir
}

func TestGenerate_TableDriven(t *testing.T) {
	fixtures := setupFixtureFiles(t)

	tests := []struct {
		name       string
		file1      string
		file2      string
		format     string
		expectPart string
		expectErr  bool
	}{
		{
			name:       "JSON files, stylish format",
			file1:      "file1.json",
			file2:      "file2.json",
			format:     "stylish",
			expectPart: "+ verbose: true",
		},
		{
			name:       "YAML files, stylish format",
			file1:      "file1.yml",
			file2:      "file2.yml",
			format:     "stylish",
			expectPart: "+ verbose: true",
		},
		{
			name:      "Unknown format",
			file1:     "file1.json",
			file2:     "file2.json",
			format:    "unknown",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path1 := filepath.Join(fixtures, tt.file1)
			path2 := filepath.Join(fixtures, tt.file2)

			out, err := code.GenDiff(path1, path2, tt.format)

			if tt.expectErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Contains(t, out, tt.expectPart)
		})
	}
}
