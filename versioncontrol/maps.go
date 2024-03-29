package versioncontrol

var GithubRepoURLForPackage = map[string]string{
	"go.opencensus.io":                   "https://github.com/census-instrumentation/opencensus-go",
	"google.golang.org/grpc":             "https://github.com/grpc/grpc-go",
	"google.golang.org/genproto":         "https://github.com/googleapis/go-genproto",
	"google.golang.org/appengine":        "https://github.com/golang/appengine",
	"google.golang.org/api":              "https://github.com/googleapis/google-api-go-client",
	"google.golang.org/protobuf":         "https://github.com/golang/protobuf",
	"cloud.google.com/go":                "https://api.github.com/repos/googleapis/google-cloud-go",
	"gopkg.in/check.v1":                  "https://github.com/go-check/check",
	"gopkg.in/yaml.v2":                   "https://github.com/go-yaml/yaml",
	"gopkg.in/yaml.v3":                   "https://github.com/go-yaml/yaml",
	"gopkg.in/square/go-jose.v2":         "https://github.com/square/go-jose",
	"go.etcd.io/bbolt":                   "https://github.com/etcd-io/bbolt",
	"go.uber.org/atomic":                 "https://github.com/uber-go/atomic",
	"go.uber.org/multierr":               "https://github.com/uber-go/multierr",
	"go.uber.org/zap":                    "https://github.com/uber-go/zap",
	"gopkg.in/resty.v1":                  "https://github.com/go-resty/resty",
	"gopkg.in/ini.v1":                    "https://github.com/go-ini/ini",
	"gopkg.in/alecthomas/kingpin.v2":     "https://github.com/alecthomas/kingpin",
	"honnef.co/go/tools":                 "https://github.com/dominikh/go-tools",
	"gopkg.in/DataDog/dd-trace-go.v1":    "https://github.com/DataDog/dd-trace-go",
	"gotest.tools":                       "https://github.com/gotestyourself/gotest.tools",
	"golang.org/x/net":                   "https://github.com/golang/net",
	"aidanwoods.dev/go-paseto":           "https://github.com/aidantwoods/go-paseto",
	"go4.org/intern":                     "https://github.com/go4org/intern",
	"go4.org/unsafe/assume-no-moving-gc": "https://github.com/go4org/unsafe-assume-no-moving-gc",
	"inet.af/netaddr":                    "https://github.com/inetaf/netaddr",
	"go.opentelemetry.io/otel":           "https://github.com/open-telemetry/opentelemetry-go",
	"go.opentelemetry.io/otel/trace":     "https://github.com/open-telemetry/opentelemetry-go/trace",
}

var GerritRepoURLForPackage = map[string]string{
	"cloud.google.com/go": "https://code-review.googlesource.com/projects/gocloud",
}
var licenseForRepo = map[string]string{
	"golang.org/x/crypto":   "BSD-3-Clause",
	"golang.org/x/sync":     "BSD-3-Clause",
	"golang.org/x/image":    "BSD-3-Clause",
	"golang.org/x/sys":      "BSD-3-Clause",
	"golang.org/x/text":     "BSD-3-Clause",
	"golang.org/x/tools":    "BSD-3-Clause",
	"golang.org/x/xerrors":  "BSD-3-Clause",
	"golang.org/x/oauth2":   "BSD-3-Clause",
	"google.golang.org/api": "BSD-3-Clause",
	"cloud.google.com/go":   "NOASSERTION",
}
