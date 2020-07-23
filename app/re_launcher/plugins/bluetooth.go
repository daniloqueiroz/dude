package plugins

import (
	"fmt"
	"github.com/daniloqueiroz/dude/app"
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/google/logger"
	"github.com/muka/go-bluetooth/bluez/profile/adapter"
)

const BLUETOOTH = "bluetooth"

type bluetoothAction struct {
}

func (ba *bluetoothAction) Category() Category {
	return System
}
func (ba *bluetoothAction) Name() string {
	return BLUETOOTH
}
func (ba *bluetoothAction) Description() string {
	return "Bluetooth system settings"
}

func wrapBluetooth(fn func() error) func() {
	return func() {
		err := fn()
		if err != nil {
			logger.Errorf("Error performing bluetooth device action", err)
		}
	}
}

func (ba *bluetoothAction) Execute() Result {
	var subActions Actions
	subActions = append(subActions, &internalAction{
		name:        "bluetoothctl",
		description: "Bluetooth manager",
		handler: func() {
			app.NewTerminalApp(system.Config.AppBluetoothCtl)
		},
	})
	btAdapter, err := adapter.GetDefaultAdapter()
	if err != nil {
		logger.Errorf("Unable to load bt devices:", err)
		return &SubActions{SubActions: subActions}
	}
	devices, err := btAdapter.GetDevices()
	if err != nil {
		logger.Errorf("Unable to load bt devices:", err)
		return &SubActions{SubActions: subActions}
	}

	for _, device := range devices {
		var desc string
		var fn func() error
		if device.Properties.Connected {
			desc = fmt.Sprintf("Disconnect from %s", device.Properties.Name)
			fn = device.Disconnect
		} else {
			desc = fmt.Sprintf("Connect to %s", device.Properties.Name)
			fn = device.Connect
		}
		subActions = append(subActions, &internalAction{
			name:        device.Properties.Name,
			description: desc,
			handler:     wrapBluetooth(fn),
		})
	}

	return &SubActions{SubActions: subActions}
}
