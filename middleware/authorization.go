package middleware

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"goblog/model"
	"goblog/utils/enum"
)

func Authorization(obj string, act string) gin.HandlerFunc {
	return func(c *gin.Context) {
		adapter, _ := gormadapter.NewAdapterByDBWithCustomTable(model.Db, &model.CasbinRule{})

		val, existed := c.Get("username")
		if !existed {
			code := enum.ERROR_USER_MIDDLEWARE_AUTH
			c.AbortWithStatusJSON(code, gin.H{"message": enum.GetCodeMsg(code)})
			return
		}

		if obj == "" {
			obj = c.Request.URL.RequestURI()
		}

		if act == "" {
			act = c.Request.Method
		}

		ok, err := enforce(val.(string), obj, act, adapter)
		if err != nil {
			code := enum.ERROR_USER_AUTHORIZING
			c.AbortWithStatusJSON(code, gin.H{"message": enum.GetCodeMsg(code)})
			return
		}
		if !ok {
			code := enum.FORBIDDEN
			c.AbortWithStatusJSON(code, gin.H{"message": enum.GetCodeMsg(code)})
			return
		}
		c.Next()
	}
}

func enforce(sub string, obj string, act string, adapter persist.Adapter) (bool, error) {
	enforcer, err := casbin.NewEnforcer("config/rbac_model.conf", adapter)
	if err != nil {
		return false, fmt.Errorf(enum.GetCodeMsg(enum.ERROR_CASBIN_LOCAL)+": %w", err)
	}
	// Load policies from DB dynamically
	err = enforcer.LoadPolicy()
	if err != nil {
		return false, fmt.Errorf(enum.GetCodeMsg(enum.ERROR_CASBIN_DYNAMICALLY)+": %w", err)
	}
	ok, err := enforcer.Enforce(sub, obj, act)
	return ok, err
}
