package request

import (
	"github.com/liwss/kubernetesAdmin/server/model/autocode"
	"github.com/liwss/kubernetesAdmin/server/model/common/request"
)

type {{.StructName}}Search struct{
    autocode.{{.StructName}}
    request.PageInfo
}