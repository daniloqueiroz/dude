package supervisor

import (
	"context"
	"fmt"
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/google/logger"
	"os"
	"sync"
	"time"
)

type control struct {
	ctx     context.Context
	cancel  context.CancelFunc
	barrier *sync.WaitGroup
}

func (c *control) Up() {
	c.barrier.Add(1)
}

func (c *control) Down() {
	c.barrier.Done()
}

func (c *control) Wait() {
	c.barrier.Wait()
	c.cancel()
}

func (c *control) Stop() {
	c.cancel()
}

func (c *control) DoneSig() <-chan struct{} {
	return c.ctx.Done()
}

func sigHandler(ctl *control, sigChn chan os.Signal, callback func()) {
	ctl.Up()
	select {
	case sig := <-sigChn:
		logger.Infof("Signal %v received", sig)
		callback()
	case <-ctl.DoneSig():
		break
	}
	ctl.Down()
}

func supervise(ctl *control, task *Child) {
	chn := make(chan error, 1)
	defer close(chn)
	defer system.OnPanic(fmt.Sprintf("task_supervisor-%v", task.Name), chn)

	ctl.Up()
	for {
		go task.run(ctl.ctx, chn)
		task.IsRunning = true
		select {
		case err := <-chn:
			task.IsRunning = false
			logger.Errorf("Task %v failed and will be restarted soon: %v", task.Name, err)
			time.Sleep(5 * time.Second)  // TODO backoff and max_retries
		case <-ctl.DoneSig():
			logger.Infof("Supervisor is done: %v", task.Name)
			ctl.Down()
			return
		}
	}
}
