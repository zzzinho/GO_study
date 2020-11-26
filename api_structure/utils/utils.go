package utils

import (
	"regexp"

	"github.com/study/api_structure/db"
)

const (
	emailRegex = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
)

func ValidateUser(user db.Register, err []string) []string {
	emailCheck := regexp.MustCompile(emailRegex).MatchString(user.Email)
	if emailCheck != true {
		err = append(err, "Invalid email")
	}
	if len(user.Password) < 4 {
		err = append(err, "Invalid password, Password should be more than 4 characters")
	}
	if len(user.Name) < 1 {
		err = append(err, "Invalid name, please enter a name")
	}

	return err
}

func ValidatePasswordReset(resetPassword db.ResetPasword) (bool, string) {
	if len(resetPassword.Password) < 4 {
		return false, "Invalid password, Password should be more than 4 characters"
	}
	if resetPassword.Password != resetPassword.ConfirmPassword {
		return false, "Password reset failed, password must be match"
	}
	return true, ""
}
