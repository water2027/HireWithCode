package work

import (
	"time"
)

type ViewWorkRequest struct {
	Limit  uint64 `json:"limit"`
	Offset uint64 `json:"offset"`
}

func (vwr *ViewWorkRequest) Examine() error {
	return nil
}

type ViewWorkResponse struct {
	ID        uint      `json:"id"`
	GithubId  string    `json:"github_id"`
	Email     string    `json:"email"`
	GithubUrl string    `json:"github_url"`
	OnlineUrl string    `json:"online_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
