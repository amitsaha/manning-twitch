package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestGenerateScaffold(t *testing.T) {

	testCases := []struct {
		conf           projectConfig
		expectedOutput string
	}{
		{
			conf: projectConfig{
				Name:         "MyProject",
				localPath:    "/path/to/dir",
				RepoURL:      "github.com/username/myproject",
				StaticAssets: false,
			},
			expectedOutput: "Generating scaffold for project MyProject in /path/to/dir",
		},
	}
	byteBuf := new(bytes.Buffer)
	for _, tc := range testCases {
		generateScaffold(byteBuf, tc.conf)
		if len(tc.expectedOutput) != 0 {
			gotOutput := byteBuf.String()
			if strings.Index(gotOutput, tc.expectedOutput) == -1 {
				t.Errorf("Expected output: %s, Got: %s", tc.expectedOutput, gotOutput)
			}
		}
		byteBuf.Reset()
	}
}
