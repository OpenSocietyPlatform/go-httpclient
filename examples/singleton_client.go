package examples

import (
	"net/http"
	"time"

	"github.com/OpenSocietyPlatform/go-httpclient/gohttp"
	"github.com/OpenSocietyPlatform/go-httpclient/gomime"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() gohttp.Client {
	headers := make(http.Header)
	headers.Set(gomime.HeaderContentType, gomime.ContentTypeJSON)

	client := gohttp.NewBuilder().
		SetHeaders(headers).
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(3 * time.Second).
		// SetHttpClient(&currentClient).
		SetUserAgent("osp-server").
		Build()

	return client
}

func doSomething() {

}
