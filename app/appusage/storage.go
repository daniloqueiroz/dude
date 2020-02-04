package appusage

import (
	"bufio"
	"github.com/daniloqueiroz/dude/app/system"
	"github.com/google/logger"
	"os"
)

type Serializer interface {
	serialize(entry interface{}) ([]byte, error)
	deserialize(data []byte) (interface{}, error)
}

type Journal struct {
	path       string
	serializer Serializer
	rwFile     *os.File
}

func NewJornal(filename string, serializer Serializer) *Journal {
	return &Journal{
		path:       filename,
		serializer: serializer,
		rwFile:     nil,
	}
}

func (j *Journal) Close() error {
	if j.rwFile != nil {
		_ = j.rwFile.Sync()
		return j.rwFile.Close()
	}
	return nil
}

func (j *Journal) Add(entry interface{}) error {
	data, err := j.serializer.serialize(entry)
	if err != nil {
		logger.Errorf("Error serializing entry: %v", err)
		return err
	}
	if j.rwFile == nil {
		file, err := os.OpenFile(j.path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			logger.Errorf("Error opening file: %v", err)
			return err
		}
		j.rwFile = file
	}
	writer := bufio.NewWriter(j.rwFile)
	_, err = writer.Write(data)
	if err != nil {
		logger.Errorf("Error writing to file: %v", err)
		return err
	}
	_, err = writer.WriteString("\n")
	if err != nil {
		logger.Errorf("Error writing to file: %v", err)
		return err
	}
	err = writer.Flush()
	if err != nil {
		logger.Errorf("Error writing file: %v", err)
		return err
	}
	err = j.rwFile.Sync()
	if err != nil {
		logger.Errorf("Error syncing file: %v", err)
		return err
	}
	return nil
}

func (j *Journal) Read(receiver chan interface{}) error {
	file, err := os.OpenFile(j.path, os.O_RDONLY, 0644)
	if err != nil {
		logger.Errorf("Error opening file: %v", err)
		return err
	}
	go func() {
		logger.Infof("Reading journal file %v", file)
		defer system.OnPanic("Journal:Read")
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			data := scanner.Bytes()
			entry, err := j.serializer.deserialize(data)
			if err != nil {
				logger.Errorf("Error deserializing entry: %v", err)
				continue
			}
			logger.Infof("Entry: %v", file)
			receiver <- entry
		}
		close(receiver)
		err := file.Close()
		if err != nil {
			logger.Errorf("Error closing file: %v", err)
		}
	}()

	return nil
}

func (j *Journal) Compact() {
	// TODO implement it
}
