package function_url

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

func Handle(ctx context.Context, event events.LambdaFunctionURLRequest, handler http.Handler) (events.LambdaFunctionURLResponse, error) {
	if event.RequestContext.HTTP.Method == "" {
		return events.LambdaFunctionURLResponse{
			StatusCode:      http.StatusNotFound,
			Headers:         map[string]string{},
			Body:            "",
			IsBase64Encoded: false,
			Cookies:         nil,
		}, nil
	}

	req := NewRequest(event)
	req = req.WithContext(ctx)
	wr := NewResponseWriter()
	handler.ServeHTTP(wr, req)
	return wr.ToResponse(), nil
}
