package greetings

import (
    "testing"
    "regexp"
)

func TestHelloName(t *testing.T) {
    name := "Edward"
    expected := regexp.MustCompile(`\b`+name+`\b`) 
    msg, err := Hello("Edward")
    if !expected.MatchString(msg) || err != nil {
        t.Fatalf(`Hello("Edward") = %q, %v, expected %#q, nil`, msg, err, expected)
    }
}

func TestHelloEmpty(t *testing.T) {
    msg, err := Hello("")
    if msg != "" || err == nil {
        t.Fatalf(`Hello("") = %q, %v, expected "", error`, msg, err)
    }
}

