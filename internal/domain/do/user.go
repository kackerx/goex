package do

import (
	"time"

	"goex1/internal/domain/enum"
)

type User struct {
	ID        uint            `json:"id,omitempty"`
	UserName  string          `json:"user_name,omitempty"`
	Password  string          `json:"password,omitempty"`
	NickName  string          `json:"nick_name,omitempty"`
	Email     string          `json:"email,omitempty"`
	Slogan    string          `json:"slogan,omitempty"`
	Status    enum.UserStatus `json:"status,omitempty"`
	Gender    enum.UserGender `json:"gender"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

type SessionInfo struct {
	UserID       uint   `json:"user_id,omitempty"`
	SessionID    string `json:"session_id,omitempty"`
	Platform     string `json:"platform,omitempty"`
	Phone        string `json:"phone,omitempty"`
	Email        string `json:"email,omitempty"`
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type TokenInfo struct {
	AccessToken   string    `json:"access_token,omitempty"`
	RefreshToken  string    `json:"refresh_token,omitempty"`
	Duration      int64     `json:"duration,omitempty"`
	SrvCreateTime time.Time `json:"srv_create_time"`
}
