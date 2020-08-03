package plugins

import (
	"fmt"
	"github.com/daniloqueiroz/dude/app"
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/google/logger"
	"github.com/muka/go-bluetooth/bluez/profile/adapter"
)

const BLUETOOTH Category = "bluetooth"

type bluetoothAction struct {
}

func (ba *bluetoothAction) Category() Category {
	return System
}
func (ba *bluetoothAction) Name() string {
	return string(BLUETOOTH)
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

func toggleBluetooth(bt *adapter.Adapter1, powerOn bool) func() {
	return func() {
		err := bt.SetPowered(powerOn)
		if err != nil {
			logger.Errorf("Error changing bluetooth power", err)
		}
	}
}

func (ba *bluetoothAction) Execute() Result {
	var subActions Actions
	subActions = append(subActions, &internalAction{
		name:        "bluetoothctl",
		description: "Bluetooth manager",
		handler: func() {
			app.NewTerminalApp(system.ExternalAppPath(system.BLUETOOTHCTL))
		},
		category: BLUETOOTH,
	})
	btAdapter, err := adapter.GetDefaultAdapter()
	if err != nil {
		logger.Errorf("Unable to load bt devices:", err)
		return &SubActions{SubActions: subActions}
	}
	isOn, err := btAdapter.GetPowered()
	if err != nil {
		logger.Errorf("Unable to power on btadapter:", err)
		return &SubActions{SubActions: subActions}
	} else {
		var desc string
		if isOn {
			desc = fmt.Sprintf("Turn bluetooth off")
		} else {
			desc = fmt.Sprintf("Turn bluetooth on")
		}
		subActions = append(subActions, &internalAction{
			name:        "toggle bluetooth",
			description: desc,
			handler:     toggleBluetooth(btAdapter, !isOn),
			category:    BLUETOOTH,
		})
	}

	devices, err := btAdapter.GetDevices()
	if err != nil {
		logger.Errorf("Unable to load bt devices:", err)
		return &SubActions{SubActions: subActions}
	}

	for _, device := range devices {
		var desc string
		var fn func() error
		devName := device.Properties.Name
		if devName == "" {
			continue
		}
		if device.Properties.Connected {
			desc = fmt.Sprintf("Disconnect from %s", devName)
			fn = device.Disconnect
		} else {
			desc = fmt.Sprintf("Connect to %s", devName)
			fn = device.Connect
		}
		subActions = append(subActions, &internalAction{
			name:        devName,
			description: desc,
			handler:     wrapBluetooth(fn),
			category:    BLUETOOTH,
		})
	}

	return &SubActions{SubActions: subActions}
}
