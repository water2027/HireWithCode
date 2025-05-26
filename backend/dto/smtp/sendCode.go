package smtp

import (
	"fmt"
	"github.com/dlclark/regexp2"
)

type SendCodeRequest struct {
	Email string `json:"email"`
}

func (scr *SendCodeRequest) Examine() error {
	reg := regexp2.MustCompile(`^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`, 0)
	isMatch, _ := reg.MatchString(scr.Email)
	if !isMatch {
		return fmt.Errorf("邮箱不合格")
	}

	return nil
}
