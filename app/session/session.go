package session

import (
	"context"
	"github.com/daniloqueiroz/dude/app"
	"github.com/daniloqueiroz/dude/app/appusage"
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/daniloqueiroz/dude/app/system/supervisor"
	"os"
	"syscall"
)

type Session struct {
	supervisor *supervisor.Supervisor
	recorder   *appusage.Recorder
}

func NewSession(r *appusage.Recorder) *Session {
	return &Session{
		supervisor : supervisor.NewSupervisor(),
		recorder: r,
	}
}

func (s *Session) Start() {
	// PID file

	app.FehProc().FireAndForget()

	s.supervisor.AddSigHandler(s.Stop, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	s.supervisor.AddProc(app.CompositorProc())
	s.supervisor.AddProc(app.PolkitProc())
	s.supervisor.AddProc(app.XSSLockProc())
	s.supervisor.AddTask("autostart-apps", s.autostartApps)
	displayMonitorSupervisor(s.supervisor)
	batteryMonitorSupervisor(s.supervisor)
	appUsageMonitorSupervisor(s.recorder, s.supervisor)

	s.supervisor.Start()
	_ = system.SimpleNotification("Session started").Show()
	s.supervisor.Wait()
	os.Exit(0)
}

func (s *Session) Stop() {
	s.supervisor.Stop()
}

func (s *Session) autostartApps(ctx context.Context) error {
	app.AutostartApps()
	<-ctx.Done()
	return nil
}
