package enum

import "time"

type UserStatus string

const (
	UserStatusNormal  UserStatus = "normal"
	UserStatusDisable UserStatus = "disable"
)

type UserGender string

const (
	UserGenderMale   = "male"
	UserGenderFeMale = "female"
)

var (
	userStatusMap = map[int]UserStatus{
		1: UserStatusNormal,
		2: UserStatusDisable,
	}

	userGenderMap = map[int]UserGender{
		1: UserGenderMale,
		2: UserGenderFeMale,
	}
)

func GetUserStatus(status int) UserStatus {
	return userStatusMap[status]
}

func GetUserGender(status int) UserGender {
	return userGenderMap[status]
}

const RedisKeyPrefix = "goex:"

const (
	RedisKeyAccessToken      = RedisKeyPrefix + "user:access_token:%s"
	RedisKeyRefreshToken     = RedisKeyPrefix + "user:refresh_token:%s"
	RedisKeySession          = RedisKeyPrefix + "user:session:%d"
	RedisKeyLockRefreshToken = RedisKeyPrefix + "user:lock_refresh_token:%s"
)

const (
	AccessTokenDuration            = 2 * time.Hour
	RefreshTokenDuration           = 24 * time.Hour * 10
	OldRefreshTokenHoldingDuration = 6 * time.Hour    // 刷新Token时老的RefreshToken保留的时间(用于发现refresh被窃取)
	PasswordTokenDuration          = 15 * time.Minute // 重置密码的验证Token的有效期
)
