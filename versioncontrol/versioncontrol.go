package versioncontrol

import "net/http"

//Client holds the necessary items to make api calls to various version control providers
//TODO Would be nice to attach repo urls and paths to this object at some point
type Client struct {
	HttpClient *http.Client
	Token string
}
