package errors

import "errors"

func IsTemporary(err error) bool {
	type t interface {
		IsTemporary() bool
	}

	for {
		e, ok := err.(t)
		if ok {
			return e.IsTemporary()
		}

		err = errors.Unwrap(err)
		if err == nil {
			return false
		}
	}
}
