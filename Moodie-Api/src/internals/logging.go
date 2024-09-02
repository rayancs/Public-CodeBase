package internals

import "fmt"

func RedLog(params string) {
	fmt.Println("\033[31m" + params + "\033[0m")
}
func GreenLog(params string) {
	fmt.Println("\033[32m" + params + "\033[0m")
}
func BlueLog(params string) {
	fmt.Println("\033[34m" + params + "\033[0m")
}
func PurpleLog(params string) {
	fmt.Println("\033[35m" + params + "\033[0m")
}
