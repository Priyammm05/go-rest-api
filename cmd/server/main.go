package main

import "fmt"

func Run() error {
	fmt.Println("Starting point")
	return nil
}

func main() {
	fmt.Println("GO REST API")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}