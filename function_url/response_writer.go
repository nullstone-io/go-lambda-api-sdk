package function_url

import (
	"bytes"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

type ResponseWriter interface {
	http.ResponseWriter
	ToResponse() events.LambdaFunctionURLResponse
}

func NewResponseWriter() ResponseWriter {
	return &responseWriter{
		headers: http.Header{},
		body:    bytes.NewBufferString(""),
	}
}

var _ ResponseWriter = &responseWriter{}

type responseWriter struct {
	headers    http.Header
	body       *bytes.Buffer
	statusCode int
}

func (l *responseWriter) Header() http.Header {
	return l.headers
}

func (l *responseWriter) Write(i []byte) (int, error) {
	return l.body.Write(i)
}

func (l *responseWriter) WriteHeader(statusCode int) {
	l.statusCode = statusCode
}

func (l *responseWriter) ToResponse() events.LambdaFunctionURLResponse {
	headers := map[string]string{}
	for k := range l.headers {
		headers[k] = l.headers.Get(k)
	}
	return events.LambdaFunctionURLResponse{
		StatusCode:      l.statusCode,
		Headers:         headers,
		Body:            l.body.String(),
		IsBase64Encoded: false,
		Cookies:         make([]string, 0),
	}
}
