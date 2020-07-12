package main

import "testing"

func TestFuelPerModulePart1(t *testing.T) {
	tests := []struct {
		mass int
		fuel int
	}{
		{mass: 12, fuel: 2},
		{mass: 14, fuel: 2},
		{mass: 1969, fuel: 654},
		{mass: 100756, fuel: 33583},
	}
	for _, test := range tests {
		t.Run(string(test.mass), func(t *testing.T) {
			if FuelPerModulePart1(test.mass) != test.fuel {
				t.Error("A module with a mass of ", test.mass, " requires ", test.fuel, " fuel")
			}
		})
	}
}

func TestFuelPerModulePart2(t *testing.T) {
	tests := []struct {
		name string
		mass int
		fuel int
	}{
		{name: "A module of mass 14", mass: 14, fuel: 2},
		{name: "A module of mass 1969", mass: 1969, fuel: 966},
		{name: "A module of mass 100756", mass: 100756, fuel: 50346},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if FuelPerModulePart2(test.mass) != test.fuel {
				t.Error(test.name, "requires ", test.fuel, " fuel")
			}
		})
	}
}
