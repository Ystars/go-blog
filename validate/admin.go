package validate

type AdminCommon struct {
	Common
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
