package dto

type ForgotPasswordRedisObj struct {
	UserId      int32  `json:"userId"`
	Username    string `json:"username"`
	TokenString string `json:"token"`
}
