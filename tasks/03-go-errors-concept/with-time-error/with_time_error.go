package errs

import (
	"fmt"
	"time"
)

type WithTimeError struct { // Реализуй меня.
	t time.Time
	err error
}

func (e WithTimeError) Error() string {
	return fmt.Sprintf("%s, occurred at: %s", e.err.Error(), e.t.Format("2006-06-07T20:48:39.478061+03:00"))
}

func (e *WithTimeError) Time() time.Time {
	return e.t
}

func (e *WithTimeError) Unwrap() error {
	return e.err
}

func NewWithTimeError(err error) error {
	return newWithTimeError(err, time.Now)
}

func newWithTimeError(err error, timeFunc func() time.Time) error {
	return &WithTimeError{
		t: timeFunc(),
		err: err,
	}
}
