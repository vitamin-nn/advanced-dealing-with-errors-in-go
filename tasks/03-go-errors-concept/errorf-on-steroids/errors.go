package errs

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`%([a-zA-Z])`)

func Errorf(format string, args ...interface{}) error {
	specs := re.FindAllStringSubmatch(format, -1)

	if len(specs) != len(args) {
		return nil
	}

	var errList []error
	for i, s := range specs {
		if s[1] == "w" {
			if err, ok := args[i].(error); ok {
				errList = append(errList, err)
			}
		}
	}

	if len(errList) == 0 {
		return nil
	}

	formatNew := strings.ReplaceAll(format, "%w", "%v")

	return &steroidErr{
		errList: errList,
		str: fmt.Sprintf(formatNew, args...),
	}
}

type steroidErr struct {
	errList []error
	str string
}

func (e steroidErr) Error() string {
	return e.str
}

func (e steroidErr) Is(err error) bool {
	for _, vErr := range e.errList {
		if errors.Is(vErr, err) {
			return true
		}
	}

	return false
}

func (e steroidErr) As(target interface{}) bool {
	for _, vErr := range e.errList {
		if errors.As(vErr, target) {
			return true
		}
	}

	return false
}