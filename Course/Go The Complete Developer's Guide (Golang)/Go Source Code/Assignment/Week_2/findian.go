package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Please input a string:")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputString := scanner.Text()

	lowerString := strings.ToLower(inputString)

	formatString := strings.ReplaceAll(lowerString, " ", "")
	if strings.Contains(formatString, "a") {
		if strings.HasPrefix(formatString, "i") {
			if strings.HasSuffix(formatString, "n") {
				fmt.Println("Found!")
				os.Exit()
			}
		}
	}

	fmt.Println("Not Found!")
}
