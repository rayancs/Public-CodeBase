package service

type ValidatorResponse struct {
	Valid  bool
	Errors string
}

func ValidateEmailStr(str string) *ValidatorResponse {
	if !matchPattern(str, *emailRegex().Pattern) {
		return &ValidatorResponse{Valid: false, Errors: emailRegex().Description}
	}
	return &ValidatorResponse{Valid: true, Errors: "Passed"}
}
