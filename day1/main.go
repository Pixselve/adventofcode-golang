package main

import "fmt"

func main() {
	inputs := []int{
		142156,
		108763,
		77236,
		78186,
		110145,
		126414,
		115436,
		133275,
		132634,
		82606,
		118669,
		90307,
		134124,
		102597,
		128607,
		109214,
		50160,
		72539,
		99033,
		145334,
		135409,
		97525,
		109865,
		142319,
		79027,
		96924,
		72530,
		85993,
		109594,
		115991,
		107998,
		112934,
		85198,
		112744,
		129637,
		95515,
		90804,
		107052,
		89707,
		93658,
		60115,
		118752,
		94315,
		59645,
		115668,
		139320,
		70734,
		56771,
		74741,
		69284,
		92228,
		145376,
		103317,
		55143,
		58370,
		54873,
		52424,
		95392,
		67892,
		90858,
		74693,
		77363,
		51496,
		79375,
		71206,
		103492,
		94065,
		72084,
		144311,
		67381,
		129958,
		86741,
		148906,
		123383,
		147575,
		136327,
		118108,
		136529,
		66356,
		70746,
		147569,
		107267,
		122434,
		69688,
		122127,
		94072,
		110203,
		50546,
		57836,
		139334,
		113240,
		96729,
		68516,
		74635,
		126951,
		138948,
		88312,
		101477,
		129730,
		93816,
	}

	sumPart1 := 0
	sumPart2 := 0

	for _, s := range inputs {
		sumPart1 += FuelPerModulePart1(s)
		sumPart2 += FuelPerModulePart2(s)
	}
	fmt.Println("The sum of the fuel requirements for part 1 is :", sumPart1)
	fmt.Println("The sum of the fuel requirements for part 2 is :", sumPart2)

}

func FuelPerModulePart1(mass int) int {
	fuel := mass/3 - 2
	if fuel >= 0 {
		return fuel
	} else {
		return 0
	}
}

func FuelPerModulePart2(mass int) int {
	switch mass {
	case 0:
		return 0
	default:
		moduleMass := FuelPerModulePart1(mass)
		return moduleMass + FuelPerModulePart2(moduleMass)
	}
}
