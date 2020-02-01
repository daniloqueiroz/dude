package appusage

import (
	"encoding/json"
	"fmt"
	"github.com/google/logger"
	"io/ioutil"
	"path"
	"sort"
	"time"
)

const ReportSuffix = "report.json"

type Report struct {
	ClassRecords  Classes
	Total         time.Duration
	Idle          time.Duration
}

type ClassRecord struct {
	Class   string
	Spent   time.Duration
	Percent float64
}

type Classes []ClassRecord
func (c Classes) Len() int           { return len(c) }
func (c Classes) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c Classes) Less(i, j int) bool { return c[i].Spent < c[j].Spent }

func NewReport(recorder *Recorder) *Report {
	var report Report
	classes := make(map[string]time.Duration)

	for track := range recorder.tracks.Tracks() {
		classes[track.Window.Class] += track.Spent
		report.Total += track.Spent
		report.Idle += track.Idle
	}
	for k, v := range classes {
		report.ClassRecords = append(report.ClassRecords, ClassRecord{
			Class:   k,
			Spent:   v,
			Percent: 100.0 * float64(v) / float64(report.Total)})
	}
	sort.Sort(sort.Reverse(report.ClassRecords))
	return &report
}

func (r *Report) WriteToFile(reportFile string) error {
	logger.Infof("Writing report to %s", reportFile)
	data, err := json.MarshalIndent(r, "", " ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(reportFile, data, 0644)
}


func ReportFileName(destDir, prefix string) string {
	return path.Join(destDir, fmt.Sprintf("%s-%s", prefix, ReportSuffix))
}