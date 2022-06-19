package request

import (
	"github.com/liwss/kubernetesAdmin/server/model/common/request"
	"github.com/liwss/kubernetesAdmin/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
