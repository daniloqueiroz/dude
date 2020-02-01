package proc

import (
	"bytes"
	"errors"
	"github.com/google/logger"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Daemon interface {
	Start(signals ...os.Signal) error
}

type simpled struct {
	sigChn  chan os.Signal
	started bool
	barrier *sync.WaitGroup
	fn func()
}

func NewDaemon(fn func()) *simpled {
	var wg sync.WaitGroup
	return &simpled{
		sigChn:  make(chan os.Signal),
		started: false,
		barrier: &wg,
		fn: fn,
	}
}

func (d *simpled) Start(signals ...os.Signal) error {
	if d.started {
		return errors.New("Daemon is already started")
	}

	signal.Notify(d.sigChn, signals...)
	d.barrier.Add(1)
	d.started = true

	go d.fn()
	go d.sigHandler()
	d.barrier.Wait()
	return nil
}

func (d *simpled) sigHandler() {
	logger.Info("Daemon waiting for signal")
	sig := <- d.sigChn
	logger.Infof("Signal %s received, shutting down daemon", sig)
	d.barrier.Done()
}

func LaunchDaemon(wd *Watchdog, name string) {
	dudePath, err := getExecutablePath()
	if err != nil {
		logger.Fatal("Unable to locate dude binary", err)
	}
	logger.Infof("Launching dude daemon %s", name)
	cmd := NewProcess(dudePath, "daemon", name)
	wd.Supervise(cmd)
}

func getExecutablePath() (string, error) {
	name := "/proc/self/exe"
	for len := 128; ; len *= 2 {
		b := make([]byte, len)
		n, e := syscall.Readlink(name, b)
		if e != nil {
			return "", &os.PathError{"readlink", name, e}
		}
		if n < len {
			if z := bytes.IndexByte(b[:n], 0); z >= 0 {
				n = z
			}
			return string(b[:n]), nil
		}
	}
}
