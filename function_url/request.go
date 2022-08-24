package function_url

import (
	"bytes"
	"github.com/aws/aws-lambda-go/events"
	"io"
	"net/http"
	"net/url"
)

func NewRequest(in events.LambdaFunctionURLRequest) *http.Request {
	inHttp := in.RequestContext.HTTP
	u := url.URL{
		Scheme:   inHttp.Protocol,
		Host:     in.RequestContext.DomainName,
		Path:     in.RawPath,
		RawQuery: in.RawQueryString,
	}
	u.RawPath = u.EscapedPath()
	req := &http.Request{
		Method:     inHttp.Method,
		URL:        &u,
		Header:     http.Header{},
		RequestURI: u.String(),
		Body:       io.NopCloser(bytes.NewBufferString(in.Body)),
	}
	for k, v := range in.Headers {
		req.Header.Set(k, v)
	}
	return req
}
