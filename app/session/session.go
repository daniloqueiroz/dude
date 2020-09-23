package session

import (
	"context"
	"github.com/daniloqueiroz/dude/app"
	"github.com/daniloqueiroz/dude/app/display"
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/daniloqueiroz/dude/app/system/supervisor"
	"os"
	"syscall"
)

type Session struct {
	supervisor *supervisor.Supervisor
}

func NewSession() *Session {
	return &Session{
		supervisor: supervisor.NewSupervisor(),
	}
}

func (s *Session) Start() {
	s.supervisor.AddSigHandler(s.Stop, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	s.supervisor.AddTask("wallpaper", wallpaper)
	s.supervisor.AddProc(app.CompositorProc())
	s.supervisor.AddProc(app.PolkitProc())
	s.supervisor.AddProc(app.Udiskie())
	s.supervisor.AddProc(display.StartScreensaver())
	s.supervisor.AddTask("autostart-apps", autostartApps)
	if system.Config.DisplayAutoConfigEnabled {
		displayMonitorSupervisor(s.supervisor)
	}
	if system.Config.BatteryMonitorEnabled {
		batteryMonitorSupervisor(s.supervisor)
	}

	s.supervisor.Start()
	_ = system.SimpleNotification("Session started").Show()
	s.supervisor.Wait()
	os.Exit(0)
}

func (s *Session) Stop() {
	s.supervisor.Stop()
}

func wallpaper(ctx context.Context) error {
	app.FehProc().FireAndForget()
	<-ctx.Done()
	return nil
}

func autostartApps(ctx context.Context) error {
	app.AutostartApps()
	<-ctx.Done()
	return nil
}
