// Gone Time Tracker -or- Where has my time gone?
package appusage

import (
	"github.com/daniloqueiroz/dude/app/system"
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

func (r *Recorder) flushTask() {
	go func() {
		defer system.OnPanic("Recorder:flushTask")
		logger.Info("FlushTask started")
		cTick := time.NewTicker(120 * time.Second)
		defer cTick.Stop()
		for range cTick.C {
			logger.Info("Flush active window usage")
			r.Update(r.active.Window)
		}
	}()
}

func (r *Recorder) Start() {
	defer r.x.Close()
	r.flushTask()
	r.x.Collect(r, time.Minute*5)
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
