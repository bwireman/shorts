package util

import (
	"fmt"
	"os"
)

func Exit(msg interface{}) {
	fmt.Printf("echo %s", msg)
	os.Exit(1)
}

func MaybeExit(msg interface{}) {
	if msg != nil {
		Exit(msg)
	}
}
