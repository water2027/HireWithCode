package route

import (
	"backend/controller"
	"backend/service"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	workController := controller.NewWorkController(&service.WorkService{})
	router.POST("/work/submit", workController.SubmitWork)
	router.POST("/work/view", workController.ViewWork)
	router.POST("/work/accept", workController.AcceptChallenge)
	router.POST("/work/send_code", workController.SendEmail)
}