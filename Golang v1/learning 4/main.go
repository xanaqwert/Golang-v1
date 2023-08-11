package main

import "fmt"

func main() {
	age := 18
	name := "cesla"

	//*! PRINT function
	//*Print
	fmt.Print("Hello, ")
	fmt.Print("World! \n")
	fmt.Print("New line \n")

	//*Println
	fmt.Println("Hello Cesla")
	fmt.Println("Goodbye Cesla")
	fmt.Println("my age is", age, "and my name is", name)

	fmt.Printf("\n")

	//* Printf (formatted strings) %_ = format specifier
	fmt.Printf("My age is %v and my name is %v \n", age, name) //*! %v is to call the variabel, the default is the first one that appear when the string end.
	fmt.Printf("my age is %q and my name is %q \n", age, name) //*! $q is to call the variabel, but it's only for quote/string
	fmt.Printf("age is of type %T \n", age)                    //*! %T to identify the type of the variabel it self
	fmt.Printf("you scored %f points! \n", 255.55)             //*! %f is for the decimal
	fmt.Printf("you scored %0.1f points! \n", 255.55)          //*! is for the deciman, but u can rounded it

	//* Printf (save formatted settings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name) //*! is to saved the var, that u wanted to, just call the variabel it's self
	fmt.Println("the saved string is:", str)

}
