package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var input string

	const (
		FIRST_LETTER   = "i"
		CONTAIN_LETTER = "a"
		LAST_LETTER    = "n"
	)

	fmt.Print("Type a phrase:")

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() == false {
		fmt.Println("Scanner has failed")
		return
	}

	input = scanner.Text()

	inputCaseInsensitive := strings.ToLower(input)

	log.Println(inputCaseInsensitive)

	hasPrefix := strings.HasPrefix(inputCaseInsensitive, FIRST_LETTER)
	hasSuffix := strings.HasSuffix(inputCaseInsensitive, LAST_LETTER)
	doesContain := strings.Contains(inputCaseInsensitive, CONTAIN_LETTER)

	if hasPrefix && hasSuffix && doesContain {
		fmt.Println("Found!")
		return
	}

	fmt.Println("Not Found!")
}
