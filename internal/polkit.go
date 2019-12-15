package internal

import (
	"github.com/daniloqueiroz/dude/internal/commons"
	"github.com/daniloqueiroz/dude/internal/commons/proc"
	"github.com/google/logger"
)

func StartPolkit() {
	process := proc.NewProcess(commons.Config.AppPolkitAgent)
	if err := process.FireAndKeepAlive(100); err != nil {
		logger.Errorf("Polkit has died: %v", err)
	}
}