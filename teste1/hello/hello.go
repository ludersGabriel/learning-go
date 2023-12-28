package main

import (
	"example/greetings"
	"fmt"
	"log"
)

func main() {
	names := []string{"Gladys", "Samantha", "Darrin"}

	message, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
