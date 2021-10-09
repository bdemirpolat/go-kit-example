package email

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	Send endpoint.Endpoint
}

// MakeEndpoints ...
func MakeEndpoints(svc Service, middlewares []endpoint.Middleware) Endpoints {
	return Endpoints{
		Send: setMiddlewares(makeSendEndpoint(svc), middlewares),
	}
}

// setMiddlewares
func setMiddlewares(e endpoint.Endpoint, m []endpoint.Middleware) endpoint.Endpoint {
	for _, middleware := range m {
		e = middleware(e)
	}
	return e
}

// makeSendEndpoint
func makeSendEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(sendRequest)
		if !ok {
			return nil, errors.New("request type is not a `sendRequest`")
		}

		res, err := s.Send(req)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}
