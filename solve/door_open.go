package solve

import (
	"fmt"
	"sort"
)

func DoorOpen(number_of_doors int, door_state int) {
	fmt.Printf("Number of doors given %d\n", number_of_doors)
	fmt.Printf("Door state given %d\n", door_state)

	// build initial door state and print
	doors := get_initial_door_state(number_of_doors, door_state)

	doors_keys := get_all_keys(doors)
	sort.Ints(doors_keys)

	// call function to solve the puzzle
	doors_changed := solve_puzzle(doors, number_of_doors)

	doors_keys_for_changed := get_all_keys(doors_changed)
	sort.Ints(doors_keys_for_changed)

	fmt.Println("Final door state")
	number_of_open_doors := 0
	number_of_closed_doors := 0
	for _, v := range doors_keys_for_changed {
		if doors_changed[v] == 0 {
			number_of_open_doors++
		} else {
			number_of_closed_doors++
		}		
		fmt.Printf("Door %d is %d\n", v, doors_changed[v])
	}

	fmt.Printf("Number of open doors %d and Number of closed doors %d \n", number_of_open_doors, number_of_closed_doors)
}

func solve_puzzle(doors map[int]int, number_of_doors int) map[int]int {
	doors_changed := get_doors_copy(doors)
	// solve the puzzle
	for i := 1; i <= number_of_doors; i++ {
		next_door_number := i
		for next_door_number <= number_of_doors {
			// change door state for the door to change
			if doors_changed[next_door_number] == 0 {
				doors_changed[next_door_number] = 1
			} else {
				doors_changed[next_door_number] = 0
			}

			next_door_number = next_door_number + i
		}
	}

	return doors_changed
}

func get_initial_door_state(number_of_doors int, door_state int) map[int]int {
	doors := make(map[int]int)
	for i := 1; i <= number_of_doors; i++ {
		doors[i] = door_state
	}
	return doors
}

func get_all_keys(doors map[int]int) []int {
	keys := make([]int, 0)
	for k := range doors {
		keys = append(keys, k)
	}
	return keys
}

func get_doors_copy(doors map[int]int) map[int]int {
	doors_copy := make(map[int]int)
	for k := range doors {
		doors_copy[k] = doors[k]
	}
	return doors_copy
}
