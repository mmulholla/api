package common

import (
	"fmt"

	schema "github.com/devfile/api/v2/pkg/apis/workspaces/v1alpha2"
)

// EventsAdded add event to to test schema and nitofy a registered follower
func (devfile *TestDevfile) EventsAdded(events *schema.Events) {
	LogInfoMessage(fmt.Sprintf("events added"))
	devfile.SchemaDevFile.Events = events
	if devfile.Follower != nil {
		devfile.Follower.AddEvent(*events)
	}
}

// EventsAdded notify a registered follower thst the ebvents have been updated
func (devfile *TestDevfile) EventsUpdated(events *schema.Events) {
	LogInfoMessage(fmt.Sprintf("events updated"))
	if devfile.Follower != nil {
		devfile.Follower.UpdateEvent(*events)
	}
}

// AddCommand creates a command of a specified type in a schema structure and pupulates it with random attributes
func (devfile *TestDevfile) AddEvents() schema.Events {
	events := schema.Events{}
	devfile.EventsAdded(&events)
	devfile.SetEventsValues(&events)
	return events
}

func (devfile *TestDevfile) SetEventsValues(events *schema.Events) {
	if GetRandomDecision(4,1) {
		numPreStart := GetRandomNumber(1, 5)
		for i := 0; i < numPreStart; i++ {
			events.PreStart = append(events.PreStart, devfile.AddCommand(schema.ApplyCommandType).Id)
		}
	}
	if GetRandomDecision(4,1) {
		numPostStart := GetRandomNumber(1, 5)
		for i := 0; i < numPostStart; i++ {
			events.PostStart = append(events.PostStart, devfile.AddCommand(schema.ExecCommandType).Id)
		}
	}
	if GetRandomDecision(4,1) {
		numPreStop := GetRandomNumber(1, 5)
		for i := 0; i < numPreStop; i++ {
			events.PreStop = append(events.PreStop, devfile.AddCommand(schema.ExecCommandType).Id)
		}
	}
	if GetRandomDecision(4,1) {
		numPostStart := GetRandomNumber(1, 5)
		for i := 0; i < numPostStart; i++ {
			events.PostStop = append(events.PostStop, devfile.AddCommand(schema.ApplyCommandType).Id)
		}
	}
	devfile.EventsUpdated(events)
}