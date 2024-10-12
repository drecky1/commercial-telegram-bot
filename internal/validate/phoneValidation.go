package validate

import (
	"regexp"
	"tg_contour_bot/utils"
)

func IsValidRussianPhoneNumber(phone string) bool {
	re := regexp.MustCompile(utils.PhoneNumberRegex)
	return re.MatchString(phone)
}
