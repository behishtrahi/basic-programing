package main

import "fmt"

func main() {
	// store height
	var height int = 9
	height = 9
	// store ridus
	var ridus int = 4
	ridus = 4
	// store pi
	var phi float64 = 3.14
	phi = 3.14
	// write the formula
	volume := height * ridus * ridus * int(phi)

	fmt.Println("volume is :", volume)
}
