package validators

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

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
