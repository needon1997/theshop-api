package validation

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"regexp"
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

const MOBILE_FORMAT_1 = "^([1-9]{1})([0-9]{9})$"  //10 digits mobile
const MOBILE_FORMAT_2 = "^([1-9]{1})([0-9]{10})$" //11 digits mobile

func ValidateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	matched, _ := regexp.MatchString(MOBILE_FORMAT_1, mobile)
	if matched {
		return matched
	}
	matched, _ = regexp.MatchString(MOBILE_FORMAT_2, mobile)
	return matched
}
