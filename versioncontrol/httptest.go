package versioncontrol

import (
	"fmt"
	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
	"net/http"
	"time"
)

func SetupHTTPRecord(fileName string) (*recorder.Recorder, *http.Client, error){
	r, err := recorder.New("./testData/"+fileName)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to setup http recorder, %v", err)
	}

	r.AddFilter(func(i *cassette.Interaction) error {
		delete(i.Request.Headers, "Authorization")
		return nil
	})

	return r, &http.Client{
		Transport:     r,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       5 * time.Second,
	}, nil
}
