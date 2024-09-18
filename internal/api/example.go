package api

import (
	"context"
	"github.com/woaijssss/tros"
	trlogger "github.com/woaijssss/tros/logx"
	v1 "tros_example_server/internal/proto/console/v1"
)

type exampleController struct {
}

var ExampleController = new(exampleController)

func (e *exampleController) Init(atx tros.AppContext) error {
	v1.RegisterExampleServiceServer(atx.ServiceRegistrar(), e)
	atx.HTTPRouter().RegisterServiceHandler(v1.RegisterExampleServiceHandler)

	return nil
}

func (e *exampleController) GetExample(ctx context.Context, request *v1.GetExampleRequest) (*v1.GetExampleResponse, error) {
	trlogger.Infof(ctx, "request param: [%+v]", request.GetId())
	return &v1.GetExampleResponse{
		Value: 10,
	}, nil
}

func (e *exampleController) PostExample(ctx context.Context, request *v1.PostExampleRequest) (*v1.PostExampleResponse, error) {
	trlogger.Infof(ctx, "request param: [%+v]", request.GetId())
	return &v1.PostExampleResponse{
		Value: 10,
	}, nil
}
