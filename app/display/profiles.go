package display

import (
	"github.com/daniloqueiroz/dude/app/system"
	"sort"
	"strings"
)

type DisplayConf struct {
	DisplayName string
	Resolution  string
	ExtraArgs   string
}

type Profile struct {
	Name          string
	displays      map[string]DisplayConf
	displaysNames string
}

func (p Profile) AppliesTo(lookup string) bool {
	return p.displaysNames == lookup
}

func (p Profile) IsEnabled(displayName string) bool {
	_, found := p.displays[strings.ToLower(displayName)]
	return found
}

func (p Profile) GetDisplay(displayName string) *DisplayConf {
	entry, found := p.displays[strings.ToLower(displayName)]
	if !found {
		return nil
	}
	return &entry
}

type Profiles []Profile

func (p Profiles) SelectProfile(names ...string) *Profile {
	lookup := strings.Join(names, ":")
	lookup = strings.ToLower(lookup)
	for _, profile := range p {
		if profile.AppliesTo(lookup) {
			return &profile
		}
	}
	return nil
}

func (p Profiles) GetProfile(profileName string) *Profile {
	for _, profile := range p {
		if profile.Name == profileName {
			return &profile
		}
	}
	return nil
}

func LoadProfiles() Profiles {
	profilesConfig := system.Config.Profiles
	profiles := make([]Profile, 0)
	for profileName, s := range profilesConfig {
		ss := s.(map[string]interface{})

		displays := make(map[string]DisplayConf)
		names := make([]string, 0)
		for displayName, c := range ss {
			conf := c.(map[string]interface{})
			resolution, found := conf["resolution"]
			if !found {
				resolution = ""
			}
			extra, found := conf["extra_args"]
			if !found {
				extra = ""
			}

			displays[displayName] = DisplayConf{
				DisplayName: displayName,
				Resolution:  resolution.(string),
				ExtraArgs:   extra.(string),
			}

			names = append(names, displayName)

		}
		sort.Strings(names)

		profiles = append(profiles, Profile{
			Name:          profileName,
			displays:      displays,
			displaysNames: strings.Join(names, ":"),
		})
	}
	return profiles
}
