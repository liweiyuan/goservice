package proxy

import (
	"goservice/stringservice/service"
	"github.com/go-kit/kit/endpoint"
	"context"
	"goservice/stringservice/transport"
	"errors"
)

//StringService的 proxy
type Proxymw struct {
	Next              service.StringService // Serve most requests via this service...
	UppercaseEndpoint endpoint.Endpoint     // ...except Uppercase, which gets served by this endpoint
}

//proxy 方法
func (mw Proxymw) Uppercase(ctx context.Context, s string) (string, error) {
	response, err := mw.UppercaseEndpoint(ctx, transport.UppercaseRequest{S: s})
	if err != nil {
		return "", nil
	}

	resp := response.(transport.UppercaseResponse)

	if resp.Err != "" {
		return resp.V, errors.New(resp.Err)
	}
	return resp.V, nil
}
