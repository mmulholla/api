package common

import (
	"fmt"

	schema "github.com/devfile/api/v2/pkg/apis/workspaces/v1alpha2"
)

// commandAdded adds a new command to the test schema data and notifies the follower
func (devfile *TestDevfile) commandAdded(command schema.Command) {
	LogInfoMessage(fmt.Sprintf("command added Id: %s", command.Id))
	devfile.SchemaDevFile.Commands = append(devfile.SchemaDevFile.Commands, command)
	if devfile.Follower != nil {
		devfile.Follower.AddCommand(command)
	}
}

// commandUpdated and notifies the follower of the command which has been updated
func (devfile *TestDevfile) commandUpdated(command schema.Command) {
	LogInfoMessage(fmt.Sprintf("command updated Id: %s", command.Id))
	if devfile.Follower != nil {
		devfile.Follower.UpdateCommand(command)
	}
}

// addEnv creates and returns a specifed number of env attributes in a schema structure
func addEnv(numEnv int) []schema.EnvVar {
	commandEnvs := make([]schema.EnvVar, numEnv)
	for i := 0; i < numEnv; i++ {
		commandEnvs[i].Name = "Name_" + GetRandomString(5, false)
		commandEnvs[i].Value = "Value_" + GetRandomString(5, false)
		LogInfoMessage(fmt.Sprintf("Add Env: %s", commandEnvs[i]))
	}
	return commandEnvs
}

// addAttributes creates returns a specifed number of attributes in a schema structure
func addAttributes(numAtrributes int) map[string]string {
	attributes := make(map[string]string)
	for i := 0; i < numAtrributes; i++ {
		AttributeName := "Name_" + GetRandomString(6, false)
		attributes[AttributeName] = "Value_" + GetRandomString(6, false)
		LogInfoMessage(fmt.Sprintf("Add attribute : %s = %s", AttributeName, attributes[AttributeName]))
	}
	return attributes
}

// addGroup creates and returns a group in a schema structure
func (devfile *TestDevfile) addGroup() *schema.CommandGroup {

	commandGroup := schema.CommandGroup{}
	commandGroup.Kind = GetRandomGroupKind()
	LogInfoMessage(fmt.Sprintf("group Kind: %s, default already set %t", commandGroup.Kind, devfile.GroupDefaults[commandGroup.Kind]))
	// Ensure only one and at least one of each type are labelled as default
	if !devfile.GroupDefaults[commandGroup.Kind] {
		devfile.GroupDefaults[commandGroup.Kind] = true
		commandGroup.IsDefault = true
	} else {
		commandGroup.IsDefault = false
	}
	LogInfoMessage(fmt.Sprintf("group isDefault: %t", commandGroup.IsDefault))
	return &commandGroup
}

// AddCommand creates a command of a specified type in a schema structure and pupulates it with random attributes
func (devfile *TestDevfile) AddCommand(commandType schema.CommandType) schema.Command {

	var command *schema.Command
	if commandType == schema.ExecCommandType {
		command = devfile.createExecCommand()
		devfile.SetExecCommandValues(command)
	} else if commandType == schema.CompositeCommandType {
		command = devfile.createCompositeCommand()
		devfile.SetCompositeCommandValues(command)
	} else if commandType == schema.ApplyCommandType {
		command = devfile.createApplyCommand()
		devfile.SetApplyCommandValues(command)
	}
	return *command
}

// createExecCommand creates and returns an empty exec command in a schema structure
func (devfile *TestDevfile) createExecCommand() *schema.Command {

	LogInfoMessage("Create an exec command :")
	command := schema.Command{}
	command.Id = GetRandomUniqueString(8, true)
	LogInfoMessage(fmt.Sprintf("command Id: %s", command.Id))
	command.Exec = &schema.ExecCommand{}
	devfile.commandAdded(command)
	return &command

}

// SetExecCommandValues randomly sets exec command attribute to random values
func (devfile *TestDevfile) SetExecCommandValues(command *schema.Command) {

	execCommand := command.Exec

	// exec command must be mentioned by a container component
	execCommand.Component = devfile.GetContainerName()

	execCommand.CommandLine = GetRandomString(4, false) + " " + GetRandomString(4, false)
	LogInfoMessage(fmt.Sprintf("....... commandLine: %s", execCommand.CommandLine))

	// If group already leave it to make sure defaults are not deleted or added
	if execCommand.Group == nil {
		if GetRandomDecision(2, 1) {
			execCommand.Group = devfile.addGroup()
		}
	}

	if GetBinaryDecision() {
		execCommand.Label = GetRandomString(12, false)
		LogInfoMessage(fmt.Sprintf("....... label: %s", execCommand.Label))
	} else {
		execCommand.Label = ""
	}

	if GetBinaryDecision() {
		execCommand.WorkingDir = "./tmp"
		LogInfoMessage(fmt.Sprintf("....... WorkingDir: %s", execCommand.WorkingDir))
	} else {
		execCommand.WorkingDir = ""
	}

	execCommand.HotReloadCapable = GetBinaryDecision()
	LogInfoMessage(fmt.Sprintf("....... HotReloadCapable: %t", execCommand.HotReloadCapable))

	if GetBinaryDecision() {
		execCommand.Env = addEnv(GetRandomNumber(1, 4))
	} else {
		execCommand.Env = nil
	}
	devfile.commandUpdated(*command)

}

// createCompositeCommand creates an empty composite command in a schema structure
func (devfile *TestDevfile) createCompositeCommand() *schema.Command {

	LogInfoMessage("Create a composite command :")
	command := schema.Command{}
	command.Id = GetRandomUniqueString(8, true)
	LogInfoMessage(fmt.Sprintf("command Id: %s", command.Id))
	command.Composite = &schema.CompositeCommand{}
	devfile.commandAdded(command)

	return &command
}

// SetCompositeCommandValues randomly sets composite command attribute to random values
func (devfile *TestDevfile) SetCompositeCommandValues(command *schema.Command) {

	compositeCommand := command.Composite
	numCommands := GetRandomNumber(1, 3)

	for i := 0; i < numCommands; i++ {
		execCommand := devfile.AddCommand(schema.ExecCommandType)
		compositeCommand.Commands = append(compositeCommand.Commands, execCommand.Id)
		LogInfoMessage(fmt.Sprintf("....... command %d of %d : %s", i, numCommands, execCommand.Id))
	}

	// If group already exists - leave it to make sure defaults are not deleted or added
	if compositeCommand.Group == nil {
		if GetRandomDecision(2, 1) {
			compositeCommand.Group = devfile.addGroup()
		}
	}

	if GetBinaryDecision() {
		compositeCommand.Label = GetRandomString(12, false)
		LogInfoMessage(fmt.Sprintf("....... label: %s", compositeCommand.Label))
	}

	if GetBinaryDecision() {
		compositeCommand.Parallel = true
		LogInfoMessage(fmt.Sprintf("....... Parallel: %t", compositeCommand.Parallel))
	}

	devfile.commandUpdated(*command)
}

// createApplyCommand creates an apply command in a schema structure
func (devfile *TestDevfile) createApplyCommand() *schema.Command {

	LogInfoMessage("Create a apply command :")
	command := schema.Command{}
	command.Id = GetRandomUniqueString(8, true)
	LogInfoMessage(fmt.Sprintf("command Id: %s", command.Id))
	command.Apply = &schema.ApplyCommand{}
	devfile.commandAdded(command)
	return &command
}

// SetApplyCommandValues randomly sets apply command attributes to random values
func (devfile *TestDevfile) SetApplyCommandValues(command *schema.Command) {
	applyCommand := command.Apply

	applyCommand.Component = devfile.GetContainerName()

	if GetRandomDecision(2, 1) {
		applyCommand.Group = devfile.addGroup()
	}

	if GetBinaryDecision() {
		applyCommand.Label = GetRandomString(63, false)
		LogInfoMessage(fmt.Sprintf("....... label: %s", applyCommand.Label))
	}

	devfile.commandUpdated(*command)
}
