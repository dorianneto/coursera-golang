package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

func main() {
	var input string
	var counter int

	list := make([]int, 3)

	for {
		fmt.Print("Type a number or press X to exit: ")
		fmt.Scan(&input)

		if input == "x" {
			break
		}

		number, err := strconv.Atoi(input)

		if err != nil {
			log.Print(err)
			fmt.Println("Wrong number, try again.")
			continue
		}

		if counter < 3 {
			list[0] = number
		} else {
			list = append(list, number)
		}

		sort.Ints(list)

		counter++

		log.Printf("len=%d cap=%d %v \n\n", len(list), cap(list), list)
	}
}
