package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	var name, address string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("What's your name? ")
	scanner.Scan()
	name = scanner.Text()

	fmt.Print("What's your address? ")
	scanner.Scan()
	address = scanner.Text()

	data, err := json.Marshal(map[string]string{
		"name":    name,
		"address": address,
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(data))
}
