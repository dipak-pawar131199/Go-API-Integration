package config

import "net/http"

const APIURL = "https://dummyjson.com/users"

func NewHTTPClient() *http.Client {
	return &http.Client{}
}
