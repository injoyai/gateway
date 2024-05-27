package push

import "github.com/injoyai/goutil/net/http"

func NewHTTP() *HTTP {
	return &HTTP{http.NewClient()}
}

type HTTP struct {
	*http.Client
}

func (this *HTTP) Publish(url string, data interface{}) error {
	return http.Url(url).SetClient(this.Client).SetBody(data).Post().Err()
}
