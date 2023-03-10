package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	MAX_CAPACITY     = 10
	PREDEFINED_INPUT = "9 2 4 10 6 3 1 5 7 8"
)

func swap(sample []int, index int) {
	current := index
	after := current + 1

	currentValue := sample[current]
	afterValue := sample[after]

	if currentValue > afterValue {
		sample[after] = currentValue
		sample[current] = afterValue
	}
}

func bubbleSort(sample []int) {
	sampleLength := len(sample) - 1

	for i := 0; i < sampleLength; i++ {
		for j := 0; j < sampleLength; j++ {
			swap(sample, j)
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Please press ENTER to use the predefined sample [%s]\n", PREDEFINED_INPUT)
	fmt.Printf("OR type up to %d space-separated integers: ", MAX_CAPACITY)

	scanner.Scan()

	var input string

	input = scanner.Text()
	if input == "" {
		input = PREDEFINED_INPUT
	}

	data := strings.Split(input, " ")

	if len(data) > MAX_CAPACITY {
		log.Fatalf("Cannot process more than %d numbers.", MAX_CAPACITY)
	}

	sample := make([]int, 0, MAX_CAPACITY)

	for _, v := range data {
		number, _ := strconv.Atoi(v)
		sample = append(sample, number)
	}

	bubbleSort(sample)

	log.Println(sample)
}
