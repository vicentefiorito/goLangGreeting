package greetings

import (
	"fmt"
	"errors"

	// import standar library math and the random module
	"math/rand"
)


// Hello returns a greeting for the named person.
func Hello(name string) (string,error) {
	// if no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}
	// If a name was received, return a value that embeds the name
	// in a greeting message
    // message := fmt.Sprintf(randomFormat(), name)
	message := fmt.Sprintf("Hello %v, Welcome to Golang! ", name)
    return message, nil
}

//Hellos returns a map that associates each of the named people
//with a greeting message. --> exportable because of the Uppercase defintion of the function
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)
	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil,err
		}
		// In the map, associate the retrieved message with
		// the name
		messages[name] = message
	}
	// return the message for each name and the error --> in this case,everything works correctly, so you return nil (not an error)
	return messages,nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
