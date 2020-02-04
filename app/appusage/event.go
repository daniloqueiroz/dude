package appusage

import (
	"encoding/json"
	"time"
)

type Event struct {
	AppName string
	Spent   time.Duration
}

type EventSerializer struct{}

func (EventSerializer) serialize(entry interface{}) ([]byte, error) {
	data, err := json.Marshal(entry)
	if err != nil {

		return nil, err
	}
	return data, nil
}

func (EventSerializer) deserialize(data []byte) (interface{}, error) {
	var track Track
	err := json.Unmarshal(data, &track)
	if err != nil {
		return nil, err
	}
	return track, nil
}
