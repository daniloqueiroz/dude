package supervisor

import (
	"context"
	"github.com/daniloqueiroz/dude/app/system"
)

type Task func(ctx context.Context) error

type Child struct {
	IsRunning bool
	Name      string
	task      Task
}

func (c *Child) run(ctx context.Context, errChn chan error) {
	defer system.OnPanic(c.Name, errChn)
	err := c.task(ctx)
	errChn <- err
}
