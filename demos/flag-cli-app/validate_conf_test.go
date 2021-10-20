package main

import (
	"errors"
	"testing"
)

func TestValidate(t *testing.T) {

	testCases := []struct {
		conf projectConfig
		errs []error
	}{
		{
			conf: projectConfig{
				Name:         "MyProject",
				localPath:    "/path/to/dir",
				RepoURL:      "github.com/username/myproject",
				StaticAssets: false,
			},
			errs: []error{},
		},
		{
			conf: projectConfig{
				StaticAssets: false,
			},
			errs: []error{
				errors.New("Project name cannot be empty"),
				errors.New("Project path cannot be empty"),
				errors.New("Project repository URL cannot be empty"),
			},
		},
	}
	for _, tc := range testCases {
		errs := validateConf(tc.conf)
		if len(tc.errs) == 0 && len(errs) != 0 {
			t.Errorf("Expected no errors, got: %v", errs)
		}

		if len(tc.errs) != 0 {
			for i, e := range tc.errs {
				if errs[i] == nil || errs[i].Error() != e.Error() {
					t.Errorf("Expected error: %v, Got: %v", e, errs[i])
				}
			}
		}
	}
}
