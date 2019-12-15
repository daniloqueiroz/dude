package proc

import (
	"bytes"
	"github.com/google/logger"
	"os"
	"syscall"
)

func DaemonExec(name string) {
	dudePath, err := getExecutablePath()
	if err != nil {
		logger.Fatal("Unable to locate dude binary", err)
	}
	logger.Info("Launching dude daemon %s", name)
	process := NewProcess(dudePath, "daemon", name)
	if err := process.FireAndKeepAlive(100); err != nil {
		logger.Errorf("Daemon %s has died: %v", name, err)
	}
}

func getExecutablePath() (string, error) {
	name := "/proc/self/exe"
	for len := 128; ; len *= 2 {
		b := make([]byte, len)
		n, e := syscall.Readlink(name, b)
		if e != nil {
			return "", &os.PathError{"readlink", name, e}
		}
		if n < len {
			if z := bytes.IndexByte(b[:n], 0); z >= 0 {
				n = z
			}
			return string(b[:n]), nil
		}
	}
}
