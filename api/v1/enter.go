package v1

import (
	"github.com/liwss/kubernetesAdmin/server/api/v1/autocode"
	"github.com/liwss/kubernetesAdmin/server/api/v1/example"
	"github.com/liwss/kubernetesAdmin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	ExampleApiGroup  example.ApiGroup
	AutoCodeApiGroup autocode.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
