package ui

import (
	"bufio"
	"fmt"
	"os"
)

func Input(title string, target *string) {
	reader := bufio.NewReader(os.Stdin)

	if *target == "" {
		fmt.Printf("%s: ", title)
	} else {
		fmt.Printf("%s: (%s) ", title, *target)
	}
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error: Failed to get user input")
		return
	}
	if input == "\n" {
		return
	}

	*target = input[:len(input)-1]
}
