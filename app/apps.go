package app

import (
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/daniloqueiroz/dude/app/system/proc"
	"github.com/google/logger"
)

func FehProc() *proc.Process {
	return proc.NewProcess(system.ExternalAppPath(system.FEH), "--bg-fill", "--randomize", system.Config.WallpaperDir)
}

func CompositorProc() *proc.Process {
	return proc.NewProcess(system.ExternalAppPath(system.PICOM), "--backend", "glx", "--vsync")
}

func XSetScreensaverTimeProc() *proc.Process {
	return proc.NewProcess(system.ExternalAppPath(system.XSET), "s", string(system.Config.ScreenSaverTimeoutSecs))
}

func XSSLockProc() *proc.Process {
	return proc.NewProcess(system.ExternalAppPath(system.XSS_LOCK), system.ExternalAppPath(system.XSECURELOCK))
}

func XDGOpen(target string) *proc.Process {
	logger.Infof("xdg-open %s", target)
	return proc.NewProcess(system.ExternalAppPath(system.XDG_OPEN), target)
}
