package main

import (
	"fmt"
)

func GetUserInput(var_type string) float64 {
	fmt.Printf("Please input the %s: ", var_type)
	var input float64
	fmt.Scan(&input)
	return input
}

// s = ½ a t2 + vot + so
func GenDisplaceFn(a float64, vo float64, so float64) func(float64) float64 {
	fn := func(time float64) float64 {
		return 0.5*a*(time*time) + (vo * time) + so
	}
	return fn
}

func main() {
	user_a := GetUserInput("acceleration")
	user_vo := GetUserInput("initial velocity")
	user_so := GetUserInput("initial displacement")

	fn := GenDisplaceFn(user_a, user_vo, user_so)

	fmt.Println("Below are the tests returning total displacement of the function returned by GenDisplaceFn() for time = 3 and time = 5:")
	fmt.Println(fn(3))
	fmt.Println(fn(5))
}
