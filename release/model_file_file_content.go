/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package zest

import (
	"encoding/json"
	"os"
)

// checks if the FileFileContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileFileContent{}

// FileFileContent Serializer for File Content.
type FileFileContent struct {
	// Artifact file representing the physical content
	Artifact *string `json:"artifact,omitempty"`
	// Path where the artifact is located relative to distributions base_path
	RelativePath string `json:"relative_path"`
	// An uploaded file that may be turned into the artifact of the content unit.
	File **os.File `json:"file,omitempty"`
	// A URI of a repository the new content unit should be associated with.
	Repository *string `json:"repository,omitempty"`
	// An uncommitted upload that may be turned into the artifact of the content unit.
	Upload *string `json:"upload,omitempty"`
}

// NewFileFileContent instantiates a new FileFileContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileFileContent(relativePath string) *FileFileContent {
	this := FileFileContent{}
	this.RelativePath = relativePath
	return &this
}

// NewFileFileContentWithDefaults instantiates a new FileFileContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileFileContentWithDefaults() *FileFileContent {
	this := FileFileContent{}
	return &this
}

// GetArtifact returns the Artifact field value if set, zero value otherwise.
func (o *FileFileContent) GetArtifact() string {
	if o == nil || IsNil(o.Artifact) {
		var ret string
		return ret
	}
	return *o.Artifact
}

// GetArtifactOk returns a tuple with the Artifact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileFileContent) GetArtifactOk() (*string, bool) {
	if o == nil || IsNil(o.Artifact) {
		return nil, false
	}
	return o.Artifact, true
}

// HasArtifact returns a boolean if a field has been set.
func (o *FileFileContent) HasArtifact() bool {
	if o != nil && !IsNil(o.Artifact) {
		return true
	}

	return false
}

// SetArtifact gets a reference to the given string and assigns it to the Artifact field.
func (o *FileFileContent) SetArtifact(v string) {
	o.Artifact = &v
}

// GetRelativePath returns the RelativePath field value
func (o *FileFileContent) GetRelativePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RelativePath
}

// GetRelativePathOk returns a tuple with the RelativePath field value
// and a boolean to check if the value has been set.
func (o *FileFileContent) GetRelativePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelativePath, true
}

// SetRelativePath sets field value
func (o *FileFileContent) SetRelativePath(v string) {
	o.RelativePath = v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *FileFileContent) GetFile() *os.File {
	if o == nil || IsNil(o.File) {
		var ret *os.File
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileFileContent) GetFileOk() (**os.File, bool) {
	if o == nil || IsNil(o.File) {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *FileFileContent) HasFile() bool {
	if o != nil && !IsNil(o.File) {
		return true
	}

	return false
}

// SetFile gets a reference to the given *os.File and assigns it to the File field.
func (o *FileFileContent) SetFile(v *os.File) {
	o.File = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *FileFileContent) GetRepository() string {
	if o == nil || IsNil(o.Repository) {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileFileContent) GetRepositoryOk() (*string, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *FileFileContent) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *FileFileContent) SetRepository(v string) {
	o.Repository = &v
}

// GetUpload returns the Upload field value if set, zero value otherwise.
func (o *FileFileContent) GetUpload() string {
	if o == nil || IsNil(o.Upload) {
		var ret string
		return ret
	}
	return *o.Upload
}

// GetUploadOk returns a tuple with the Upload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileFileContent) GetUploadOk() (*string, bool) {
	if o == nil || IsNil(o.Upload) {
		return nil, false
	}
	return o.Upload, true
}

// HasUpload returns a boolean if a field has been set.
func (o *FileFileContent) HasUpload() bool {
	if o != nil && !IsNil(o.Upload) {
		return true
	}

	return false
}

// SetUpload gets a reference to the given string and assigns it to the Upload field.
func (o *FileFileContent) SetUpload(v string) {
	o.Upload = &v
}

func (o FileFileContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileFileContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Artifact) {
		toSerialize["artifact"] = o.Artifact
	}
	toSerialize["relative_path"] = o.RelativePath
	if !IsNil(o.File) {
		toSerialize["file"] = o.File
	}
	if !IsNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}
	if !IsNil(o.Upload) {
		toSerialize["upload"] = o.Upload
	}
	return toSerialize, nil
}

type NullableFileFileContent struct {
	value *FileFileContent
	isSet bool
}

func (v NullableFileFileContent) Get() *FileFileContent {
	return v.value
}

func (v *NullableFileFileContent) Set(val *FileFileContent) {
	v.value = val
	v.isSet = true
}

func (v NullableFileFileContent) IsSet() bool {
	return v.isSet
}

func (v *NullableFileFileContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileFileContent(val *FileFileContent) *NullableFileFileContent {
	return &NullableFileFileContent{value: val, isSet: true}
}

func (v NullableFileFileContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileFileContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


