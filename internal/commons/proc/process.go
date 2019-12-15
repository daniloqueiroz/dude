package proc

import (
	"errors"
	"github.com/google/logger"
	"os/exec"
)

type ProcState string

const (
	NotStarted ProcState = "not_started"
	Running    ProcState = "running"
	Success    ProcState = "success"
	Failed     ProcState = "failed"
)

type Process struct {
	cmd           string
	args          []string
	goCmd         *exec.Cmd
	output        string
	processEndChn chan ProcState
	state         ProcState
}

func NewProcess(cmd string, args ...string) *Process {
	return &Process{
		cmd:           cmd,
		args:          args,
		goCmd:         nil,
		output:        "",
		state:         NotStarted,
		processEndChn: make(chan ProcState),
	}
}

func (p *Process) supervisor() {
	logger.Infof("Starting process %s %s", p.cmd, p.args)
	output, err := p.goCmd.CombinedOutput()
	if err != nil {
		logger.Errorf("Process %s %s failed: %v", p.cmd, p.args, err)
		p.state = Failed
	} else {
		logger.Infof("Process %s %s ended successfully: %s", p.cmd, p.args, output)
		p.state = Success
	}
	p.output = string(output)
	p.goCmd = nil
	p.processEndChn <- p.state
}

func (p *Process) fireWithSupervisor() error {
	if p.state != NotStarted {
		return errors.New("process is already started")
	}
	p.state = Running
	p.goCmd = exec.Command(p.cmd, p.args...)
	go p.supervisor()
	return nil
}

func (p *Process) State() ProcState {
	return p.state
}

func (p *Process) FireAndForget() error {
	if p.state != NotStarted {
		return errors.New("process is already started")
	}
	p.state = Running
	p.goCmd = exec.Command(p.cmd, p.args...)
	return p.goCmd.Start()
}

func (p *Process) FireAndWait() error {
	_, err := p.FireAndWaitForOutput()
	return err
}

func (p *Process) FireAndWaitForOutput() (string, error) {
	if p.state == NotStarted {
		if err := p.fireWithSupervisor(); err != nil {
			return "", err
		}
	}

	if  p.state == Running {
		logger.Infof("Waiting process result of %s %s", p.cmd, p.args)
		<-p.processEndChn
	}

	if p.state == Success {
		return p.output, nil
	} else {
		return p.output, errors.New("process ended with an error")
	}
}

func (p *Process) FireAndKeepAlive(maxRestarts int) error {
	if maxRestarts == 0 {
		return errors.New("max number of restarted reached")
	} else if p.state != Running {
		_ = p.fireWithSupervisor()
	}

	state := <-p.processEndChn
	logger.Infof("Process ended with state %v, trying to restart", state)
	return p.FireAndKeepAlive(maxRestarts - 1)
}
