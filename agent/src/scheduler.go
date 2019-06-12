package main

import (
	"encoding/json"
	"fmt"
	"time"

	"enix.io/banana/src/models"
	"github.com/imdario/mergo"
	"k8s.io/klog"
)

// routineCmd : Command implementation for 'daemon'
type routineCmd struct{}

// newRoutineCmd : Creates routine command from command line args
func newRoutineCmd(*launchArgs) (*routineCmd, error) {
	return &routineCmd{}, nil
}

// execute : Start the routine using specified backend
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
	for name, schedule := range config.ScheduledBackups {
		backupState, exists := state.LastBackups[name]
		fullConfig := schedule
		mergo.Merge(&fullConfig.Config, config)
		if !exists || backupState.Status == "Failed" {
			err := cmd.doBackup(name, state, &fullConfig)
			if err != nil {
				return err
			}
		} else {
			timeSinceLastBackup := time.Since(backupState.Time)
			interval := time.Duration(schedule.Interval * float32(time.Hour) * 24)
			if timeSinceLastBackup > interval {
				err := cmd.doBackup(name, state, &fullConfig)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (cmd *routineCmd) doBackup(name string, state *State, config *models.ScheduledBackupConfig) error {
	if state.LastBackups[name] == nil {
		state.LastBackups[name] = &BackupState{}
	}

	state.LastBackups[name].Time = time.Now()
	state.LastBackups[name].Status = "Failed"
	klog.Infof("backing up %s", name)

	if config.Target == "" {
		return fmt.Errorf("missing target directory in schedule for backup %s", name)
	}

	typ := "incremental"
	if state.LastBackups[name].Type == "" || state.LastBackups[name].IncrCountSinceLastFull >= config.FullEvery-1 {
		typ = "full"
	}

	backupCmd := &backupCmd{
		Type:   typ,
		Name:   name,
		Target: config.Target,
	}
	backupCmd.execute(&config.Config)

	state.LastBackups[name].Status = "Success"
	state.LastBackups[name].Type = typ
	if typ == "full" {
		state.LastBackups[name].IncrCountSinceLastFull = 0
	} else {
		state.LastBackups[name].IncrCountSinceLastFull++
	}

	return nil
}
