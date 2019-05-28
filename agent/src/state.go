package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"

	"enix.io/banana/src/models"
)

// BackupState : Represents the state of a backup, including the last time it ran
type BackupState struct {
	Time time.Time `json:"time"`
}

// State : Represents the state of an agent, aka. its last backups
type State struct {
	Version     int8                   `json:"version"`
	LastBackups map[string]BackupState `json:"last_backups"`
}

func (s *State) loadFromDisk(config *models.Config) error {
	state, err := ioutil.ReadFile(config.StatePath)
	if err != nil {
		if os.IsNotExist(err) {
			dir, _ := path.Split(config.StatePath)
			os.Mkdir(dir, 00644)
			_, err := os.Create(config.StatePath)
			if err != nil {
				return err
			}

			s.Version = 1
			s.LastBackups = make(map[string]BackupState)
			return nil
		}

		return err
	}

	err = json.Unmarshal(state, s)
	if err != nil {
		return err
	}

	if s.Version < 1 {
		s.Version = 1
	} else if s.Version != 1 {
		return fmt.Errorf("incompatible state file version %d", s.Version)
	}

	return nil
}

func (s *State) saveToDisk(config *models.Config) error {
	bytes, err := json.Marshal(s)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(config.StatePath, bytes, 00644)
}
