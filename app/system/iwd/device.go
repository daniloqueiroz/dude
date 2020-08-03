package iwd

import (
	"github.com/godbus/dbus/v5"
)

const (
	objectDevice = "net.connman.iwd.Device"
)

// Device refers to the iwd network device like "wlan0" for example: /net/connman/iwd/0/4
type Device struct {
	Path    dbus.ObjectPath
	Adapter dbus.ObjectPath
	Address string
	Mode    string
	Name    string
	Powered bool
	conn    *dbus.Conn
}

func (d *Device) SetPowered(powered bool) error {
	objectManager := d.conn.Object(objectIwd, d.Path)
	return objectManager.Call("org.freedesktop.DBus.Properties.Set", 0, objectDevice, "Powered", dbus.MakeVariant(powered)).Store()
}
