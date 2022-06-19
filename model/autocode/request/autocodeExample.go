// 自动生成模板SysDictionaryDetail
package request

import (
	"github.com/liwss/kubernetesAdmin/server/model/autocode"
	"github.com/liwss/kubernetesAdmin/server/model/common/request"
)

// 如果含有time.Time 请自行import time包
type AutoCodeExampleSearch struct {
	autocode.AutoCodeExample
	request.PageInfo
}
