package restfulv2adapter

import (
	"github.com/emicklei/go-restful/v2"
	"k8s.io/kube-openapi/pkg/common"
)

var _ common.Route = &RouteAdapter{}

type RouteAdapter struct {
	Route restful.Route
}

func (r *RouteAdapter) StatusCodeResponses() map[int]common.StatusCodeResponse {
	errs := make(map[int]common.StatusCodeResponse, len(r.Route.ResponseErrors))
	for status, rErr := range r.Route.ResponseErrors {
		errs[status] = &ResponseErrorAdapter{rErr}
	}

	return errs
}

func (r *RouteAdapter) OperationName() string {
	return r.Route.Operation
}

func (r *RouteAdapter) Method() string {
	return r.Route.Method
}

func (r *RouteAdapter) Path() string {
	return r.Route.Path
}

func (r *RouteAdapter) Parameters() []common.Parameter {
	var params []common.Parameter
	for _, rParam := range r.Route.ParameterDocs {
		params = append(params, &ParamAdapter{*rParam})
	}
	return params
}

func (r *RouteAdapter) Description() string {
	return r.Route.Doc
}

func (r *RouteAdapter) Consumes() []string {
	return r.Route.Consumes
}

func (r *RouteAdapter) Produces() []string {
	return r.Route.Produces
}

func (r *RouteAdapter) Metadata() map[string]interface{} {
	return r.Route.Metadata
}

func (r *RouteAdapter) RequestPayloadSample() interface{} {
	return r.Route.ReadSample
}

func (r *RouteAdapter) ResponsePayloadSample() interface{} {
	return r.Route.WriteSample
}
