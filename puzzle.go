package main

import (
	"fmt"

	"github.com/thatisuday/commando"
	"watchblade.com/lab/puzzle/solve"
)

func main() {
	fmt.Println("Starting the puzzle")
	commando.SetExecutableName("puzzle").SetVersion("0.0.1").SetDescription("Solve a puzzle")

	// configure the root command
	commando.Register(nil).AddArgument("door", "find door open and close state", "").
		AddFlag("number", "number of doors", commando.Int, 10).
		AddFlag("state", "the door is open or close initially. open=0, close=1", commando.Int, 1).
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			// read values from the command line
			number_of_doors := flags["number"].Value.(int)
			door_state := flags["state"].Value.(int)

			// check of teh door state is more than 1 then set it to 1			
			if door_state > 1 { door_state = 1 }

			// solve the puzzle
			solve.DoorOpen(number_of_doors, door_state)
		})

	commando.Parse(nil)

}
