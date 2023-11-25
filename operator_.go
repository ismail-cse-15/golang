package main

func main() {

	var (
		number1 = 10
		number2 = 6
		number3 = 10
	)

	// basic arithmetic operator (+,-,*,/,%,++,--)
	println("Summation of two number:", number1+number2)
	println("Subctration of two number:", number1-number2)
	println("Multiplication of two number:", number1*number2)
	println("Division of two number:", number1/number2)
	println("Reminder after Division:", number1%number2)
	number1++
	println("Number one after increment:", number1)
	number1--
	println("Number one after decrement:", number1)

	// basic comparison operator
	println(number1 == number2, number1 == number3)
	println(number1 > number2, number2 < number1, number1 > number3)
	println(number1 != number2, number1 != number3)
	println(number1 >= number3, number1 <= number3, number1 <= number2, number2 >= number3)

	//

}
