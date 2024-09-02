package internals

import (
	"fmt"
	"os"
)

func ErrorExit(msg string, err error) {
	fmt.Printf("%s \n%v", msg, err)
	os.Exit(1)
}
