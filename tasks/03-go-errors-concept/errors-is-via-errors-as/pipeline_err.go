package pipe

import (
	"errors"
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

func IsPipelineError(err error, user, pipelineName string) bool {
	var errT *PipelineError
	if errors.As(err, &errT) {
		return errT.User == user && errT.Name == pipelineName
	}

	return false
}
