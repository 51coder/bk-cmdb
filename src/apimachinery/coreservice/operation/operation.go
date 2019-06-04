package operation

import (
	"context"
	"net/http"

	"configcenter/src/apimachinery/rest"
	"configcenter/src/common/metadata"
)

type OperationClientInterface interface {
	ModelInstAggregate(ctx context.Context, h http.Header, data interface{}) (resp *metadata.AggregateStringResponse, err error)
	AggregateBizHost(ctx context.Context, h http.Header, data interface{}) (resp *metadata.AggregateIntResponse, err error)
	CommonAggregate(ctx context.Context, h http.Header, data metadata.ChartOption) (resp *metadata.Response, err error)
	SearchInstCount(ctx context.Context, h http.Header, data interface{}) (resp *metadata.Response, err error)
	CreateOperationChart(ctx context.Context, h http.Header, data interface{}) (resp *metadata.Response, err error)
	SearchOperationChart(ctx context.Context, h http.Header, data interface{}) (resp *metadata.SearchChartConfig, err error)
	DeleteOperationChart(ctx context.Context, h http.Header, id string) (resp *metadata.Response, err error)
	UpdateOperationChart(ctx context.Context, h http.Header, data interface{}) (resp *metadata.Response, err error)
	SearchOperationChartData(ctx context.Context, h http.Header, data interface{}) (resp *metadata.Response, err error)
	UpdateOperationChartPosition(ctx context.Context, h http.Header, data interface{}) (resp *metadata.Response, err error)
}

func NewOperationClientInterface(client rest.ClientInterface) OperationClientInterface {
	return &operation{client: client}
}

type operation struct {
	client rest.ClientInterface
}
