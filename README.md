##Dep Report
This repo defines the tool that the b5 team uses to build a comprehensive report of our dependencies for review by security and other stakeholders.

###Running Locally or in CI
Ensure that you have 3 environment variables set:
 - `GITHUB_OAUTH_TOKEN` - String, used as credentials for the API calls to github
 - `DEP_REPORT_PRODUCT` - String, defines which product is using the report on a given run
 - `DEP_WEBHOOK` - String, defines the slack webhook url that we use to notify the team if a gerrit dependency license is not provided 
 
###Github Actions
A very simple implementation of github actions are configured on this repo. At this time, it only runs `go test -v ./...` but could be expanded upon in the future. It also protects the master branch from merges when the tests don't pass.
The code for this is defined at `/.github/workflows/go.yml`.
###Failure Notification
Based on guidance from security, we need to take action between releases when a dependency license fails to be provided in the report.
The existing logic allows a release pipeline to pass, but notifies us in the slack channel "dep-issues" so that we can create an issue and correct the problem ASAP.
If the slack notification fails, the pipeline fails.
