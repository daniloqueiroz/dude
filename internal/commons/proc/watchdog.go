package proc

import (
	"errors"
	"github.com/google/logger"
	"os"
	"os/signal"
	"sync"
)

type Watchdog struct {
	processes  []*Process
	updatesChn chan *Process
	barrier    *sync.WaitGroup
	stopChn    chan os.Signal
	started    bool
}

func NewWatchdog() *Watchdog {
	var wg sync.WaitGroup
	return &Watchdog{
		processes:  make([]*Process, 0),
		updatesChn: make(chan *Process),
		barrier:    &wg,
		stopChn:    make(chan os.Signal),
		started:    false,
	}
}

func (dog *Watchdog) fireAndWatch(proc *Process) {
	_ = proc.FireAndWait()
	if dog.started {
		dog.updatesChn <- proc
	}
}

func (dog *Watchdog) supervisor() {
	logger.Info("Watchdog supervisor started")
	for {
		select {
		case proc := <-dog.updatesChn:
			logger.Infof("Starting process %s", proc.cmd)
			go dog.fireAndWatch(proc)
		case sig := <-dog.stopChn:
			logger.Infof("Signal %v received, stopping watchdog", sig)
			dog.shutdown()
			return
		}
	}
}

func (dog *Watchdog) Supervise(proc *Process) error {
	if dog.started {
		return errors.New("Watchdog is already started")
	}

	dog.processes = append(dog.processes, proc)
	return nil
}

func (dog *Watchdog) Start(signals ...os.Signal) error {
	if dog.started {
		return errors.New("Watchdog is already started")
	}

	signal.Notify(dog.stopChn, signals...)
	dog.barrier.Add(1)
	dog.started = true

	go dog.supervisor()
	for _, proc := range dog.processes {
		dog.updatesChn <- proc
	}

	dog.barrier.Wait()
	return nil
}

func (dog *Watchdog) shutdown() {
	dog.started = false
	close(dog.updatesChn)
	for _, proc := range dog.processes {
		if err := proc.goCmd.Process.Kill(); err == nil {
			logger.Infof("Process %s killed", proc.cmd)
		} else {
			logger.Infof("Error killing process %s: %v", proc.cmd, err)
		}
	}
	dog.barrier.Done()
}
