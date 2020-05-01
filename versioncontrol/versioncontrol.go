package versioncontrol

import "net/http"

//Client holds the necessary items to make api calls to various version control providers
type Client struct {
	HttpClient *http.Client
	Token string
}
