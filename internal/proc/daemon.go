package proc

import (
	"bytes"
	"github.com/google/logger"
	"os"
	"syscall"
)

func DaemonExec(wd *Watchdog, name string) {
	dudePath, err := getExecutablePath()
	if err != nil {
		logger.Fatal("Unable to locate dude binary", err)
	}
	logger.Infof("Launching dude daemon %s", name)
	cmd := NewProcess(dudePath, "daemon", name)
	wd.Supervise(cmd)
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
