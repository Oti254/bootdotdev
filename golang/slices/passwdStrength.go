package main

func isValidPassword(password string) bool{
	isCorrectLen := false
	hasDigit := false
	hasUpper := false
	if len(password) >= 5 && len(password) <= 12{
		isCorrectLen = true
	}
	for _, passchar := range password {
		if (passchar >= 'A' && passchar <= 'Z') {
			hasUpper = true
		}
		if (passchar >= '0' && passchar <= '9') {
			hasDigit = true
		}
	}
	return isCorrectLen && hasUpper && hasDigit
}
