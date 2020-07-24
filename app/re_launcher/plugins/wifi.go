package plugins

import (
	"fmt"
	"github.com/daniloqueiroz/dude/app"
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/daniloqueiroz/dude/app/system/iwd"
	"github.com/godbus/dbus/v5"
	"github.com/google/logger"
	"github.com/xeonx/timeago"
	"time"
)

const WIFI = "wifi"

type wifiAction struct {
}

func (wa *wifiAction) Category() Category {
	return System
}
func (wa *wifiAction) Name() string {
	return WIFI
}
func (wa *wifiAction) Description() string {
	return "Wifi system settings"
}

func wrapWifi(dbus *dbus.Conn, network *iwd.Network) func() {
	return func() {
		defer system.OnPanic("wrapwifi", make(chan error))
		err := network.Connect(dbus)
		if err != nil {
			logger.Errorf("Error connecting to network", err)
		}
	}
}

func (wa *wifiAction) Execute() Result {
	defer system.OnPanic("wifi:Execute", make(chan error))
	var subActions Actions
	subActions = append(subActions, &internalAction{
		name:        "iwctl",
		description: "Wifi manager",
		handler: func() {
			app.NewTerminalApp(system.ExternalAppPath(system.IWCTL))
		},
	})
	// TODO disconnect

	dbus, err := dbus.SystemBus()
	if err != nil {
		logger.Errorf("Unable to load wifi info:", err)
		return &SubActions{SubActions: subActions}
	}
	iwdObj := iwd.New(dbus)

	networks := make(map[string]*iwd.Network)
	for _, network := range iwdObj.Networks {
		networks[network.Name] = &network
	}

	for _, knownNetwork := range iwdObj.KnownNetworks {
		if network, ok := networks[knownNetwork.Name]; ok {
			var lastUsed string
			timeUsed, err := time.Parse(time.RFC3339, knownNetwork.LastConnectedTime)
			if err != nil {
				logger.Errorf("Error formating time", err)
				lastUsed = ""
			} else {
				lastUsed = fmt.Sprintf("(Last used %s)", timeago.English.Format(timeUsed))
			}
			subActions = append(subActions, &internalAction{
				name:        knownNetwork.Name,
				description: fmt.Sprintf("Connect to %s %s", knownNetwork.Name, lastUsed),
				handler:     wrapWifi(dbus, network),
			})
		}
	}

	return &SubActions{SubActions: subActions}
}
