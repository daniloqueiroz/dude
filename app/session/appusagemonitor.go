package session

import (
	"context"
	"github.com/BurntSushi/xgb"
	"github.com/daniloqueiroz/dude/app/appusage"
	"github.com/daniloqueiroz/dude/app/system/supervisor"
	"github.com/google/logger"
	"time"
)

func appUsageMonitorSupervisor(rec *appusage.Recorder, supervisor *supervisor.Supervisor) {
	events := make(chan xgb.Event, 1)
	supervisor.AddTask("AppUsageMonitorPeriodicFlusher", func(ctx context.Context) error {
		cTick := time.NewTicker(120 * time.Second)
		defer cTick.Stop()
		for range cTick.C {
			rec.Flush()
		}
		return nil
	})
	supervisor.AddTask("AppUsageMonitorSubscriber", func(ctx context.Context) error {
		logger.Info()
		rec.RegisterListener(events)
		return nil
	})
	supervisor.AddTask("AppUsageMonitor", func(ctx context.Context) error {
		rec.HandleEvents(events)
		return nil
	})

}
