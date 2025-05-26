package work

import (
	"fmt"
	"strings"
)

type SubmitWorkRequest struct {
	GithubUrl string `json:"github_url"`
	OnlineUrl string `json:"online_url"`
	Email     string `json:"email"`
	Code      string `json:"code"`
}

func (swr *SubmitWorkRequest) Examine() error {
	if strings.TrimSpace(swr.GithubUrl) == "" {
		return fmt.Errorf("github地址为空")
	}
	if strings.TrimSpace(swr.OnlineUrl) == "" {
		return fmt.Errorf("在线体验地址为空")
	}
	return nil
}
