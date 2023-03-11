package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var scanner *bufio.Scanner

func prompt(value *float64, message string) {
	fmt.Print(message)

	for scanner.Scan() {
		v, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Printf("Wrong type, float expected. %s", message)
			continue
		}

		*value = v

		break
	}
}

func genDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	s := func(t float64) float64 {
		// s = ½ a t2 + vot + so
		return ((0.5 * a) * (math.Pow(t, 2))) + (v0 * t) + s0
	}

	return s
}

func main() {
	var a, v0, s0, t float64

	scanner = bufio.NewScanner(os.Stdin)

	prompt(&a, "Type the acceleration: ")
	prompt(&v0, "Type the initial velocity: ")
	prompt(&s0, "Type the initial displacement: ")

	fn := genDisplaceFn(a, v0, s0)

	prompt(&t, "Do you want to see the displacement after how many seconds? ")

	fmt.Printf("Result: %.2f", fn(t))
}
