package system

import (
	"github.com/godbus/dbus"
	"github.com/google/logger"
)

type Notification interface {
	Show()
}

type NotificationEvent struct {
	Title   string
	Message string
	Icon    string
}

func (n NotificationEvent) Show() error {
	conn, err := dbus.SessionBus()
	if err != nil {
		logger.Error("Unable to connect to dbus", err)
		return err
	}
	obj := conn.Object("org.freedesktop.Notifications", dbus.ObjectPath("/org/freedesktop/Notifications"))

	call := obj.Call("org.freedesktop.Notifications.Notify", 0, "dude", uint32(0), n.Icon, n.Title, n.Message, []string{}, map[string]dbus.Variant{}, int32(-1))
	if call.Err != nil {
		logger.Error("Unable to send notification", call.Err)
		return call.Err
	}
	logger.Info("Notification sent")
	return nil
}

func SimpleNotification(message string) NotificationEvent {
	notification := NotificationEvent{}
	notification.Message = message
	notification.Title = "dude"
	notification.Icon = Config.DudeIcon
	return notification
}

func TitleNotification(title, message string) NotificationEvent {
	notification := NotificationEvent{}
	notification.Message = message
	notification.Title = title
	notification.Icon = Config.DudeIcon
	return notification
}
