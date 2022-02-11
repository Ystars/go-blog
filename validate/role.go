package validate

type RoleValidate struct {
	Common
	User string `json:"user" binding:"required"`
	Role string `json:"role" binding:"required"`
}
