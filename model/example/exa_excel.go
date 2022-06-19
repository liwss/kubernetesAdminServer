package example

import "github.com/liwss/kubernetesAdmin/server/model/system"

type ExcelInfo struct {
	FileName string               `json:"fileName"` // 文件名
	InfoList []system.SysBaseMenu `json:"infoList"`
}
