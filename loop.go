package main

func main() {
	var sum int = 0
	var result = [8]float32{3.49, 3.29, 3.39, 3.51, 3.72, 3.90, 3.90, 3.86}
	// simple for loop
	for i := 1; i <= 10; i++ {
		sum += i
	}

	println("Summation of number from 1 to 10", sum)

	// simple for loop with continue statement
	for i := 10; i <= 100; i += 10 {
		if i == 20 {
			continue
		}
		println(i)
	}

	// simple for loop with break statement
	for i := 10; i <= 100; i += 10 {
		if i == 20 {
			break
		}
		println(i)
	}

	//Nested for loop
	for i := 1; i <= 10; i++ {

		for j := 1; j <= i; j++ {
			print("*")
		}
		print("\n")
	}

	//while loop using for keyword
	println("while loop")
	number := 1
	for number <= 10 {
		print(number, "\t")
		number++
	}
	// range
	println("\n range function applied on array, slice, map")
	for _, value := range result {
		print(value, "\t")
	}
}
