package proc

import (
	"errors"
	"fmt"
	"github.com/google/logger"
	"os/exec"
)

type State string

const (
	NotStarted State = "not_started"
	Running    State = "running"
	Success    State = "success"
	Failed     State = "failed"
)

type Process struct {
	cmd   string
	args  []string
	goCmd *exec.Cmd
	state State
}

func NewProcess(cmd string, args ...string) *Process {
	return &Process{
		cmd:   cmd,
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
			logger.Errorf("Error killing process %: %v", p.cmd, err)
		}
	}
}

func (p *Process) FireAndForget() error {
	if p.State() != NotStarted {
		return errors.New(fmt.Sprintf("Process %s is already started", p.cmd))
	}
	p.goCmd = exec.Command(p.cmd, p.args...)
	p.updateState(Running)
	return p.goCmd.Start()
}

func (p *Process) FireAndWait() error {
	_, err := p.FireAndWaitForOutput()
	return err
}

func (p *Process) FireAndWaitForOutput() (string, error) {
	if p.State() == Running {
		logger.Errorf("Process %s is already running", p.cmd)
		return "", errors.New(fmt.Sprintf("process %s is already running", p.cmd))
	} else {
		p.goCmd = exec.Command(p.cmd, p.args...)
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
