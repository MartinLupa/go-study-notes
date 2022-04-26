/*
STEP 1: go mod init example.com/greetings
The go mod init command creates a go.mod file to track your code's dependencies.
At the begining, the file includes only the name of your module and the Go version your code supports.
But as you add dependencies, the go.mod file will list the versions your code depends on.
This keeps builds reproducible and gives you direct control over which module versions to use.
*/

package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

/*
Hello returns a greeting for the named person.
This function takes a name parameter whose type is string. The function also returns a string.
In Go, a function whose name starts with a capital letter can be called by a function not in the same package.
This is known in Go as an exported name.
*/

// ===================================================================================
// Common error handling in Go: Return an error as a value so the caller can check for it.
func Hello(name string) (string, error) {
	//If no name was given, return an error with a message.
	if name == "" || len(strings.TrimSpace(name)) == 0 {
		return "", errors.New("please write your name")
	}

	//If a name was received, return a value that embeds the name.
	//Sprintf formats according to a format specifier and returns the resulting string.
	//Create a message using a random format.
	message := fmt.Sprintf(randomFormat(), name)
	//nil is a zero value, meaning no errors.
	return message, nil
}

// ===================================================================================
//Hellos returns a map that associates each of the named people with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	//A map to associate names with messages.
	messages := make(map[string]string)
	//Loop through the received slice of names, calling the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		//In the map, associate the retrieved message with the name.
		messages[name] = message
	}

	return messages, nil
}

// ===================================================================================
//Add an init function to seed the rand package with the current time.
//Go executes init functions automatically at program startup, after global variables have been initialized. For more about init functions, see Effective Go.
// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// ===================================================================================
//randomFormat returns one of a set of greeting messages. The returned message is selected at random.
func randomFormat() string {
	//A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	//Return a randomly selected message from format by specifying a random index for the formats slice.
	return formats[rand.Intn(len(formats))]
}
