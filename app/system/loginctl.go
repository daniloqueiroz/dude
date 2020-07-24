package system

import (
	"github.com/daniloqueiroz/go-systemd/login1"
	"github.com/google/logger"
)

func LockScreen() error {
	conn, err := login1.New()
	if err != nil {
		logger.Error("Unable to connect to DBUS", err)
		return err
	}
	sessions, err := conn.ListSessions()
	if err != nil {
		logger.Error("Unable to connect get sessions", err)
		return err
	}
	for _, session := range sessions {
		if session.User == CurrentUserID() {
			conn.LockSession(session.ID)
		}
	}
	logger.Info("Session locked")
	return nil
}

func Suspend() {
	conn, err := login1.New()
	if err != nil {
		logger.Error("Unable to connect to DBUS", err)
	}
	conn.Suspend(false)
}

func Reboot() {
	conn, err := login1.New()
	if err != nil {
		logger.Error("Unable to connect to DBUS", err)
	}
	conn.Reboot(false)
}

func Shutdown() {
	conn, err := login1.New()
	if err != nil {
		logger.Error("Unable to connect to DBUS", err)
	}
	conn.PowerOff(false)
}
