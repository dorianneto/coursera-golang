package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const FIELD_SIZE_LIMIT = 20

type Name struct {
	fname string
	lname string
}

func generateFieldValue(value string) string {
	if len(value) > FIELD_SIZE_LIMIT {
		return value[:FIELD_SIZE_LIMIT]
	}

	return value
}

func main() {
	var filename string
	var slice []Name

	fmt.Print("What file do you want to open? ")
	fmt.Scan(&filename)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Cannot open the file", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		name := strings.Split(scanner.Text(), " ")

		fname := generateFieldValue(name[0])
		lname := generateFieldValue(name[1])

		slice = append(slice, Name{
			fname: fname,
			lname: lname,
		})
	}

	for _, n := range slice {
		fmt.Printf("fname=%s lname=%s\n", n.fname, n.lname)
	}
}
