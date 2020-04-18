package versionControl

import "testing"

func TestRepoNameFromGithubPackage(t *testing.T) {
	tests := []struct{
		description string
		packageName string
		wantRepoName string
	}{
		{
			description: "Should return repo name when package name is found in map",
			packageName: "go.opencensus.io",
			wantRepoName: "census-instrumentation/opencensus-go",
		},
		{
			description: "Should return given packagename param when not found in map",
			packageName: "https://github.com/BurntSushi/toml",
			wantRepoName: "BurntSushi/toml",
		},
	}

	for _, test := range tests {
		gotRepoName, err := repoNameFromGithubPackage(test.packageName)
		if err != nil {
			t.Fatalf("unable to get repo name from package: %v", err)
		}
		if gotRepoName != test.wantRepoName {
			t.Errorf("repo name returned did not match expected repo name, want: %s, got: %s", test.wantRepoName, gotRepoName)
		}
	}
}
