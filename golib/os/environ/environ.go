package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	envs := os.Environ()

	for _, env := range envs {
		fmt.Println(env)
	}
	fmt.Println()

	result, err := os.Executable()
	if err != nil {
		log.Fatal("Executable:", err)
	}
	fmt.Println(result)

	fmt.Println(os.ExpandEnv("SHELL is $SHELL\n"))

	fmt.Printf("LOGNAME is %s.\n", os.Getenv("LOGNAME"))
}
