// Gone Time Tracker -or- Where has my time gone?
package appusage

import (
	"fmt"
	"github.com/daniloqueiroz/dude/internal"
	"github.com/google/logger"
	"path"
	"time"
)

func (w Window) String() string {
	return fmt.Sprintf("%s %s", w.Class, w.Name)
}

type Track struct {
	Seen   time.Time
	Spent  time.Duration
	Idle   time.Duration
	Window Window
}

func (t Track) String() string {
	return fmt.Sprintf("%s %s", t.Seen.Format("2006/01/02 15:04:05"), t.Spent)
}

type Recorder struct {
	x           Xorg
	tracks      *TrackStore
	active      *Window
	zzz         bool
	DataDir     string
	CurrentWeek string
}

func (r *Recorder) Snooze(idle time.Duration) {
	if r.active != nil && !r.zzz && r.tracks.Has(*r.active) {
		track, err := r.tracks.Get(*r.active)
		if err == nil {
			track.Idle += idle
			r.tracks.Put(track)
		}
		r.zzz = true
	}
}

func (r *Recorder) Wakeup() {
	if r.active != nil && r.zzz && r.tracks.Has(*r.active) {
		track, err := r.tracks.Get(*r.active)
		if err == nil {
			track.Seen = time.Now()
			r.tracks.Put(track)
		}
		r.zzz = false
	}
}

func (r *Recorder) Update(win Window) {
	logger.Infof("Update Event Received. Active Window %#v / New Window %#v", r.active, win)
	if !r.zzz {
		if r.active != nil && r.tracks.Has(*r.active) {
			track, _ := r.tracks.Get(*r.active)
			track.Spent += time.Since(track.Seen)

			if err := r.tracks.Put(track); err != nil {
				logger.Errorf("Error updating %#v->%#v: %#v", r.active, track, err)
			}
		}

		logger.Infof("Active window updated to %#v", win)
		r.active = &win
		var newTrack *Track
		if r.tracks.Has(win) {
			newTrack, _ = r.tracks.Get(win)
		} else {
			newTrack = &Track{
				Window: win,
			}
		}
		newTrack.Seen = time.Now()
		if err := r.tracks.Put(newTrack); err != nil {
			logger.Errorf("Error updating %#v->%#v: %#v", r.active, newTrack, err)
		}
	}
}

func (r *Recorder) backgroundTasks() {
	go func() {
		defer internal.OnPanic("compacter")
		logger.Info("Starting compact task")
		cTick := time.NewTicker(30 * time.Minute)
		defer cTick.Stop()
		for range cTick.C {
			logger.Info("Compact task running")
			r.tracks.compact()
		}
	}()
	go func() {
		defer internal.OnPanic("report writer")
		logger.Info("Starting report writer task")
		cTick := time.NewTicker(1 * time.Minute)
		defer cTick.Stop()
		for range cTick.C {
			logger.Info("Report task running")
			name := ReportFileName(r.DataDir, r.CurrentWeek)
			report := NewReport(r)
			err := report.WriteToFile(name)
			if err != nil {
				print("reporter!")
				logger.Errorf("Error writing report: %v", err)
			}
		}
	}()
	go func() {
		defer internal.OnPanic("app_usage update")
		logger.Info("Starting update task")
		cTick := time.NewTicker(30 * time.Second)
		defer cTick.Stop()
		for range cTick.C {
			r.Update(*r.active)
		}
	}()
}

func (r *Recorder) Start() {
	defer r.x.Close()
	r.backgroundTasks()
	r.x.Collect(r, time.Minute*5)
}

func NewRecorder(display, dataDir string) (*Recorder, error) {
	X := Connect(display)
	time_now := time.Now()
	year, wk_num := time_now.ISOWeek()
	currentWeek := fmt.Sprintf("%d-w%d", year, wk_num)
	store, err := NewTrackStore(path.Join(dataDir, currentWeek))
	if err != nil {
		return nil, err
	}

	tracker := &Recorder{
		x:           X,
		tracks:      store,
		active:      nil,
		zzz:         false,
		DataDir:     dataDir,
		CurrentWeek: currentWeek,
	}
	// TODO timer to flush active
	return tracker, nil
}
