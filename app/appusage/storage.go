package appusage

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/google/logger"
	"github.com/prologic/bitcask"
)

const (
	separator = "=::="
)

type TrackStore struct {
	db *bitcask.Bitcask
}

func NewTrackStore(storageDir string) (*TrackStore, error) {
	opts := []bitcask.Option{
		bitcask.WithMaxDatafileSize(10 << 20),
		bitcask.WithSync(true),
		bitcask.WithMaxKeySize(256),
	}
	bc, err := bitcask.Open(storageDir, opts...)
	if err != nil {
		logger.Fatalf("Unable to open bitcask: %v", err)
		return nil, err
	}
	store := &TrackStore{db: bc}
	return store, nil
}

func (t *TrackStore) compact() {
	logger.Infof("Compacting TrackStore")
	t.db.Merge()
}

func (t *TrackStore) getKey(win Window) []byte {
	bytes := make([]byte, 32)
	hash := sha256.Sum256([]byte(fmt.Sprint(win.Class, win.Name)))
	for i, b := range hash {
		bytes[i] = b
	}
	return bytes
}

func (t *TrackStore) Get(win Window) (*Track, error) {
	data, err := t.db.Get(t.getKey(win))
	if err != nil {
		logger.Errorf("Unable to retrieve track %s: %v", t.getKey(win), err)
		return nil, err
	}
	return deserialize(data)
}

func (t *TrackStore) Put(rec *Track) error {
	data, err := serialize(rec)
	if err != nil {
		return err
	}
	return t.db.Put(t.getKey(rec.Window), data)
}

func (t *TrackStore) Has(win Window) bool {
	return t.db.Has(t.getKey(win))
}

func (t *TrackStore) Tracks() chan *Track {
	chn := make(chan *Track)
	go func() {
		defer close(chn)
		for key := range t.db.Keys() {
			data, _ := t.db.Get(key)
			track, _ := deserialize(data)
			chn <- track
		}
	}()
	return chn
}

func (t *TrackStore) Len() int {
	return t.db.Len()
}

func serialize(track *Track) ([]byte, error) {
	data, err := json.Marshal(track)
	if err != nil {
		logger.Errorf("Unable to serialize track: %v", err)
		return nil, err
	}
	return data, nil
}

func deserialize(data []byte) (*Track, error) {
	var track Track
	err := json.Unmarshal(data, &track)
	if err != nil {
		logger.Errorf("Unable to deserialize track: %v", err)
		return nil, err
	}
	return &track, nil
}
