package controller

import (
	"backend/dto"
	"backend/dto/smtp"
	"backend/dto/work"
	"backend/service"

	"github.com/gin-gonic/gin"
)

// 前端只有两个提交表单的地方, 两个接口应该就差不多了. 再加一个看的接口和就差不多了

type WorkController struct {
	workService *service.WorkService
}

func NewWorkController(workService *service.WorkService) *WorkController {
	return &WorkController{
		workService: workService,
	}
}

// 接受挑战
func (wc *WorkController) AcceptChallenge(c *gin.Context) {
	req := work.AcceptChallengeRequest{}
	err := dto.BindData(c, &req)
	if err != nil {
		dto.ErrorResponse(c, dto.WithCode(dto.AttrError), dto.WithMessage(err.Error()))
		return
	}

	result, err := wc.workService.AcceptChallenge(req)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
		return
	}

	dto.SuccessResponse(c, dto.WithData(result))
}

// 提交作品
func (wc *WorkController) SubmitWork(c *gin.Context) {
	req := work.SubmitWorkRequest{}
	err := dto.BindData(c, &req)
	if err != nil {
		dto.ErrorResponse(c, dto.WithCode(dto.AttrError), dto.WithMessage(err.Error()))
		return
	}

	result, err := wc.workService.SubmitWork(req)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
		return
	}

	dto.SuccessResponse(c, dto.WithData(result))
}

// 查看作品
func (wc *WorkController) ViewWork(c *gin.Context) {
	req := work.ViewWorkRequest{}
	err := dto.BindData(c, &req)
	if err != nil {
		dto.ErrorResponse(c, dto.WithCode(dto.AttrError))
		return
	}

	result, err := wc.workService.ViewWork(req)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
		return
	}

	dto.SuccessResponse(c, dto.WithData(result))
}

// 发送验证码
func (wc *WorkController) SendEmail(c *gin.Context) {
	req := smtp.SendCodeRequest{}
	err := dto.BindData(c, &req)
	if err != nil {
		dto.ErrorResponse(c, dto.WithCode(dto.AttrError), dto.WithMessage(err.Error()))
		return
	}

	result, err := wc.workService.SendCode(req)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
		return
	}

	dto.SuccessResponse(c, dto.WithData(result))
}
