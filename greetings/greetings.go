package greetings

import (
    "fmt"
    "errors"
    "math/rand"
    "time"
)

// Hello returns a greeting for the named person
func Hello(name string) (string,error) {
    
    // If no name was given, return an error
    if name == "" {
        return "", errors.New("missing name")
    }

    message := fmt.Sprintf(randomFormat(), name)
    // message := fmt.Sprintf("Hi %v welcome to my first go module!", name)
    return message, nil
}

func Hellos(names []string) (map[string]string, error) {
    messages := make(map[string]string)

    for _, name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }
        messages[name] = message
    }
    return messages, nil
}

// Set seed so we have random choice of greeting each time randomFormat() is ran
// Default behavior is deterministic
func init() {
    rand.Seed(time.Now().UnixNano())
}

// return a random meeting
func randomFormat() string {
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    return formats[rand.Intn(len(formats))]
}
