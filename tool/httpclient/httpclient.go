package httpclient

import (
	"crypto/tls"
	"sync"

	"github.com/go-resty/resty/v2"
)

var httpclient *resty.Client
var lock sync.Mutex

func SetUp() {
	if httpclient == nil {
		lock.Lock()
		if httpclient == nil {
			httpclient = resty.New()
			httpclient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
		}
		lock.Unlock()
	}
}

func NewRequest() *resty.Request {
	if httpclient == nil {
		lock.Lock()
		if httpclient == nil {

			httpclient = resty.New()
			httpclient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
		}
		lock.Unlock()
	}
	return httpclient.R()
}
