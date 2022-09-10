package requests

import (
	"errors"
	"fmt"
	"regexp"
)

const maxPageSize = 100

// Реализуй нас.
var (
	errIsNotRegexp     error = errors.New("exp is not regexp")
	errInvalidPage     error = errors.New("invalid page")
	errInvalidPageSize error = errors.New("invalid page size")
)

// Реализуй мои методы.
type ValidationErrors []error

func (v ValidationErrors) Error() string {
	var errs string

	if len(v) > 0 {
		errs = "validation errors:"
	}

	for _, vErr := range v {
		errs = fmt.Sprintf("%s\n\t%v", errs, vErr)
	}

	return errs
}

func (v ValidationErrors) Is(err error) bool {
	for _, vErr := range v {
		if errors.Is(vErr, err) {
			return true
		}
	}

	return false
}

type SearchRequest struct {
	Exp      string
	Page     int
	PageSize int
}

func (r SearchRequest) Validate() error {
	var errs ValidationErrors

	_, err := regexp.Compile(r.Exp)
	if err != nil {
		errs = append(errs, fmt.Errorf("%w: %v", errIsNotRegexp, err))
	}

	if r.Page <= 0 {
		errs = append(errs, fmt.Errorf("%w: %d", errInvalidPage, r.Page))
	}

	if r.PageSize <= 0 {
		errs = append(errs, fmt.Errorf("%w: %d <= 0", errInvalidPageSize, r.PageSize))
	}

	if r.PageSize > maxPageSize {
		errs = append(errs, fmt.Errorf("%w: %d > %d", errInvalidPageSize, r.PageSize, maxPageSize))
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}
