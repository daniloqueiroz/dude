package app

import (
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/daniloqueiroz/dude/app/system/proc"
)

func FehProc() *proc.Process {
	return proc.NewProcess(system.Config.AppFeh, "--bg-fill", "--randomize", system.Config.WallpaperDir)
}

func CompositorProc() *proc.Process {
	return proc.NewProcess(system.Config.AppCompositor, "--backend", "glx", "--vsync", "-d", ":0")
}

func XSetScreensaverTimeProc() *proc.Process {
	return proc.NewProcess(system.Config.AppXset, "s", string(system.Config.ScreenSaverTimeoutSecs))
}

func XSSLockProc() *proc.Process {
	return proc.NewProcess(system.Config.AppXssLock, system.Config.AppXsecurelock)
}