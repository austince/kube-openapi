package restfulv2adapter

import (
	"github.com/emicklei/go-restful/v2"
	"k8s.io/kube-openapi/pkg/common"
)

func AdaptWebServices(webServices []*restful.WebService) []common.RouteContainer {
	var containers []common.RouteContainer
	for _, ws := range webServices {
		containers = append(containers, &WebServiceAdapter{ws})
	}
	return containers
}

