package supervisor

import (
	"context"
	"github.com/daniloqueiroz/dude/app/system/proc"
	"os"
	"os/signal"
	"sync"
	"time"
)

type Supervisor struct {
	children []*Child
	ctl      *control
}

func NewSupervisor() *Supervisor {
	ctx, cancel := context.WithCancel(context.Background())
	control := &control{ctx: ctx, cancel: cancel, barrier: &sync.WaitGroup{}}
	return &Supervisor{
		children: make([]*Child, 0),
		ctl:      control,
	}
}

func (s *Supervisor) Stop() {
	s.ctl.Stop()
}

func (s *Supervisor) Start() {
	for _, child := range s.children {
		go supervise(s.ctl, child)
	}
}

func (s *Supervisor) Wait() {
	time.Sleep(1 * time.Second)
	s.ctl.Wait()
}

func (s *Supervisor) AddTask(name string, task Task) *Child {
	child := &Child{
		IsRunning: false,
		Name:      name,
		task:      task,
	}
	s.children = append(s.children, child)
	return child
}

func (s *Supervisor) AddProc(proc *proc.Process) *Child {
	return s.AddTask(proc.Cmd, func(ctx context.Context) error {
		return proc.FireAndWait()
	})
}

func (s *Supervisor) AddSigHandler(handler func(), signals ...os.Signal) {
	sigChn := make(chan os.Signal, 1)
	signal.Notify(sigChn, signals...)
	go sigHandler(s.ctl, sigChn, handler)
}
