package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() string {
	return a.food
}

func (a *Animal) Move() string {
	return a.locomotion
}

func (a *Animal) Speak() string {
	return a.noise
}

func main() {
	animals := map[string]Animal{
		"cow":   {food: "grass", locomotion: "walk", noise: "moo"},
		"bird":  {food: "worms", locomotion: "fly", noise: "peep"},
		"snake": {food: "mice", locomotion: "slither", noise: "hsss"},
	}

	actions := [3]string{"eat", "move", "speak"}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Insert an animal followed by an action (like cow eat): \n\n")

	for {
		fmt.Print("> ")
		scanner.Scan()

		input := scanner.Text()
		data := strings.Split(input, " ")

		if len(data) != 2 {
			fmt.Print("Paramters are incorrect\n\n")
			continue
		}

		animal := animals[strings.ToLower(data[0])]
		if reflect.ValueOf(animal).IsZero() == true {
			fmt.Print("Incorrect animal. Availables: cow, bird and snake\n\n")
			continue
		}

		action := strings.ToLower(data[1])
		found := false

		for _, v := range actions {
			if action == v {
				found = true
				break
			}
		}

		if found == false {
			fmt.Print("Incorrect action. Availables: eat, move and speak\n\n")
			continue
		}

		output := reflect.ValueOf(&animal).MethodByName(strings.Title(action)).Call([]reflect.Value{})

		fmt.Printf("%s\n\n", output)
	}
}
