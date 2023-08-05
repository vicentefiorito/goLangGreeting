package main

import (
	// standard library imports
    "fmt"
	"log"
	// import user input
	"bufio"
	// imports operating system
	"os"
	// Imports the greetings module
    "example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

	// creates a scanner to read input
	scanner := bufio.NewScanner(os.Stdin)
	// ask the user to input their name for the greeting
	fmt.Printf("What is your name?: ")
	scanner.Scan()
	username := scanner.Text()

	// A slice of names
	// names := []string{
	// 	"Gladys",
	// 	"Samantha",
	// 	"Darrin",
	// }

 	// Request a greeting message.
    // message, err := greetings.Hello("Mike")

	// Request a greeting message for the NAMES!
	messages,err := greetings.Hello(username)
    // If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }

    // If no error was returned, print the returned message
    // to the console.

	// this is for a single greeting of a name
    // fmt.Println(message)

	// This is for multiple greetings of a slice of names
	fmt.Println(messages)

}