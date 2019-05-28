package main

import (
	"encoding/json"
	"fmt"
	"time"

	"enix.io/banana/src/models"
	"k8s.io/klog"
)

// routineCmd : Command implementation for 'daemon'
type routineCmd struct{}

// newRoutineCmd : Creates backup command from command line args
func newRoutineCmd(*launchArgs) (*routineCmd, error) {
	return &routineCmd{}, nil
}

// execute : Start the backup using specified backend
func (cmd *routineCmd) execute(config *models.Config) error {
	sendMessageToMonitor("routine_start", config, cmd, "")
	klog.Info("starting banana routine")

	state := &State{}
	err := state.loadFromDisk(config)
	if err != nil {
		sendMessageToMonitor("routine_crashed", config, cmd, err.Error())
		return err
	}

	err = cmd.runTasks(state, config)
	if err != nil {
		sendMessageToMonitor("routine_failed", config, cmd, err.Error())
		return err
	}

	err = state.saveToDisk(config)
	if err != nil {
		sendMessageToMonitor("routine_crashed", config, cmd, err.Error())
		return err
	}

	sendMessageToMonitor("routine_done", config, cmd, "")
	return nil
}

// jsonMap : Convert struct to an anonymous map with given JSON keys
func (cmd *routineCmd) jsonMap() (out map[string]interface{}) {
	raw, _ := json.Marshal(cmd)
	json.Unmarshal(raw, &out)
	return
}

func (cmd *routineCmd) runTasks(state *State, config *models.Config) error {
	for name, schedule := range config.Schedule {
		backupState, exists := state.LastBackups[name]
		if !exists {
			cmd.doBackup(name, state, config)
		} else {
			timeSinceLastBackup := time.Since(backupState.Time)
			interval := time.Duration(schedule.Interval * int(time.Hour) * 24)
			if timeSinceLastBackup > interval {
				cmd.doBackup(name, state, config)
			}
		}
	}

	return nil
}

func (cmd *routineCmd) doBackup(name string, state *State, config *models.Config) {
	state.LastBackups[name] = BackupState{Time: time.Now()}
	fmt.Printf("backup %s now\n", name)
}
