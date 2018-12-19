package main

import "fmt"

func main() {
	fmt.Println("hello")

	var age int = 40

	var favNum float64 = 3.384

	//var randNum := 1

	//randNum = "snth" //compile error

	fmt.Println(age, " ", favNum)
	fmt.Println(age, favNum)
	fmt.Println("san\n ")

	const pi float64 = 3.14

	// loop

	i := 1

	for i <= 10 {
		fmt.Println("i", i, "\n")

		i++

	}

	for j := 0; j < 5; j++ {
		fmt.Println("j", j, "\n")

	}

	yourAge := 18

	if yourAge >= 16 {
		fmt.Println("You can Drive\n ")

	} else if yourAge >= 18 {
		fmt.Println("you can vote\n ")

	}

	switch yourAge {
	case 16:
		fmt.Println("Go Drive\n ")
	case 18:
		fmt.Println("Go Vote\n ")

	}

	var faNums2 [5]float64 // an array of max size 5 of type float
	faNums2[0] = 123
	faNums2[1] = 4123
	faNums2[2] = 1223
	faNums2[3] = 12223
	faNums2[4] = 12.3

	fmt.Println("faNums2[3]", faNums2[3], "\n")

	// shorthand
	favNums3 := [5]float64{1, 2, 3, 4, 5}

	// iterate array
	for i, value := range favNums3 { // can use _ instead of i if not needed
		fmt.Println(value, i)

	}

} // end of main
