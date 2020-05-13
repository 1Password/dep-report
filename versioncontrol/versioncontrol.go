package versioncontrol

import (
	"net/http"
)

//Client holds the necessary items to make api calls to various version control providers
type Client struct {
	HttpClient *http.Client
	Token string
	//This webhook url is used to notify us when we fail to get a dependency's license
	SlackWebhook string
}
