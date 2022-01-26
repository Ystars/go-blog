package V1

import (
	"github.com/gin-gonic/gin"
	"goblog/model"
	"goblog/utils/enum"
	"goblog/utils/jwt"
	"goblog/utils/response"
	"net/http"
)

var adminModel model.Admin

type AdminApi struct{}

// Login 后台登录
func (t *AdminApi) Login(c *gin.Context) {
	_ = c.ShouldBindJSON(&adminModel)
	var token string

	code := adminModel.CheckLogin(adminModel.Username, adminModel.Password)

	if code == enum.SUCCESS {
		jwt.NewJWT().SetToken(c, jwt.SetTokenData{
			Username: adminModel.Username,
			ID:       int(adminModel.ID),
			Issuer:   "goblog",
		})
		return
	}
	response.OkWithData(response.R{
		"status":  code,
		"data":    adminModel.Username,
		"id":      adminModel.ID,
		"message": enum.GetCodeMsg(code),
		"token":   token,
	}, c)

}

func (t *AdminApi) Show(c *gin.Context) {
	_ = c.ShouldBindJSON(&adminModel)
	user, _ := c.Get("username")
	c.JSON(http.StatusOK, gin.H{
		"status":  enum.SUCCESS,
		"aaa":     112,
		"message": user,
	})
}
