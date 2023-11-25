package main

func main() {
	// Basic output function print, println
	println("1 + 1 =", 1+1)

	print("hello world!\n")

	print("Welcome to learning GoLang")

	// basic variable declaration
	// maltiple variable declaration
	var name, age = "ismail", 27
	println(name, " is ", age, "years old")
	var a, b, c = 10, 15, 30
	var d, e int
	var student string
	var area float64
	println("a :", a, "b :", b, " c:", c)
	println("d :", d, "e :", e)
	println("Student one  name: ", student)
	//printf(" Sdudent one name: %v", student)
	d, e = 40, 50
	student = "Abir"
	student2 := "Pias"
	println("d :", d, "e :", e)
	println("Student one  name: ", student, "Student two  name: ", student2)
	const g float32 = 9.81
	const (
		PI             = 3.1416
		COUNTRY string = "Bangladesh"
		FLAG           = true
	)

	area = PI * (3 * 3)

	println("Area of corcle:", area)

	/// basic array declaration
	var number = [10]int{10, 20, 30}
	println(number[3])

}
