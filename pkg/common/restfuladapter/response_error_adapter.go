package restfuladapter

import (
	"github.com/emicklei/go-restful"
	"k8s.io/kube-openapi/pkg/common"
)

var _ common.StatusCodeResponse = &ResponseErrorAdapter{}

type ResponseErrorAdapter struct {
	Err restful.ResponseError
}

func (r *ResponseErrorAdapter) Message() string {
	return r.Err.Message
}

func (r *ResponseErrorAdapter) Model() interface{} {
	return r.Err.Model
}

func (r *ResponseErrorAdapter) Code() int {
	return r.Err.Code
}
