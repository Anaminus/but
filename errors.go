// The but package provides helpers for handling messages while at the bottom
// of the call stack.
package but

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strings"
)

// IfError prints err to stderr if the error is non-nil. Extra arguments are
// converted to a string which, if present, annotates the error. Returns true
// if the error is non-nil.
func IfError(err error, args ...interface{}) bool {
	if err != nil {
		if len(args) > 0 {
			err = errors.Wrap(err, fmt.Sprint(args...))
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
		if len(args) > 0 {
			if format, ok := args[0].(string); ok {
				err = errors.Wrapf(err, format, args[1:])
			}
		}
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
