package response

import "github.com/liwss/kubernetesAdmin/server/model/example"

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
