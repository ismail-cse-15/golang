package main

func main() {
	day := 10
	day_no := 4
	// ssingle switch case
	switch day {

	case 1:
		println("Saturday")
	case 2:
		println("Sunday")

	case 3:
		println("Monday")
	case 4:
		println("Tuesday")

	case 5:
		println("wednesday")

	case 6:
		println("Thursday")
	case 7:
		println("Friday")
	default:
		println("Not a weekday")
	}
	// Multiple switch case
	switch day_no {

	case 1, 3, 5, 7:
		println("Odd weekday")
	case 2, 4, 6:
		println("Even weekday")

	default:
		println("Not a weekday")
	}

}
