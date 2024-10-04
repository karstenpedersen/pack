package utils

import (
	"fmt"
	"bufio"
	"os"
)


func Input(title string, def string) string {
	reader := bufio.NewReader(os.Stdin)

	if def == "" {
		fmt.Printf("%s: ", title)
	} else {
		fmt.Printf("%s(%s): ", title, def)
	}
	input, err := reader.ReadString('\n')
	if err != nil || input == "\n" {
		return def
	}

	return input[:len(input) - 1]
}