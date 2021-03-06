// The but package provides helpers for handling messages while at the bottom
// of the call stack.
package but

import (
	"fmt"
	"os"
	"strings"
)

// IfError prints err to stderr if the error is non-nil. Extra arguments are
// converted to a string which, if present, annotates the error. Returns true
// if the error is non-nil.
func IfError(err error, args ...interface{}) bool {
	if err != nil {
		if len(args) > 0 {
			err = fmt.Errorf(fmt.Sprint(args...)+": %w", err)
		}
		fmt.Fprintln(os.Stderr, err)
		return true
	}
	return false
}

// IfErrorf prints err to stderr if the err is non-nil. Extra arguments are
// formatted as a string, according to the format argument. If present, this
// string annotates the error. Returns true if the error is non-nil.
func IfErrorf(err error, format string, args ...interface{}) bool {
	if err != nil {
		args = append(args, err)
		err = fmt.Errorf(format+": %w", args...)
		fmt.Fprintln(os.Stderr, err)
		return true
	}
	return false
}

// IfFatal prints err to stderr and exits, if the error is non-nil. Extra
// arguments are converted to a string which, if present, annotates the error.
func IfFatal(err error, args ...interface{}) {
	if err != nil {
		IfError(err, args...)
		os.Exit(1)
	}
}

// IfFatalf prints err to stderr and exits, if the err is non-nil. Extra
// arguments are formatted as a string, according to the format argument. If
// present, this string annotates the error.
func IfFatalf(err error, format string, args ...interface{}) {
	if err != nil {
		IfErrorf(err, format, args...)
		os.Exit(1)
	}
}

// Log prints the given arguments to stderr.
func Log(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
}

// Logf formats the arguments according to format, and prints the result to
// stderr.
func Logf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
}

// Fatal prints the given arguments to stderr and exits.
func Fatal(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
	os.Exit(1)
}

// Fatalf formats the arguments according to format, prints the result to
// stderr, and exits.
func Fatalf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}

// Errors groups together multiple errors as a single error.
type Errors struct {
	// Msg is an optional message to be displayed before the list of errors.
	Msg string
	// Errs is the list of errors.
	Errs []error
}

// Error implements the error interface. Errors are displayed one per line,
// each with indentation.
func (err Errors) Error() string {
	s := make([]string, len(err.Errs)+1)
	if err.Msg != "" {
		s[0] = err.Msg
	} else {
		s[0] = "\n\t"
	}
	for i, e := range err.Errs {
		s[i+1] = e.Error()
	}
	return strings.Join(s, "\n\t")
}

// Errors returns the list of errors.
func (err Errors) Errors() []error {
	return err.Errs
}
