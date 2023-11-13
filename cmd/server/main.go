package main

import "fmt"

// * Responsible for the instantiation and startup of our Go app
func Run() error {
	fmt.Println("Starting up our application...")
	return nil
}

func main() {
	fmt.Println("Go REST API Course")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
