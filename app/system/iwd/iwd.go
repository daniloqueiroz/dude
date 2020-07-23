package iwd

import (
	"github.com/godbus/dbus/v5"
)

const (
	objectIwd = "net.connman.iwd"
)

// Iwd is a struct over all major iwd components
type Iwd struct {
	KnownNetworks []KnownNetwork
	Networks      []Network
}

// New parses the net.connman.iwd object index and initializes an iwd object
func New(conn *dbus.Conn) Iwd {
	var objects map[dbus.ObjectPath]map[string]map[string]dbus.Variant
	objectManager := conn.Object(objectIwd, "/")
	objectManager.Call("org.freedesktop.DBus.ObjectManager.GetManagedObjects", 0).Store(&objects)
	i := Iwd{
		make([]KnownNetwork, 0),
		make([]Network, 0),
	}
	for k, v := range objects {
		for resource, obj := range v {
			switch resource {
			case objectKnownNetwork:
				i.KnownNetworks = append(i.KnownNetworks, KnownNetwork{
					Path:              k,
					AutoConnect:       obj["AutoConnect"].Value().(bool),
					Hidden:            obj["Hidden"].Value().(bool),
					LastConnectedTime: obj["LastConnectedTime"].Value().(string),
					Name:              obj["Name"].Value().(string),
					Type:              obj["Type"].Value().(string),
				})
			case objectNetwork:
				i.Networks = append(i.Networks, Network{
					Path:      k,
					Connected: obj["Connected"].Value().(bool),
					Name:      obj["Name"].Value().(string),
					Type:      obj["Type"].Value().(string),
				})
			}
		}
	}
	return i
}
