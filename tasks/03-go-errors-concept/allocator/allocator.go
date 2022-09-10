package allocator

const (
	Admin          = 777
	MinMemoryBlock = 1024
)

type NotPermittedError struct{}

type ArgOutOfDomainError struct{}

func (NotPermittedError) Error() string {
	return "operation not permitted"
}

func (ArgOutOfDomainError) Error() string {
	return "numerical argument out of domain of func"
}

func Allocate(userID, size int) ([]byte, error) {
	if userID != Admin {
		return nil, new(NotPermittedError)
	}

	if size < MinMemoryBlock {
		return nil, new(ArgOutOfDomainError)
	}

	return make([]byte, size), nil
}
