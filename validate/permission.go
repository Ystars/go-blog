package validate

type PermissionValidate struct {
	Base
	User   string `json:"user" binding:"required"`
	Rule   string `json:"rule" binding:"required"`
	Method string `json:"method" binding:"required"`
}
