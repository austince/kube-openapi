package restfuladapter

import (
	"github.com/emicklei/go-restful"
	"k8s.io/kube-openapi/pkg/common"
)

func AdaptWebServices(webServices []*restful.WebService) []common.RouteContainer {
	var containers []common.RouteContainer
	for _, ws := range webServices {
		containers = append(containers, &WebServiceAdapter{ws})
	}
	return containers
}

