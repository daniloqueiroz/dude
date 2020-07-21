package plugins

import (
	"fmt"
	"github.com/daniloqueiroz/dude/app"
	"github.com/daniloqueiroz/dude/app/display"
	"github.com/daniloqueiroz/dude/app/system"
)

const (
	LOCK_SCREEN = "lock-screen"
	SUSPEND     = "suspend"
	REBOOT      = "reboot"
	SHUTDOWN    = "shutdown"
	TERMINAL    = "terminal"
	DISPLAY     = "display"
)

type internalAction struct {
	name        string
	description string
	handler     func()
}

func (ia *internalAction) Category() Category {
	return System
}
func (ia *internalAction) Name() string {
	return ia.name
}
func (ia *internalAction) Description() string {
	return ia.description
}

func (ia *internalAction) Execute() Result {
	ia.handler()
	return Result("la")
}

type internalPlugin struct {
	internalCmds Actions
}

func (a *internalPlugin) Category() Category {
	return System
}

func (a *internalPlugin) FindActions(input string) Actions {
	return FilterAction(input, a.internalCmds)
}

func InternalPluginNew() LauncherPlugin {
	return &internalPlugin{
		internalCmds: loadInternalActions(),
	}
}

func loadInternalActions() Actions {
	return Actions{
		&internalAction{
			name:        LOCK_SCREEN,
			description: "Locks the screen",
			handler: func() {
				system.LockScreen()
			},
		},
		&internalAction{
			name:        SUSPEND,
			description: "Suspends the computer",
			handler: func() {
				system.Suspend()
			},
		},
		&internalAction{
			name:        REBOOT,
			description: "Reboots the computer",
			handler: func() {
				system.Reboot()
			},
		},
		&internalAction{
			name:        SHUTDOWN,
			description: "Turns off the computer",
			handler: func() {
				system.Shutdown()
			},
		},
		&internalAction{
			name:        TERMINAL,
			description: "Starts a new Terminal Window",
			handler: func() {
				app.NewTmuxTerminal()
			},
		},
		&internalAction{
			name:        DISPLAY,
			description: "Load display profile",
			handler: func() {
				profile := display.AutoConfigureDisplay()
				system.SimpleNotification(fmt.Sprintf("Profile %s applied", profile)).Show()
			},
		},
	}
}
