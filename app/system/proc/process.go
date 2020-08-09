package proc

import (
	"errors"
	"fmt"
	"github.com/google/logger"
	"github.com/rkoesters/xdg/basedir"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strconv"
	"syscall"
)

type State string

const (
	NotStarted State = "not_started"
	Running    State = "running"
	Success    State = "success"
	Failed     State = "failed"
)

type Process struct {
	Cmd   string
	args  []string
	goCmd *exec.Cmd
	state State
}

func NewProcess(cmd string, args ...string) *Process {
	return &Process{
		Cmd:   cmd,
		args:  args,
		goCmd: nil,
		state: NotStarted,
	}
}

func (p *Process) State() State {
	return p.state
}

func (p *Process) updateState(newState State) {
	p.state = newState
}

func (p *Process) Stop() {
	if p.State() == Running {
		if err := p.goCmd.Process.Kill(); err != nil {
			logger.Errorf("Error killing process %: %v", p.Cmd, err)
		}
	}
}

func (p *Process) FireAndForget() error {
	if p.State() != NotStarted {
		return errors.New(fmt.Sprintf("Process %s is already started", p.Cmd))
	}
	p.goCmd = exec.Command(p.Cmd, p.args...)
	p.updateState(Running)
	return p.goCmd.Start()
}

func (p *Process) FireAndWait() error {
	out, err := p.FireAndWaitForOutput()
	logger.Infof("%s output: %s", p.Cmd, out)
	return err
}

func (p *Process) FireAndWaitForOutput() (string, error) {
	if p.State() == Running {
		logger.Errorf("Process %s is already running", p.Cmd)
		return "", errors.New(fmt.Sprintf("process %s is already running", p.Cmd))
	} else {
		p.goCmd = exec.Command(p.Cmd, p.args...)
		p.updateState(Running)
	}

	output, err := p.goCmd.CombinedOutput()
	if err != nil {
		p.updateState(Failed)
	} else {
		p.updateState(Success)
	}

	p.goCmd = nil
	return string(output), err
}

func CreatePidFile(name string) {
	pidFile := path.Join(basedir.RuntimeDir, fmt.Sprintf("%s.pid", name))
	logger.Infof("%s pid file %s", name, pidFile)
	if piddata, err := ioutil.ReadFile(pidFile); err == nil {
		if pid, err := strconv.Atoi(string(piddata)); err == nil {
			if process, err := os.FindProcess(pid); err == nil {
				if err := process.Signal(syscall.Signal(0)); err == nil {
					logger.Fatalf("pid already running: %d", pid)
				}
			}
		}
	}

	err := ioutil.WriteFile(pidFile, []byte(fmt.Sprintf("%d", os.Getpid())), 0664)
	if err != nil {
		logger.Fatalf("Unable to write pid file")
	}
}
