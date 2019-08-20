package main

import (
	"fmt"
)

func recoverName() {
	if r := recover(); r != nil {
		fmt.Println("Recover from: ", r)
	}
}

func fullName(firstName *string, lastName *string) {
	defer fmt.Println("defered from fullName().")
	defer recoverName()

	if firstName == nil {
		panic("Runtime error: firstname cannot be nil. ")
	}

	if lastName == nil {
		panic("Runtime error: lastname cannot be nil. ")
	}

	fmt.Printf("Hello, %s %s! \n", *firstName, *lastName)
	fmt.Println("Return normally from fullName().")
}

func main() {
	defer fmt.Println("defered from main().")

	firstName := "Elon"
	//lastName := "Mask"
	fullName(&firstName, nil)

	fmt.Println("Return normally from main().")

}
