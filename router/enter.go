package router

import (
	"github.com/liwss/kubernetesAdmin/server/router/autocode"
	"github.com/liwss/kubernetesAdmin/server/router/example"
	"github.com/liwss/kubernetesAdmin/server/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Example  example.RouterGroup
	Autocode autocode.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
