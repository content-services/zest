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

// checks if the RpmPackage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RpmPackage{}

// RpmPackage A Serializer for Package.  Add serializers for the new fields defined in Package and add those fields to the Meta class keeping fields from the parent class as well. Provide help_text.
type RpmPackage struct {
	// Artifact file representing the physical content
	Artifact *string `json:"artifact,omitempty"`
	// Path where the artifact is located relative to distributions base_path
	RelativePath *string `json:"relative_path,omitempty"`
	// An uploaded file that may be turned into the artifact of the content unit.
	File **os.File `json:"file,omitempty"`
	// A URI of a repository the new content unit should be associated with.
	Repository *string `json:"repository,omitempty"`
	// An uncommitted upload that may be turned into the artifact of the content unit.
	Upload *string `json:"upload,omitempty"`
}

// NewRpmPackage instantiates a new RpmPackage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRpmPackage() *RpmPackage {
	this := RpmPackage{}
	return &this
}

// NewRpmPackageWithDefaults instantiates a new RpmPackage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRpmPackageWithDefaults() *RpmPackage {
	this := RpmPackage{}
	return &this
}

// GetArtifact returns the Artifact field value if set, zero value otherwise.
func (o *RpmPackage) GetArtifact() string {
	if o == nil || IsNil(o.Artifact) {
		var ret string
		return ret
	}
	return *o.Artifact
}

// GetArtifactOk returns a tuple with the Artifact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmPackage) GetArtifactOk() (*string, bool) {
	if o == nil || IsNil(o.Artifact) {
		return nil, false
	}
	return o.Artifact, true
}

// HasArtifact returns a boolean if a field has been set.
func (o *RpmPackage) HasArtifact() bool {
	if o != nil && !IsNil(o.Artifact) {
		return true
	}

	return false
}

// SetArtifact gets a reference to the given string and assigns it to the Artifact field.
func (o *RpmPackage) SetArtifact(v string) {
	o.Artifact = &v
}

// GetRelativePath returns the RelativePath field value if set, zero value otherwise.
func (o *RpmPackage) GetRelativePath() string {
	if o == nil || IsNil(o.RelativePath) {
		var ret string
		return ret
	}
	return *o.RelativePath
}

// GetRelativePathOk returns a tuple with the RelativePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmPackage) GetRelativePathOk() (*string, bool) {
	if o == nil || IsNil(o.RelativePath) {
		return nil, false
	}
	return o.RelativePath, true
}

// HasRelativePath returns a boolean if a field has been set.
func (o *RpmPackage) HasRelativePath() bool {
	if o != nil && !IsNil(o.RelativePath) {
		return true
	}

	return false
}

// SetRelativePath gets a reference to the given string and assigns it to the RelativePath field.
func (o *RpmPackage) SetRelativePath(v string) {
	o.RelativePath = &v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *RpmPackage) GetFile() *os.File {
	if o == nil || IsNil(o.File) {
		var ret *os.File
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmPackage) GetFileOk() (**os.File, bool) {
	if o == nil || IsNil(o.File) {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *RpmPackage) HasFile() bool {
	if o != nil && !IsNil(o.File) {
		return true
	}

	return false
}

// SetFile gets a reference to the given *os.File and assigns it to the File field.
func (o *RpmPackage) SetFile(v *os.File) {
	o.File = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *RpmPackage) GetRepository() string {
	if o == nil || IsNil(o.Repository) {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmPackage) GetRepositoryOk() (*string, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *RpmPackage) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *RpmPackage) SetRepository(v string) {
	o.Repository = &v
}

// GetUpload returns the Upload field value if set, zero value otherwise.
func (o *RpmPackage) GetUpload() string {
	if o == nil || IsNil(o.Upload) {
		var ret string
		return ret
	}
	return *o.Upload
}

// GetUploadOk returns a tuple with the Upload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmPackage) GetUploadOk() (*string, bool) {
	if o == nil || IsNil(o.Upload) {
		return nil, false
	}
	return o.Upload, true
}

// HasUpload returns a boolean if a field has been set.
func (o *RpmPackage) HasUpload() bool {
	if o != nil && !IsNil(o.Upload) {
		return true
	}

	return false
}

// SetUpload gets a reference to the given string and assigns it to the Upload field.
func (o *RpmPackage) SetUpload(v string) {
	o.Upload = &v
}

func (o RpmPackage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RpmPackage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Artifact) {
		toSerialize["artifact"] = o.Artifact
	}
	if !IsNil(o.RelativePath) {
		toSerialize["relative_path"] = o.RelativePath
	}
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

type NullableRpmPackage struct {
	value *RpmPackage
	isSet bool
}

func (v NullableRpmPackage) Get() *RpmPackage {
	return v.value
}

func (v *NullableRpmPackage) Set(val *RpmPackage) {
	v.value = val
	v.isSet = true
}

func (v NullableRpmPackage) IsSet() bool {
	return v.isSet
}

func (v *NullableRpmPackage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRpmPackage(val *RpmPackage) *NullableRpmPackage {
	return &NullableRpmPackage{value: val, isSet: true}
}

func (v NullableRpmPackage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRpmPackage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


