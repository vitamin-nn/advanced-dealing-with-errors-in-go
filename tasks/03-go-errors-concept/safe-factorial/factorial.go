package factorial

import (
	"errors"
	"fmt"
)

const maxDepth = 256

// Реализуй нас.
var (
	ErrNegativeN = errors.New("N is negative")
	ErrTooDeep = fmt.Errorf("maximum deep is %d", maxDepth)
)

// Calculate рекурсивно считает факториал входного числа n.
// Если число меньше нуля, то возвращается ошибка ErrNegativeN.
// Если для вычисления факториала потребуется больше maxDepth фреймов, то Calculate вернёт ErrTooDeep.
func Calculate(n int) (int, error) {
	if n > maxDepth {
		return 0, ErrTooDeep
	}

	if n == 0 {
		return 1, nil
	}

	if n < 0 {
		return 0, ErrNegativeN
	}

	next, err := Calculate(n - 1)
	if err != nil {
		return 0, err
	}

	return next * n, nil
}
