package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestSetupParseFlags(t *testing.T) {

	testCases := []struct {
		args                   []string
		err                    error
		expectedConf           projectConfig
		expectedOutputContains string
	}{
		{
			args: []string{"-n", "MyProject", "-d", "/path/to/dir", "-r", "github.com/username/myproject"},
			err:  nil,
			expectedConf: projectConfig{
				Name:         "MyProject",
				localPath:    "/path/to/dir",
				RepoURL:      "github.com/username/myproject",
				StaticAssets: false},
		},
		{
			args:         []string{"foo"},
			err:          errors.New("No positional parameters expected"),
			expectedConf: projectConfig{},
		},
		{
			args:                   []string{"-h"},
			err:                    errors.New("flag: help requested"),
			expectedConf:           projectConfig{},
			expectedOutputContains: "Usage of scaffold-gen:",
		},
	}

	byteBuf := new(bytes.Buffer)
	for _, tc := range testCases {
		c, err := setupParseFlags(byteBuf, tc.args)
		if tc.err == nil && err != nil {
			t.Errorf("Expected non-nil error, got: %v", err)
		}

		if tc.err != nil {
			if err == nil || err.Error() != tc.err.Error() {
				t.Errorf("Expected error: %v, Got: %v", tc.err, err)
			}
		}

		if c != tc.expectedConf {
			t.Errorf("Expected:%#v Got: %#v", c, tc.expectedConf)
		}

		if len(tc.expectedOutputContains) != 0 {
			// setupParseFlags() -> writes a message -> the byteBuf
			gotOutput := byteBuf.String()
			if strings.Index(gotOutput, tc.expectedOutputContains) == -1 {
				t.Errorf("Expected output: %s, Got: %s", tc.expectedOutputContains, gotOutput)
			}
		}
		byteBuf.Reset()
	}
}
