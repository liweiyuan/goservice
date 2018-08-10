package proxy

/*func ProxyMiddleware(ctx context.Context,instance string,logger log.Logger) service.ServiceMiddleware{

}*/
/*
func ProxyMiddleware(proxyURL string) service.ServiceMiddleware{
	return func(next service.StringService) service.StringService {
		return Proxymw{next,makeUppercaseEndpoint(proxyURL)}
	}
}
func makeUppercaseEndpoint(proxyURL string) endpoint.Endpoint {
	return httptransport.NewClient("")
}
*/

//StringService的 proxy
/*type Proxymw struct {
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
}*/
