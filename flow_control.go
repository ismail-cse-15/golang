package main

func main() {
	var number1 = 10
	if number1%2 == 0 {

		println(number1, "is even number")
	} else {
		println(number1, "is odd number")

	}

	if number1 < 0 {
		println("It is less than zero")
	} else if number1 > 0 {
		println("It is greater than zero")
	} else {
		println("It is zero")
	}

}
