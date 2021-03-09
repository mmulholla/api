package common

import (
	"fmt"
	"math/rand"

	header "github.com/devfile/api/v2/pkg/devfile"
	attributes "github.com/devfile/api/v2/pkg/attributes"


)

// metadataAdded adds new metadata to the test schema data and notifies the follower
func (devfile *TestDevfile) metadataAdded(metadata header.DevfileMetadata) {
	LogInfoMessage("metaData added")
	devfile.SchemaDevFile.Metadata = metadata
	if devfile.Follower != nil {
		devfile.Follower.SetMetaData(metadata)
	}
}

// metadataUpdated notifies the follower of the metadata which has been updated
func (devfile *TestDevfile) metadataUpdated(metadata header.DevfileMetadata) {
	LogInfoMessage("metaData updated")
	if devfile.Follower != nil {
		devfile.Follower.UpdateMetaData(metadata)
	}
}

// AddMetadata creates metadata for addition to the devfile schema
func (devfile *TestDevfile) AddMetadata() header.DevfileMetadata {
	metadata := header.DevfileMetadata{}
	devfile.metadataAdded(metadata)
	devfile.SetMetadataValues(&devfile.SchemaDevFile.Metadata)
	return metadata
}

// SetMetadataValues randomly sets metadata attributes in the supplied metadata
func (devfile *TestDevfile) SetMetadataValues(metadata *header.DevfileMetadata)  {

	if GetRandomDecision(2, 1) {
		attrs := attributes.Attributes{}

		attrs.PutString(GetRandomString(8,false),GetRandomString(8,false))

		if GetBinaryDecision() {
			var arrayValues []string
			arraySize := GetRandomNumber(2,5)
			for i := 0; i < arraySize; i++ {
				arrayValues = append(arrayValues,GetRandomString(GetRandomNumber(4,24),false))
			}
			attrs.Put(GetRandomString(GetRandomNumber(4,24),false), arrayValues, nil)
		}

		if GetBinaryDecision() {
			mapValues := make(map[string]string)
			mapValues[GetRandomString(GetRandomNumber(4,24),false)] = GetRandomString(GetRandomNumber(4,24),false)
			mapValues[GetRandomString(GetRandomNumber(4,24),false)] = GetRandomString(GetRandomNumber(4,24),false)
			attrs.Put(GetRandomString(GetRandomNumber(4,24),false),mapValues,nil)
		}

		if GetBinaryDecision() {
			attrs.PutInteger(GetRandomString(GetRandomNumber(4,24),false),GetRandomNumber(1,9999))
			attrs.PutBoolean(GetRandomString(GetRandomNumber(4,24),false),GetBinaryDecision())
			attrs.PutFloat(GetRandomString(GetRandomNumber(4,24),false),rand.Float64())
		}
		metadata.Attributes = attrs
	}


	if GetRandomDecision(2, 1) {
		numWords := GetRandomNumber(2, 6)
		for i := 0; i < numWords; i++ {
			if i > 0 {
				metadata.Description += " "
			}
			metadata.Description += GetRandomString(8, false)
		}
		LogInfoMessage(fmt.Sprintf("Set metadata description : %s", metadata.Description))
	}

	if GetRandomDecision(2, 1) {
		metadata.DisplayName = GetRandomString(GetRandomNumber(8,12),false)
		LogInfoMessage(fmt.Sprintf("Set metadata display name : %s",metadata.DisplayName))
	}

	if GetRandomDecision(2, 1) {
		metadata.GlobalMemoryLimit = fmt.Sprintf("%dG",GetRandomNumber(64, 256))
		LogInfoMessage(fmt.Sprintf("Set metadata global memory limit name : %s",metadata.GlobalMemoryLimit))
	}

	if GetRandomDecision(2, 1) {
		metadata.Icon = fmt.Sprintf("%s.%s",GetRandomString(GetRandomNumber(4, 32), false),".ico")
		LogInfoMessage(fmt.Sprintf("Set metadata icon : %s",metadata.Icon))
	}

	if GetRandomDecision(2, 1) {
		metadata.Name = GetRandomString(GetRandomNumber(8,12),false)
		LogInfoMessage(fmt.Sprintf("Set metadata name : %s",metadata.Name))
	}

	if GetRandomDecision(2, 1) {
		var tags []string
		NumTags := GetRandomNumber(1,5)
		for i := 0; i < NumTags; i++ {
			tags = append(tags,GetRandomString(GetRandomNumber(4,12),false))
			LogInfoMessage(fmt.Sprintf("Add metadata tag : %s",tags[i]))
		}
		metadata.Tags = tags
	}

	if GetRandomDecision(2, 1) {

		version := fmt.Sprintf("%d.%d.%d",GetRandomNumber(1,10),GetRandomNumber(1,10),GetRandomNumber(1,10))
		if GetBinaryDecision() {
			version = fmt.Sprintf("%s-%s", version, GetRandomString(5, true))
		}
		metadata.Version = version
		LogInfoMessage(fmt.Sprintf("Set metadata version : %s",metadata.Version))
	}

	devfile.metadataUpdated(*metadata)
}

