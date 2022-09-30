package errors

import "fmt"

// Wrapf работает аналогично fmt.Errorf, только поддерживает nil-ошибки.
func Wrapf(err error, f string, v ...any) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf(f+": %w", append(v, err)...)
}
