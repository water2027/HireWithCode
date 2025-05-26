package work

import (
	"fmt"
	"strings"
	"github.com/dlclark/regexp2"
)

type AcceptChallengeRequest struct {
	GithubId string `json:"github_id"`
	Email    string `json:"email"`
	Code     string `json:"code"`
}

func (acr *AcceptChallengeRequest) Examine() error {
	if strings.TrimSpace(acr.GithubId) == "" {
		return fmt.Errorf("id为空")
	}
	if strings.TrimSpace(acr.Email) == "" {
		return fmt.Errorf("邮箱为空")
	}
	if strings.TrimSpace(acr.Code) == "" {
		return fmt.Errorf("验证码为空")
	}
	reg := regexp2.MustCompile(`^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`, 0)
	isMatch, _ := reg.MatchString(acr.Email)
	if !isMatch {
		return fmt.Errorf("邮箱不合格")
	}
	return nil
}
