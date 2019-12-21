package daemons

import (
	"fmt"
	"github.com/daniloqueiroz/dude/internal/proc"
	"github.com/daniloqueiroz/dude/internal/system"
	"github.com/daniloqueiroz/dude/pkg"
	"github.com/daniloqueiroz/dude/pkg/display"
	"github.com/google/logger"
	"github.com/rkoesters/xdg/basedir"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"syscall"
	"time"
)


func Sessiond() {
	writePidFile()
	wd := proc.NewWatchdog()
	display.StartCompositor(wd)
	pkg.StartScreensaver(wd)
	pkg.StartPolkit(wd)
	pkg.AutostartApps()
	proc.LaunchDaemon(wd, "powerd")
	proc.LaunchDaemon(wd, "trackerd")
	proc.LaunchDaemon(wd, "displayd")
	system.SimpleNotification("Session started").Show()

	wd.Start(syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	time.Sleep(1 * time.Second)
	logger.Info("Session ended")
}

func writePidFile() {
	pidFile := path.Join(basedir.RuntimeDir, "sessiond.pid")
	logger.Infof("Sessiond pid file %s", pidFile)
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