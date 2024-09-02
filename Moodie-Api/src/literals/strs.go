package literals

func PortNumber() string   { return ":8080" }
func UserEndPoint() string { return "/user" }
func Bearer() string       { return "Bearer" }

// We Faced A Problem While Logging You in \nPlease Try again , if the problem conscists contact Customer Service
func GeneralError(doing string) string {
	return "We Faced A Problem While " + doing + ". Please Try again, if the problem persists, contact Customer Service."
}
