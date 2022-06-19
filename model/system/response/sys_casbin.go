package response

import (
	"github.com/liwss/kubernetesAdmin/server/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
