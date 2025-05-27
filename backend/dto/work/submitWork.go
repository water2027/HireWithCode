package work

import (
	"fmt"
	"github.com/dlclark/regexp2"
	"strings"
)

type SubmitWorkRequest struct {
	GithubUrl string `json:"github_url"`
	OnlineUrl string `json:"online_url"`
	Email     string `json:"email"`
	Code      string `json:"code"`
}

func (swr *SubmitWorkRequest) Examine() error {
	emailReg := regexp2.MustCompile(`^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`, 0)
	isMatch, _ := emailReg.MatchString(swr.Email)
	if !isMatch {
		return fmt.Errorf("无效的邮箱")
	}

	if !strings.HasPrefix(swr.GithubUrl, "http://") && !strings.HasPrefix(swr.GithubUrl, "https://") {
		return fmt.Errorf("无效的仓库链接") 
	}
	if !strings.HasPrefix(swr.OnlineUrl, "http://") && !strings.HasPrefix(swr.OnlineUrl, "https://") {
		return fmt.Errorf("无效的在线体验链接") 
	}

	return nil
}
