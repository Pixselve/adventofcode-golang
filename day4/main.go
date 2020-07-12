package main

import "fmt"

func main() {
	fmt.Println(NumberOfPossiblePasswordFirstPart(256310, 732736))
	fmt.Println(NumberOfPossiblePasswordSecondPart(256310, 732736))
}

func AscendingOrder(input uint) bool {
	intSlice := IntToSlice(input, []uint{})
	for index, currentInt := range intSlice {
		if index > 0 && intSlice[index-1] > currentInt {
			return false
		}
	}
	return true
}

func AtLeastARepeat(input uint) bool {
	intSlice := IntToSlice(input, []uint{})
	for index, currentInt := range intSlice {
		if index > 0 && intSlice[index-1] == currentInt {
			return true
		}
	}
	return false
}

func OnlyRepeatTwoTimes(input uint) bool {
	intSlice := IntToSlice(input, []uint{})
	intMap := make(map[uint]uint)
	for _, u := range intSlice {
		intMap[u]++
	}
	for _, u2 := range intMap {
		if u2 == 2 {
			return true
		}
	}
	return false
}

func IntToSlice(n uint, sequence []uint) []uint {
	if n != 0 {
		i := n % 10
		sequence = append([]uint{i}, sequence...)
		return IntToSlice(n/10, sequence)
	}
	return sequence
}

func NumberOfPossiblePasswordFirstPart(a, b uint) uint {
	combinations := uint(0)
	for i := a; i <= b; i++ {
		if AscendingOrder(i) && AtLeastARepeat(i) {
			combinations++
		}
	}

	return combinations
}

func NumberOfPossiblePasswordSecondPart(a, b uint) uint {
	combinations := uint(0)
	for i := a; i <= b; i++ {
		if AscendingOrder(i) && OnlyRepeatTwoTimes(i) {
			combinations++
		}
	}

	return combinations
}
