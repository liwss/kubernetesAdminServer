package service

import (
	"github.com/liwss/kubernetesAdmin/server/service/autocode"
	"github.com/liwss/kubernetesAdmin/server/service/example"
	"github.com/liwss/kubernetesAdmin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	ExampleServiceGroup  example.ServiceGroup
	AutoCodeServiceGroup autocode.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
