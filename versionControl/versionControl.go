package versionControl

import (
	_ "golang.org/x/net/idna"
	"net/http"
	"time"
)

var Client = &http.Client{Timeout: 5 * time.Second}

var RepoURLForPackage = map[string]string{
	"go.opencensus.io":            "https://github.com/census-instrumentation/opencensus-go",
	"google.golang.org/grpc":      "https://github.com/grpc/grpc-go",
	"google.golang.org/genproto":  "https://github.com/googleapis/go-genproto",
	"google.golang.org/appengine": "https://github.com/golang/appengine",
	"google.golang.org/api":       "https://code-review.googlesource.com/projects/google-api-go-client",
	"cloud.google.com/go":         "https://code-review.googlesource.com/projects/gocloud",
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
