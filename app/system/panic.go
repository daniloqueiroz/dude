package system

import (
	"errors"
	"fmt"
	"github.com/google/logger"
)

func OnPanic(name string, errChn chan error) {
	r := recover()
	if r != nil {
		logger.Errorf("Panic on %s: %+v", name, r)
		errChn <- errors.New(fmt.Sprintf("Panic: %+v", r))
	}
}
