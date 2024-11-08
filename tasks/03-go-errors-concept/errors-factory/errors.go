package errors

// NewError возвращает новое значение-ошибку, текст которой является msg.
// Две ошибки с одинаковым текстом, созданные через NewError, не равны между собой:
//
//  NewError("end of file") != NewError("end of file")
//
func NewError(msg string) error {
	return &NewErr{
		msg: msg,
	}
}

type NewErr struct {
	msg string
}

func (e *NewErr) Error() string {
	return e.msg
}