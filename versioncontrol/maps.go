package versioncontrol

var GithubRepoURLForPackage = map[string]string{
	"go.opencensus.io":                "https://github.com/census-instrumentation/opencensus-go",
	"google.golang.org/grpc":          "https://github.com/grpc/grpc-go",
	"google.golang.org/genproto":      "https://github.com/googleapis/go-genproto",
	"google.golang.org/appengine":     "https://github.com/golang/appengine",
	"google.golang.org/api":           "https://api.github.com/repos/googleapis/google-api-go-client",
	"cloud.google.com/go":             "https://api.github.com/repos/googleapis/google-cloud-go",
	"gopkg.in/check.v1":               "https://github.com/go-check/check",
	"gopkg.in/yaml.v2":                "https://github.com/go-yaml/yaml",
	"gopkg.in/square/go-jose.v2":      "https://github.com/square/go-jose",
	"go.etcd.io/bbolt":                "https://github.com/etcd-io/bbolt",
	"go.uber.org/atomic":              "https://github.com/uber-go/atomic",
	"go.uber.org/multierr":            "https://github.com/uber-go/multierr",
	"go.uber.org/zap":                 "https://github.com/uber-go/zap",
	"gopkg.in/resty.v1":               "https://github.com/go-resty/resty",
	"gopkg.in/ini.v1":                 "https://github.com/go-ini/ini",
	"gopkg.in/alecthomas/kingpin.v2":  "https://github.com/alecthomas/kingpin",
	"honnef.co/go/tools":              "https://github.com/dominikh/go-tools",
	"gopkg.in/DataDog/dd-trace-go.v1": "https://github.com/DataDog/dd-trace-go",
	"gotest.tools":                    "https://github.com/gotestyourself/gotest.tools",
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
	"golang.org/x/xerrors":  "BSD-3-Clause",
	"golang.org/x/oauth2":   "BSD-3-Clause",
	"google.golang.org/api": "BSD-3-Clause",
	"cloud.google.com/go":   "NOASSERTION",
}
