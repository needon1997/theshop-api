package validation

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

const ERROR string = "error"

func ValidateFormJSON(c *gin.Context, form interface{}) error {
	err := c.ShouldBindJSON(form)
	if err != nil {
		if _, ok := err.(validator.ValidationErrors); !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				ERROR: err.Error(),
			})
			return err
		}
		c.JSON(http.StatusOK, gin.H{
			ERROR: err.Error(),
		})
		return err
	}
	return nil
}
