package system

import (
	"github.com/google/logger"
	"os/user"
)

func CurrentUserID() string {
	usr, err := user.Current()
	if err != nil {
		logger.Fatal("Unable to detect home dir", err)
	}
	return usr.Username
}

func HomeDir() string {
	usr, err := user.Current()
	if err != nil {
		logger.Fatal("Unable to detect home dir", err)
	}
	return usr.HomeDir
}
