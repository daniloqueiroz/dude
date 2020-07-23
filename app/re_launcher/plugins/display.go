package plugins

import (
	"fmt"
	"github.com/daniloqueiroz/dude/app/display"
)

const DISPLAY = "display"

type displayAction struct {
}

func (da *displayAction) Category() Category {
	return System
}
func (da *displayAction) Name() string {
	return DISPLAY
}
func (da *displayAction) Description() string {
	return "Display system settings"
}

func (da *displayAction) Execute() Result {
	profiles := display.LoadProfiles()
	var subActions Actions
	subActions = append(subActions, &internalAction{
		name:        "auto",
		description: "Auto apply display profile",
		handler: func() {
			display.AutoConfigureDisplay()
		},
	})
	for _, profile := range profiles {
		subActions = append(subActions, &internalAction{
			name:        profile.Name,
			description: fmt.Sprintf("Apply display profile %s", profile.Name),
			handler: func() {
				display.ApplyProfile(profile.Name)
			},
		})
	}

	return &SubActions{SubActions: subActions}
}
