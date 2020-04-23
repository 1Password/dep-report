package versioncontrol

var GithubRepoURLForPackage = map[string]string{
	"go.opencensus.io":            "https://github.com/census-instrumentation/opencensus-go",
	"google.golang.org/grpc":      "https://github.com/grpc/grpc-go",
	"google.golang.org/genproto":  "https://github.com/googleapis/go-genproto",
	"google.golang.org/appengine": "https://github.com/golang/appengine",
	"google.golang.org/api":       "https://api.github.com/repos/googleapis/google-api-go-client",
	"cloud.google.com/go":         "https://api.github.com/repos/googleapis/google-cloud-go",
	"gopkg.in/check.v1":           "https://github.com/go-check/check",
	"gopkg.in/yaml.v2":            "https://github.com/go-yaml/yaml",
}

var GerritRepoURLForPackage = map[string]string{
	"google.golang.org/api": "https://code-review.googlesource.com/projects/google-api-go-client",
	"cloud.google.com/go":   "https://code-review.googlesource.com/projects/gocloud",
}
var licenseForRepo = map[string]string{
	"golang.org/x/crypto":   "BSD-3-Clause",
	"golang.org/x/sync":     "BSD-3-Clause",
	"golang.org/x/image":    "BSD-3-Clause",
	"golang.org/x/net":      "BSD-3-Clause",
	"golang.org/x/sys":      "BSD-3-Clause",
	"golang.org/x/text":     "BSD-3-Clause",
	"golang.org/x/tools":    "BSD-3-Clause",
	"golang.org/x/oauth2":   "BSD-3-Clause",
	"google.golang.org/api": "BSD-3-Clause",
	"cloud.google.com/go":   "NOASSERTION",
}
