package main

import "fmt"

func main() {

	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Println("Nice!")
	case grade == "B", grade == "C":
		fmt.Println("Good!")
	case grade == "D":
		fmt.Println("OK")
	case grade == "F":
		fmt.Println("Not good")
	default:
		fmt.Println("Bad")
	}
	fmt.Println("Your grade is", grade)
}
