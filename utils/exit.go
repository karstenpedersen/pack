package utils

import (
	"fmt"
	"os"
)

func Exit(message ...any) {
	fmt.Println(message...)
	os.Exit(1)
}
