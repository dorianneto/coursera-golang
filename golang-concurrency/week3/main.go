package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	DEFAULT_VALUE  = "5 8 4 2 1 6 7 3 10 12 11 9"
	PARTITION_SIZE = 4
)

func main() {
	var rawData []int

	scan := bufio.NewScanner(os.Stdin)

	fmt.Printf("Write a series of integers separated by space or just press enter (default: %s): ", DEFAULT_VALUE)
	scan.Scan()

	input := scan.Text()
	if input == "" {
		input = DEFAULT_VALUE
	}

	for _, i := range strings.Split(input, " ") {
		number, err := strconv.Atoi(i)
		if err != nil {
			log.Fatalln("All the values MUST be a number")
		}

		rawData = append(rawData, number)
	}

	perPartition := int(math.Ceil(float64(len(rawData)) / PARTITION_SIZE))

	index := 0
	value := perPartition

	var output []int

	for i := 0; i < PARTITION_SIZE; i++ {
		c := make(chan []int)

		partition := rawData[index:value]

		go func(o *[]int, p []int, c chan []int) {
			*o = append(*o, partition...)

			c <- *o
		}(&output, partition, c)

		<-c

		index += perPartition
		value += perPartition
	}

	sort.Ints(output)

	fmt.Println(output)
}
