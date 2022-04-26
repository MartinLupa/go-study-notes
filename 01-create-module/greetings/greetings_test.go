package greetings

import (
	"regexp"
	"strings"
	"testing"
)

/*
The go test command executes test functions (whose names begin with Test) in test files (whose names end with _test.go).
You can add the -v flag to get verbose output that lists all of the tests and their results.
*/

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string or with blank spaces,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if len(strings.TrimSpace(msg)) != 0 || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
