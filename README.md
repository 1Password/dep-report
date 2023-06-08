# dep-report

`dep-report` is a custom written tool used to report on the golang dependencies in an application.

It functions by reading the application `go.mod` file, and then interacting with the remote repositories for each dependency in order to build a report with entries formatted like:

```
    {
      "name": "github.com/DataDog/datadog-go",
      "source": "github",
      "License": "MIT",
      "website": "https://api.github.com/repos/DataDog/datadog-go",
      "installed": {
        "version": "1.4.1",
        "time": "2021-05-05T11:24:08Z",
        "commit": "fbbbcbc72f95c23c28bbfe2bf008a9958db049a2"
      },
      "latest": {
        "version": "v5.1.1",
        "time": "2022-05-05T16:04:48Z",
        "commit": "553de96e699a42be8b401607fbbbce81d4942790"
      }
    }
```

## Running the Tool

* In order to run the tool, you must first setup a [Github Personal Access Token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token)
* To verify the PAT is configured correctly, you can test out running the tool against it's own deps:
```
GITHUB_OAUTH_TOKEN=<your token> go run main.go
```
* If this works, then install the tool globally via `go install .`
* This tool must be run in the root directory of the application to be reported on (i.e. in the same location as `go.mod`)
```
> cd my/go/app
> GITHUB_OAUTH_TOKEN=<your token> dep-report
```

## Troubleshooting

### `Unable to determine repo source for...`

This is the most common issue encountered with this tool. 

The code for the tool relies on a [mapping for particular dependency repo sources](https://github.com/1Password/dep-report/blob/master/versioncontrol/maps.go). 

While this is not ideal, it is relatively easy to fix.

For example, a recent failure reported the following issue:

```
unable to generate report: failed to create report object from dependency: { v1.26.0 google.golang.org/protobuf}: unable to determine repo source for google.golang org/protobuf
```

And the fix required was simply to add a mapping for the `google.golang.org/protobuf` dependency to it's source repo on github: `https://github.com/golang/protobuf`

```
	"google.golang.org/protobuf":      "https://github.com/golang/protobuf",
```
