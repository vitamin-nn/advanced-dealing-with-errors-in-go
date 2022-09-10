package errs

type Unwrapper interface {
	Unwrap() error
}

func Unwrap(err error) error {
	for {
		errP, ok := err.(Unwrapper)
		if !ok {
			return err
		}

		err = errP.Unwrap()
	}
}
