package V1

import (
	"github.com/gin-gonic/gin"
	"goblog/service"
	"goblog/utils/enum"
	"goblog/utils/response"
	"goblog/validate"
)

type PermissionApi struct{}

func (t *PermissionApi) Add(c *gin.Context) {
	var permissionValidate validate.PermissionValidate
	if err, ok := permissionValidate.Check(c, &permissionValidate); !ok {
		response.OkWithData(response.R{
			"status":  enum.ERROR,
			"message": err,
		}, c)
		return
	}
	_, err := service.NewAuthorize().AddPermissionForUser(permissionValidate.User, permissionValidate.Rule, permissionValidate.Method)
	if err != nil {
		response.OkWithData(response.R{
			"status":  enum.ERROR,
			"message": err,
		}, c)
		return
	}
	response.Ok(c)
}
