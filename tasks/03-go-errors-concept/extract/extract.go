package errs

type listUnwarpper interface {
	Unwrap() []error
}

type unwrapper interface {
	Unwrap() error
}

// Extract достаёт из цепочки err набор sentinel-ошибок,
// игнорируя "оборачивающие" их ошибки.
func Extract(err error) []error {
	var result []error

	if unwrapedErr, ok := err.(listUnwarpper); ok {
		for _, e := range unwrapedErr.Unwrap() {
			result = append(result, Extract(e)...)
		}
	} else if unwrapedErr, ok := err.(unwrapper); ok {
		result = append(result, Extract(unwrapedErr.Unwrap())...)
	} else {
		result = []error{err}
	}

	return result
}
