package service

import (
	"log"
	"regexp"
)

func matchPattern(userStr string, pattern regexp.Regexp) bool {

	return pattern.MatchString(userStr)
}

// patterns
type RegexDetails struct {
	Pattern     *regexp.Regexp
	Description string
}

func emailRegex() *RegexDetails {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatal("Error compiling regex:", err)
	}
	return &RegexDetails{Pattern: re, Description: "Must Be A Valid Email"}
}
