package plugins

import (
	"fmt"
	"github.com/daniloqueiroz/dude/app"
	"github.com/daniloqueiroz/dude/app/system"
	"net/url"
	"strings"
)

type webAction struct {
	name string
	url  string
	desc string
}

func (wa *webAction) Category() Category {
	return Web
}
func (wa *webAction) Name() string {
	return wa.name
}
func (wa *webAction) Description() string {
	return wa.desc
}

func (wa *webAction) Execute() Result {
	app.XDGOpen(wa.url).FireAndForget()
	return Empty{}
}

type webPlugin struct {
}

func (w *webPlugin) Category() Category {
	return Web
}

func (w *webPlugin) FindActions(input string) Actions {
	var fullURL string
	if !strings.HasPrefix("http://", input) || !strings.HasPrefix("https://", input) {
		fullURL = fmt.Sprintf("http://%s", input)
	} else {
		fullURL = input
	}
	return Actions{
		&webAction{
			name: "Web search",
			desc: fmt.Sprintf("Search for '%s'", input),
			url:  fmt.Sprintf("%s?q=%s", system.Config.LauncherWebQueryURL, url.QueryEscape(input)),
		},
		&webAction{
			name: "Open URL",
			desc: fmt.Sprintf("Go to '%s'", input),
			url:  fullURL,
		},
	}
}

func WebPluginNew() LauncherPlugin {
	return &webPlugin{}

}
