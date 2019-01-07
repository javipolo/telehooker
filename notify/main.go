package notify

import (
	"fmt"
)

func justPrint(s string) {
	fmt.Println(s)
}

// Send notification
func Send(source string, msg string) {
	justPrint(msg)
}
