package actions

import (
	"fmt"
	"github.com/daniloqueiroz/dude/internal"
	"github.com/daniloqueiroz/dude/internal/commons/system"
	"github.com/daniloqueiroz/dude/internal/laucher"
)

const (
	INTERNAL_PREFIX = ":"
	LOCK_SCREEN     = "lock-screen"
	SUSPEND         = "suspend"
	TERMINAL        = "terminal"
)

type Internal struct {
	name        string
	description string
	exec        func()
}

func (p Internal) Input() string {
	return fmt.Sprintf("%s%s", INTERNAL_PREFIX, p.name)
}

func (p Internal) Description() string {
	return p.description
}

func (p Internal) Exec() {
	p.exec()
}

func (p Internal) String() string {
	return p.Input()
}

func loadInternalActions(actions map[string]laucher.Action) {
	// :display [single, mirror, auto]
	// :shutdown
	// :volume [up, down, mute, mic(?)]
	// :brightness [up, down]
	// :keyboard <layout> -> modifies keyboard layout
	// :terminal -> launches st
	// launcher only operations
	// :kill <program>
	// :pass
	// ::-> window switch
	// :o <whatever> -> xdg-open
	// :e <file> -> howl <file>
	// :! <cmd> -> execute command on terminal
	commands := []Internal{
		{
			name:        LOCK_SCREEN,
			description: "Locks the screen",
			exec: func() {
				system.LockScreen()
			},
		},
		{
			name:        SUSPEND,
			description: "Suspends the computer",
			exec: func() {
				system.Suspend()
			},
		},
		{
			name:        TERMINAL,
			description: "Starts a new Terminal Window",
			exec: func() {
				internal.NewTmuxTerminal()
			},
		},
	}
	for _, action := range commands {
		actions[action.Input()] = action
	}
}
