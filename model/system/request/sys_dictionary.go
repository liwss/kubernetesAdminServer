package request

import (
	"github.com/liwss/kubernetesAdmin/server/model/common/request"
	"github.com/liwss/kubernetesAdmin/server/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
