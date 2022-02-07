package validate

type RoleValidate struct {
	Base
	User string `json:"user" binding:"required"`
	Role string `json:"role" binding:"required"`
}
