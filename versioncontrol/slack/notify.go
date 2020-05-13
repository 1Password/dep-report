package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"time"
)

type slackRequestBody struct {
	Text string `json:"text"`
}

func FailureNotify(product string, dep string, webhookURL string) error{
	if webhookURL == "" {
		return errors.New("unable to post failure notification, webhookURL missing")
	}
	err := sendSlackNotification(webhookURL, fmt.Sprintf("Failed to retrieve license for product: %s, dependency: %s", product, dep))
	if err != nil {
		return errors.Wrapf(err, "unable to post failure notification to slack, product: %s, dependency: %s", product, dep)
	}
	return nil
}

func sendSlackNotification(webhookURL string, msg string) error{
	slackBody, _ := json.Marshal(slackRequestBody{Text: msg})
	req, err := http.NewRequest(http.MethodPost, webhookURL, bytes.NewBuffer(slackBody))
	if err != nil {
		return errors.Wrap(err, "unable to prepare request")
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "unable to make request to slack webhook")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrapf(err, "unable to read slack response body")
	}

	if string(body) != "ok" {
		return fmt.Errorf("non-ok response returned from Slack: %v", string(body))
	}

	return nil
}
