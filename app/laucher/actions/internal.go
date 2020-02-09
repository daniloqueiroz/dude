package actions

import (
	"fmt"
	"github.com/daniloqueiroz/dude/app"
	"github.com/daniloqueiroz/dude/app/display"
	"github.com/daniloqueiroz/dude/app/laucher"
	"github.com/daniloqueiroz/dude/app/system"
)

const (
	LOCK_SCREEN     = "lock-screen"
	SUSPEND         = "suspend"
	TERMINAL        = "terminal"
	DISPLAY         = "display"
)

type Internal struct {
	internalCmds laucher.Actions
}

func (i *Internal) Find(input string) laucher.Actions {
	if i.internalCmds == nil {
		i.loadInternalActions()
	}
	return laucher.FilterAction(input, i.internalCmds)
}

func (i *Internal) loadInternalActions() {
	// :display [single, mirror, auto]
	// :shutdown
	// :volume [up, down, mute, mic(?)]
	// :brightness [up, down]
	// :keyboard <layout> -> modifies keyboard layout
	// launcher only operations
	// :kill <program>
	// :pass
	// ::-> window switch
	// :o <whatever> -> xdg-open
	// :e <file> -> howl <file>
	// :! <cmd> -> execute command on terminal
	i.internalCmds = laucher.Actions{
		{
			Details: laucher.ActionMeta{
				Name:       LOCK_SCREEN,
				Description: "Locks the screen",
				Category:    laucher.System,
			},
			Exec:  func() {
				system.LockScreen()
			},
		},
		{
			Details: laucher.ActionMeta{
				Name:       SUSPEND,
				Description: "Suspends the computer",
				Category:    laucher.System,
			},
			Exec:  func() {
				system.Suspend()
			},
		},
		{
			Details: laucher.ActionMeta{
				Name:       TERMINAL,
				Description: "Starts a new Terminal Window",
				Category:    laucher.System,
			},
			Exec:  func() {
				app.NewTmuxTerminal()
			},
		},
		{
			Details: laucher.ActionMeta{
				Name:       DISPLAY,
				Description: "Load display profile",
				Category:    laucher.System,
			},
			Exec:  func() {
				profile := display.AutoConfigureDisplay()
				system.SimpleNotification(fmt.Sprintf("Profile %s applied", profile)).Show()
			},
		},
	}
}
