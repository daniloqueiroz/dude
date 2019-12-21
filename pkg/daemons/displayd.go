package daemons

import (
	"github.com/BurntSushi/xgb"
	"github.com/daniloqueiroz/dude/internal/proc"
	"github.com/daniloqueiroz/dude/pkg/display"
	"github.com/google/logger"
	"syscall"
	"time"
)

func Displayd() {
	daemon := proc.NewDaemon(displayd)
	daemon.Start(syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
}

func displayd() {
	logger.Info("Starting displayd")
	// give it a bit of time before configuring output
	time.Sleep(2 * time.Second)
	display.AutoConfigureDisplay()
	chn := make(chan xgb.Event)
	go func() {
		for {
			<-chn
			logger.Infof("Xrandr event received, autoconfiguring displays")
			display.AutoConfigureDisplay()
		}
	}()
	go display.ListenOutputEvents(chn)
	logger.Info("Displayd is running")
}
