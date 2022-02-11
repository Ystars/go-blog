package validate

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"goblog/utils/maps"
)

type Common struct{}

// BindJsonValidate 校验
func (e *Common) BindJsonValidate(obj interface{}, c *gin.Context) error {
	if err := c.ShouldBindJSON(&obj); err != nil {
		valid, ok := err.(validator.ValidationErrors)
		if !ok {
			return err
		}
		var transfer maps.Transfer
		return fmt.Errorf(transfer.MapToJson(valid.Translate(trans)))
	}

	return nil
}
