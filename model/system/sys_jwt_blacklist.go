package system

import (
	"github.com/liwss/kubernetesAdmin/server/global"
)

type JwtBlacklist struct {
	global.GVA_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
