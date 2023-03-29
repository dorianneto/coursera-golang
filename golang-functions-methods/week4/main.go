package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
)

const (
	COMMAND_NEWANIMAL = "newanimal"
	COMMAND_QUERY     = "query"
)

var scanner *bufio.Scanner
var executed bool

type callback func()

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (c Cow) Eat() {
	fmt.Println("grass")
}
func (c Cow) Move() {
	fmt.Println("walk")
}
func (c Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct{}

func (c Bird) Eat() {
	fmt.Println("worms")
}
func (c Bird) Move() {
	fmt.Println("fly")
}
func (c Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct{}

func (c Snake) Eat() {
	fmt.Println("mice")
}
func (c Snake) Move() {
	fmt.Println("slither")
}
func (c Snake) Speak() {
	fmt.Println("hsss")
}

func prompt(input *[]string) error {
	fmt.Print(">")
	scanner.Scan()

	data := strings.Split(scanner.Text(), " ")
	if len(data) != 3 {
		return errors.New("Incorrect input format")
	}

	*input = data

	return nil
}

func run(shouldRun bool, callback callback) {
	if !shouldRun {
		return
	}

	callback()

	executed = true
}

func main() {
	var input []string

	animals := make(map[string]Animal)
	factory := map[string]Animal{
		"cow":   Cow{},
		"bird":  Bird{},
		"snake": Snake{},
	}

	scanner = bufio.NewScanner(os.Stdin)

	for {
		if err := prompt(&input); err != nil {
			fmt.Println(err)
			continue
		}

		command, name, option := input[0], input[1], input[2]

		run(COMMAND_NEWANIMAL == command, func() {
			data, found := factory[option]
			if !found {
				fmt.Println("Wrong animal")
				return
			}

			_, duplicated := animals[name]
			if duplicated {
				fmt.Println("Animal already exists")
				return
			}

			animals[name] = data

			fmt.Println("Created it!")
		})

		run(COMMAND_QUERY == command, func() {
			action, found := animals[name]
			if !found {
				fmt.Println("Animal not found")
				return
			}

			method := reflect.ValueOf(action).MethodByName(strings.Title(option))
			if !method.IsValid() {
				fmt.Println("Wrong action")
				return
			}

			method.Call([]reflect.Value{})
		})

		if !executed {
			fmt.Println("No command found")
		}
	}
}
