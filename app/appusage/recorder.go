// Gone Time Tracker -or- Where has my time gone?
package appusage

import (
	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/screensaver"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/google/logger"
	"time"
)

type Track struct {
	Seen   time.Time
	Window Window
}

type Recorder struct {
	x       Xorg
	journal *Journal
	active  *Track
	zzz     bool
}

func (r *Recorder) Update(win Window) {
	if !r.zzz {
		if r.active != nil {
			event := r.active

			_ = r.journal.Add(
				Event{
					Spent:   time.Since(event.Seen),
					AppName: event.Window.Class,
				})
		}

		r.active = &Track{
			Seen:   time.Now(),
			Window: win,
		}
	}
}

func (r *Recorder) Snooze() {
	if r.active != nil && !r.zzz {
		r.Update(r.active.Window)
		r.zzz = true
	}
}

func (r *Recorder) Wakeup() {
	if r.active != nil && r.zzz {
		r.active.Seen = time.Now()
		r.zzz = false
	}
}

func (r *Recorder) Flush() {
	logger.Info("Flush active window usage")
	r.Update(r.active.Window)
}

func (r *Recorder) RegisterListener(events chan xgb.Event) {
	defer r.x.Close()
	defer close(events)
	r.x.Subscribe(events)
}

func (r *Recorder) HandleEvents(events chan xgb.Event) {
	if win, ok := r.x.window(); ok {
		r.Update(win)
	}

	timeout := time.Minute*5
	for {
		select {
		case event := <-events:
			switch e := event.(type) {
			case xproto.PropertyNotifyEvent:
				if win, ok := r.x.window(); ok {
					r.Wakeup()
					r.Update(win)
				}
			case screensaver.NotifyEvent:
				switch e.State {
				case screensaver.StateOn:
					r.Snooze()
				default:
					r.Wakeup()
				}
			}
		case <-time.After(timeout):
			r.Snooze()
		}
	}
}

func (r *Recorder) Start() {
	events := make(chan xgb.Event, 1)
	go r.RegisterListener(events)
	go r.HandleEvents(events)
}

func NewRecorder(display string, store *Journal) (*Recorder, error) {
	X := Connect(display)

	tracker := &Recorder{
		x:       X,
		journal: store,
		active:  nil,
		zzz:     false,
	}
	return tracker, nil
}
