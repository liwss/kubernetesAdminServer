package response

import "github.com/liwss/kubernetesAdmin/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
