package examples

import (
	"time"

	"github.com/espana33/go-httpclient.git/gohttp"
)

var(
	httpClient = getHttpClient()
)

func getHttpClient() gohttp.Client{
	client:= gohttp.NewBuilder().
		SetConnectionTimeout(2*time.Second).
		SetResponseTimeout(3*time.Second).
		Build()
	return client
}