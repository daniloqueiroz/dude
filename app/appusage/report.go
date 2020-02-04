package appusage

import (
	"github.com/google/logger"
	"sort"
	"time"
)

const ReportSuffix = "report.json"

type Report struct {
	ClassRecords Classes
	Total        time.Duration
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

func NewReport(journal *Journal) (*Report, error) {
	var report Report
	classes := make(map[string]time.Duration)
	receiver := make(chan interface{})
	err := journal.Read(receiver)
	if err != nil {
		return nil, err
	}

	for entry := range receiver {
		event := entry.(Event)
		logger.Infof("Track received: %v", event)
		classes[event.AppName] += event.Spent
		report.Total += event.Spent
	}

	logger.Infof("Classes")
	for k, v := range classes {
		report.ClassRecords = append(report.ClassRecords, ClassRecord{
			Class:   k,
			Spent:   v,
			Percent: 100.0 * float64(v) / float64(report.Total)})
	}
	logger.Infof("Sort")
	sort.Sort(sort.Reverse(report.ClassRecords))
	logger.Infof("Done")
	return &report, nil
}
