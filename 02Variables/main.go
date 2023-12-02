package main

import "fmt"

const MyConst string = "dalfkdsf"

func main() {
	fmt.Println("Variables Files")

	var username string = "Ayush"
	fmt.Println(username)
	fmt.Printf("Type of this variable is %T \n", username)

	var isLoggedin bool = false
	fmt.Println(isLoggedin)
	fmt.Printf("Type of this variable is %T \n", isLoggedin)

	var number int = 300
	fmt.Println(number)
	fmt.Printf("Type of this variable is %T \n", isLoggedin)

	var floatingNumber float64 = 45.99
	fmt.Println(floatingNumber)
	fmt.Printf("Type of this variable is %T \n", floatingNumber)

	//Implcit Variable
	var webTest = "www.google.com"
	fmt.Println(webTest)

	//without var
	testNumber := 5000 //Cannot be declared outside the method
	fmt.Println(testNumber)

	//Printing const
	fmt.Println(MyConst)
}
