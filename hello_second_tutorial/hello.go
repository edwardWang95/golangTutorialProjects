package main

import (
    "fmt"
    "log"

    "example.com/greetings"
)

func main() {
    // Set properties of the predefined Logger, include the log entry prefix and flag to disable printing the time, source file and line #
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // names
    names := []string{"Olivia", "Rebecca", "Larry"}

    // Request a greeting message
    //message, err := greetings.Hello("Edward")
    messages, err := greetings.Hellos(names)
    // print an error if it is returned and exit the program
    if err != nil {
        log.Fatal(err)
    }
    
    //Get a greeting message and print it
    fmt.Println(messages)
}
