package pipe

import "fmt"

type UserError struct {
	Operation string
	User      string
}

func (u *UserError) Error() string {
	return fmt.Sprintf("user %s cannot do op %s", u.User, u.Operation)
}

type PipelineError struct {
	User        string
	Name        string
	FailedSteps []string
}

func (p *PipelineError) Error() string {
	return fmt.Sprintf("pipeline %q error", p.Name)
}

func (p *PipelineError) As(target interface{}) bool {
	if errT, ok := target.(**UserError); ok {
		*errT = &UserError{
			Operation: p.Name,
			User: p.User,
		}

		return true
	}

	return false
}
// Добавь метод As для типа *PipelineError.
