package user_auth

type APILoginInput struct {
	Username string `form:"username" json:"username" binding:"required"`
	Passwd   string `form:"passwd" json:"passwd" bdinding:"required"`
}