package response

import "github.com/liwss/kubernetesAdmin/server/model/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
