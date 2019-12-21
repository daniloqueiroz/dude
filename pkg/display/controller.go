package display

import (
	"errors"
	"github.com/daniloqueiroz/dude/internal"
	"github.com/daniloqueiroz/dude/internal/proc"
	"github.com/google/logger"
	"strings"
)

func ApplyProfile(profileName string) error {
	screens := DetectOutputs()
	profile := LoadProfiles(internal.Config.Profiles).GetProfile(profileName)
	if profile == nil {
		return errors.New("no profile found")
	}
	applyProfile(profile, screens)
	return nil
}

func AutoConfigureDisplay() string {
	screens := DetectOutputs()
	var screenNames []string
	for _, s := range screens {
		if s.IsConnected {
			screenNames = append(screenNames, s.Name)
		}
	}
	profiles := LoadProfiles(internal.Config.Profiles)
	selected := profiles.SelectProfile(screenNames...)
	applyProfile(selected, screens)
	return selected.Name
}

func applyProfile(selected *Profile, screens []*Screen) {
	var params []string
	for _, screen := range screens {
		params = append(params, "--output", screen.Name)
		if selected.IsEnabled(screen.Name) && screen.IsConnected {
			conf := selected.GetDisplay(screen.Name)
			if conf.Resolution != "" && conf.Resolution != "auto" {
				params = append(params, "--mode", conf.Resolution)
			} else {
				params = append(params, "--mode", screen.BestMode)
			}
			if conf.ExtraArgs != "" {
				params = append(params, strings.Split(conf.ExtraArgs, " ")...)
			}
		} else {
			params = append(params, "--off")
		}
	}

	logger.Infof("Applying profile %s -> xrandr params %v", selected.Name, params)
	err := proc.NewProcess(internal.Config.AppXrandr, params...).FireAndWait()
	if err != nil {
		logger.Errorf("Error applying profile %s: %v", selected.Name, err)
	}
	SetWallpaper()
}
