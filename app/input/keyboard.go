package input

import (
	"errors"
	"fmt"
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/daniloqueiroz/dude/app/system/proc"
)

type Keyboard struct {
	Name      string
	layout    string
	model     string
	variant   string
	isDefault bool
}

func (k Keyboard) setxkbmapParams() []string {
	var sb []string
	if k.model != "" {
		sb = append(sb, "-model")
		sb = append(sb, k.model)
	}
	if k.layout != "" {
		sb = append(sb, "-layout")
		sb = append(sb, k.layout)
	}
	if k.variant != "" {
		sb = append(sb, "-variant")
		sb = append(sb, k.variant)
	}
	return sb
}

type Keyboards []Keyboard

func (k Keyboards) GetKeyboard(kbName string) *Keyboard {
	for _, kb := range k {
		if kb.Name == kbName {
			return &kb
		}
	}
	return nil
}

func (k Keyboards) GetDefaultKeyboard() *Keyboard {
	for _, kb := range k {
		if kb.isDefault {
			return &kb
		}
	}
	return nil
}


func LoadKeyboards(config map[string]interface{}) Keyboards {
	keyboards := make([]Keyboard, 0)
	for keyName, s := range config {
		var keyboard = Keyboard{
			Name: keyName,
			isDefault: false,
		}
		ss := s.(map[string]interface{})
		for prop, value := range ss {
			if prop == "layout" {
				keyboard.layout = value.(string)
			} else if prop == "model" {
				keyboard.model = value.(string)
			} else if prop == "variant" {
				keyboard.variant = value.(string)
			} else if prop == "default" && value.(bool) {
				keyboard.isDefault = true
			}
		}
		keyboards = append(keyboards, keyboard)
	}
	return keyboards
}

func SetKeyboard(kbName string) error {
	kbs := LoadKeyboards(system.Config.Keyboards)
	kb := kbs.GetKeyboard(kbName)
	if kb == nil {
		return errors.New(fmt.Sprintf("no keyboard %s found", kbName))
	}
	return setxkbmap(kb)
}

func SetDefaultKeyboard() (*Keyboard, error) {
	kbs := LoadKeyboards(system.Config.Keyboards)
	kb := kbs.GetDefaultKeyboard()
	if kb == nil {
		return nil, errors.New("no default keyboard found")
	}
	err := setxkbmap(kb)
	return kb, err
}

func setxkbmap(kb *Keyboard) error {
	return proc.NewProcess(system.Config.AppSetxkbmap, kb.setxkbmapParams()...).FireAndWait()
}