package client

import "github.com/injoyai/goutil/net/http"

func NewHTTP() (*HTTP, error) {
	return &HTTP{http.NewClient()}, nil
}

type HTTP struct {
	*http.Client
}

func (this *HTTP) Publish(url string, data interface{}) error {
	return http.Url(url).SetClient(this.Client).SetBody(data).Post().Err()
}
