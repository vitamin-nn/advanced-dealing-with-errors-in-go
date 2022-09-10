package pipe

import (
	"fmt"
)

type PipelineError struct {
	User        string
	Name        string
	FailedSteps []string
}

func (p *PipelineError) Error() string {
	return fmt.Sprintf("pipeline %q error", p.Name)
}

func (p *PipelineError) Is(target error) bool {
	if errT, ok := target.(*PipelineError); ok {
		return p.Name == errT.Name && p.User == errT.User
	}
	return false
}

// Добавь метод Is для типа *PipelineError.
