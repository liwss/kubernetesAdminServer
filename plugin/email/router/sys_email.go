package router

import (
	"github.com/gin-gonic/gin"
	"github.com/liwss/kubernetesAdmin/server/middleware"
	"github.com/liwss/kubernetesAdmin/server/plugin/email/api"
)

type EmailRouter struct {
}

func (s *EmailRouter) InitEmailRouter(Router *gin.RouterGroup) {
	emailRouter := Router.Use(middleware.OperationRecord())
	var EmailApi = api.ApiGroupApp.EmailApi.EmailTest
	var SendEmail = api.ApiGroupApp.EmailApi.SendEmail
	{
		emailRouter.POST("emailTest", EmailApi)  // 发送测试邮件
		emailRouter.POST("sendEmail", SendEmail) // 发送邮件
	}
}
