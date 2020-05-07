package slack

import (
	"github.com/jarcoal/httpmock"
	"testing"
)

func TestFailureNotify(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://hooks.slack.com/services/test/test/test", httpmock.NewStringResponder(200, "ok"))

	product := "test-product"
	dep := "test"
	webhookURL := "https://hooks.slack.com/services/test/test/test"

	if err := FailureNotify(product, dep, webhookURL); err != nil {
		t.Errorf("FailureNotify failed: %v", err)
	}
}
