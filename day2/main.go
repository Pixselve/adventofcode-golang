package main

import (
	"errors"
	"fmt"
)

var puzzleInput = []uint{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 9, 1, 19, 1, 19, 5, 23, 1, 23, 5, 27, 2, 27, 10, 31, 1, 31, 9, 35, 1, 35, 5, 39, 1, 6, 39, 43, 2, 9, 43, 47, 1, 5, 47, 51, 2, 6, 51, 55, 1, 5, 55, 59, 2, 10, 59, 63, 1, 63, 6, 67, 2, 67, 6, 71, 2, 10, 71, 75, 1, 6, 75, 79, 2, 79, 9, 83, 1, 83, 5, 87, 1, 87, 9, 91, 1, 91, 9, 95, 1, 10, 95, 99, 1, 99, 13, 103, 2, 6, 103, 107, 1, 107, 5, 111, 1, 6, 111, 115, 1, 9, 115, 119, 1, 119, 9, 123, 2, 123, 10, 127, 1, 6, 127, 131, 2, 131, 13, 135, 1, 13, 135, 139, 1, 9, 139, 143, 1, 9, 143, 147, 1, 147, 13, 151, 1, 151, 9, 155, 1, 155, 13, 159, 1, 6, 159, 163, 1, 13, 163, 167, 1, 2, 167, 171, 1, 171, 13, 0, 99, 2, 0, 14, 0}

func main() {
	Part1()
	result, err := Part2()
	if err != nil {
		fmt.Println("No solution for part 2", err)
	}

	fmt.Println("What is 100 * noun + verb?", 100*result.noun+result.verb)
}

func Part1() uint {
	resultInput, err := OpcodeManager(0, InputWithNounAndVerb(puzzleInput, NounAndVerb{
		noun: 12,
		verb: 2,
	}))
	if err != nil {
		fmt.Print("Error")
	}
	fmt.Println("What value is left at position 0 after the program halts?", resultInput[0])
	return resultInput[0]
}

type NounAndVerb struct {
	noun uint
	verb uint
}

func Part2() (NounAndVerb, error) {

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			resultInput, err := OpcodeManager(0, InputWithNounAndVerb(puzzleInput, NounAndVerb{
				noun: uint(noun),
				verb: uint(verb),
			}))
			if err != nil {
				return NounAndVerb{}, err
			} else if resultInput[0] == 19690720 {
				return NounAndVerb{
					noun: uint(noun),
					verb: uint(verb),
				}, nil
			}

		}
	}
	return NounAndVerb{}, errors.New("result not found")
}

func InputWithNounAndVerb(input []uint, params NounAndVerb) []uint {
	newInput := make([]uint, len(input))
	copy(newInput, input)
	newInput[1] = params.noun
	newInput[2] = params.verb
	return newInput
}

func OpcodeManager(position uint, integers []uint) ([]uint, error) {
	if integers[position] == 99 {
		return integers, nil
	} else if integers[position] == 1 || integers[position] == 2 {
		a := integers[integers[position+1]]
		b := integers[integers[position+2]]
		c := integers[position+3]

		if integers[position] == 1 {
			integers[c] = a + b
		} else {
			integers[c] = a * b
		}

		newIntegers, err := OpcodeManager(position+4, integers)
		if err != nil {
			return nil, err
		} else {
			return newIntegers, nil
		}
	} else {
		return nil, errors.New(fmt.Sprint("Unexpected opcode ", integers[position]))
	}

}
