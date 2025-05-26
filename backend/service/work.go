package service

import (
	"backend/database"
	"backend/dto/smtp"
	workDto "backend/dto/work"
	workModel "backend/model/work"
	"backend/utils"
	"fmt"
)

// repo层就算了吧, 这么小的东西也懒得写了. 直接用model吧
type WorkService struct {
}

func (ws *WorkService) SendCode(req smtp.SendCodeRequest) (string, error) {
	randomCode := utils.GetRandomCode()

	smtpService := NewSmtpService()
	err := smtpService.SendCode(req.Email, randomCode)
	if err != nil {
		return "", err
	}
	return "发送成功", nil
}

func (ws *WorkService) AcceptChallenge(req workDto.AcceptChallengeRequest) (string, error) {
	// 事已至此, 先验证邮箱吧
	smtpService := NewSmtpService()
	ok, err := smtpService.VerifyCode(req.Email, req.Code)
	if err != nil {
		return "", fmt.Errorf("服务器错误")
	}
	if !ok {
		return "", fmt.Errorf("验证码错误")
	}

	// 验证完了就存起来吧
	db := database.GetMysqlDb()
	modelDb := db.Model(&workModel.Work{})
	var emailCount int64 = 0
	modelDb.Where("email = ?", req.Email).Count(&emailCount)
	if emailCount > 0 {
		return "", fmt.Errorf("邮箱注册过了")
	}
	var idCount int64 = 0
	modelDb.Where("github_id = ?", req.GithubId).Count(&idCount)
	if idCount > 0 {
		return "", fmt.Errorf("这个id注册过了")
	}

	work := workModel.Work{
		Email:    req.Email,
		GithubId: req.GithubId,
	}
	if err := modelDb.Create(&work).Error; err != nil {
		return "", fmt.Errorf("服务器错误")
	}
	return "接收成功", nil
}

func (ws *WorkService) SubmitWork(req workDto.SubmitWorkRequest) (string, error) {
	smtpService := NewSmtpService()
	ok, err := smtpService.VerifyCode(req.Email, req.Code)
	if err != nil {
		return "", fmt.Errorf("服务器错误")
	}
	if !ok {
		return "", fmt.Errorf("验证码错误")
	}

	db := database.GetMysqlDb()
	modelDb := db.Model(&workModel.Work{})
	var work workModel.Work
	if err := modelDb.Where("email = ?", req.Email).First(&work).Error; err != nil {
		return "", fmt.Errorf("用户未注册，请先接受挑战")
	}

	work.GithubUrl = req.GithubUrl
	work.OnlineUrl = req.OnlineUrl

	if err := modelDb.Save(&work).Error; err != nil {
		return "", fmt.Errorf("服务器错误")
	}

	return "提交成功", nil
}

func (ws *WorkService) ViewWork(req workDto.ViewWorkRequest) ([]workDto.ViewWorkResponse, error) {
	db := database.GetMysqlDb()
	modelDb := db.Model(&workModel.Work{})

	var works []workModel.Work
	query := modelDb.Order("created_at DESC")

	if req.Limit > 0 {
		query = query.Limit(int(req.Limit))
	}

	if req.Offset > 0 {
		query = query.Offset(int(req.Offset))
	}

	if err := query.Find(&works).Error; err != nil {
		return nil, fmt.Errorf("服务器错误")
	}

	var responses []workDto.ViewWorkResponse
	for _, work := range works {
		response := workDto.ViewWorkResponse{
			ID:        work.ID,
			GithubId:  work.GithubId,
			Email:     work.Email,
			GithubUrl: work.GithubUrl,
			OnlineUrl: work.OnlineUrl,
			CreatedAt: work.CreatedAt,
			UpdatedAt: work.UpdatedAt,
		}
		responses = append(responses, response)
	}

	return responses, nil
}
