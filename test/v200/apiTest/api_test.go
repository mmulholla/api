package apiTest

import (
	"testing"

	schema "github.com/devfile/api/v2/pkg/apis/workspaces/v1alpha2"
	api "github.com/devfile/api/v2/test/v200/utils/api"
	common "github.com/devfile/api/v2/test/v200/utils/common"
)

func Test_ExecCommand(t *testing.T) {
	testContent := common.TestContent{}
	testContent.CommandTypes = []schema.CommandType{schema.ExecCommandType}
	testContent.EditContent = false
	testContent.FileName = common.GetDevFileName()
	api.RunTest(testContent, t)
}

func Test_ApplyCommand(t *testing.T) {
	testContent := common.TestContent{}
	testContent.CommandTypes = []schema.CommandType{schema.ApplyCommandType}
	testContent.EditContent = false
	testContent.FileName = common.GetDevFileName()
	api.RunTest(testContent, t)
}

func Test_CompositeCommand(t *testing.T) {
	testContent := common.TestContent{}
	testContent.CommandTypes = []schema.CommandType{schema.CompositeCommandType}
	testContent.EditContent = false
	testContent.FileName = common.GetDevFileName()
	api.RunTest(testContent, t)
}

func Test_MultiCommand(t *testing.T) {
	testContent := common.TestContent{}
	testContent.CommandTypes = []schema.CommandType{schema.ExecCommandType,
		schema.CompositeCommandType,
		schema.ApplyCommandType}
	testContent.EditContent = true
	testContent.FileName = common.GetDevFileName()
	api.RunTest(testContent, t)
}

func Test_ContainerComponent(t *testing.T) {
	testContent := common.TestContent{}
	testContent.ComponentTypes = []schema.ComponentType{schema.ContainerComponentType}
	testContent.EditContent = false
	testContent.FileName = common.GetDevFileName()
	api.RunTest(testContent, t)
}

func Test_VolumeComponent(t *testing.T) {
	testContent := common.TestContent{}
	testContent.ComponentTypes = []schema.ComponentType{schema.VolumeComponentType}
	testContent.FileName = common.GetDevFileName()
	api.RunTest(testContent, t)
}

func Test_MultiComponent(t *testing.T) {
	testContent := common.TestContent{}
	testContent.ComponentTypes = []schema.ComponentType{
		schema.ContainerComponentType,
		schema.VolumeComponentType}
	testContent.FileName = common.GetDevFileName()
	api.RunTest(testContent, t)
}

func Test_Projects(t *testing.T) {
	testContent := common.TestContent{}
	testContent.ProjectTypes = []schema.ProjectSourceType{
		schema.GitProjectSourceType,
		schema.GitHubProjectSourceType,
		schema.ZipProjectSourceType}
	testContent.FileName = common.GetDevFileName()
	api.RunTest(testContent, t)
}

func Test_StarterProjects(t *testing.T) {
	testContent := common.TestContent{}
	testContent.ProjectTypes = []schema.ProjectSourceType{
		schema.GitProjectSourceType,
		schema.GitHubProjectSourceType,
		schema.ZipProjectSourceType}
	testContent.FileName = common.GetDevFileName()
	api.RunTest(testContent, t)
}

func Test_Everything(t *testing.T) {
	testContent := common.TestContent{}
	testContent.CommandTypes = []schema.CommandType{
		schema.ExecCommandType,
		schema.CompositeCommandType,
		schema.ApplyCommandType}
	testContent.ComponentTypes = []schema.ComponentType{
		schema.ContainerComponentType,
		schema.VolumeComponentType}
	testContent.ProjectTypes = []schema.ProjectSourceType{
		schema.GitProjectSourceType,
		schema.GitHubProjectSourceType,
		schema.ZipProjectSourceType}
	testContent.ProjectTypes = []schema.ProjectSourceType{
		schema.GitProjectSourceType,
		schema.GitHubProjectSourceType,
		schema.ZipProjectSourceType}
	testContent.FileName = common.GetDevFileName()
	api.RunTest(testContent, t)
}
