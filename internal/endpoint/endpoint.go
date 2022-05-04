package endpoint

import (
	"context"

	"golang-with-kafka/internal/entity"
	"golang-with-kafka/internal/service"

	"github.com/go-kit/kit/endpoint"
)

func MakeServiceEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, in interface{}) (interface{}, error) {
		req := in.(*entity.Request)

		return svc.Call(ctx, req)
	}
}
