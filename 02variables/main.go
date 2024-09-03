package main

import "fmt"

const LoginToken string = "sdvsdvsdv"  // public caps on variable first letter
const loginToken2 string = "sdvsdvsdv" // private

func main() {
	var username string = "sathish"
	fmt.Println(username)
	fmt.Printf("variable type is %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable type is %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable type is %T \n", smallVal)

	var smallFloat float32 = 255.325643634734634634634
	fmt.Println(smallFloat)
	fmt.Printf("variable type is %T \n", smallFloat)

	var largeFloat float64 = 255.325643634734634634634
	fmt.Println(largeFloat)
	fmt.Printf("variable type is %T \n", largeFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable type is %T \n", anotherVariable)

	// implicit type. no need to declare type
	var website = "http://qualzure.com"
	fmt.Println(website)

	// no need to use keyword var
	numberOfUser := 30
	fmt.Println(numberOfUser)
}
