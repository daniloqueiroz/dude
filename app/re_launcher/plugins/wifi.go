package plugins

import (
	"fmt"
	"github.com/daniloqueiroz/dude/app"
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/daniloqueiroz/dude/app/system/iwd"
	dbus5 "github.com/godbus/dbus/v5"
	"github.com/google/logger"
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

func wrapWifi(dbus *dbus5.Conn, network *iwd.Network) func() {
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
			app.NewTerminalApp(system.Config.AppWifiCtl)
		},
	})
	// TODO disconnect

	dbus, err := dbus5.SystemBus()
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
			subActions = append(subActions, &internalAction{
				name:        knownNetwork.Name,
				description: fmt.Sprintf("Connect to %s", knownNetwork.Name),
				handler:     wrapWifi(dbus, network),
			})
		}
	}

	return &SubActions{SubActions: subActions}
}
