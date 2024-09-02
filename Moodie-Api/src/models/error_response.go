package models

type ErrorResponse struct {
	Code    int
	Error   []string
	Details []string
}
