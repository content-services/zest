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
	"fmt"
)

// checks if the MavenMavenArtifact type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MavenMavenArtifact{}

// MavenMavenArtifact A Serializer for MavenArtifact.
type MavenMavenArtifact struct {
	// A URI of a repository the new content unit should be associated with.
	Repository *string `json:"repository,omitempty"`
	// A dictionary of arbitrary key/value pairs used to describe a specific Content instance.
	PulpLabels *map[string]string `json:"pulp_labels,omitempty"`
	// Artifact file representing the physical content
	Artifact string `json:"artifact"`
	// Path where the artifact is located relative to distributions base_path
	RelativePath string `json:"relative_path"`
	AdditionalProperties map[string]interface{}
}

type _MavenMavenArtifact MavenMavenArtifact

// NewMavenMavenArtifact instantiates a new MavenMavenArtifact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMavenMavenArtifact(artifact string, relativePath string) *MavenMavenArtifact {
	this := MavenMavenArtifact{}
	this.Artifact = artifact
	this.RelativePath = relativePath
	return &this
}

// NewMavenMavenArtifactWithDefaults instantiates a new MavenMavenArtifact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMavenMavenArtifactWithDefaults() *MavenMavenArtifact {
	this := MavenMavenArtifact{}
	return &this
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *MavenMavenArtifact) GetRepository() string {
	if o == nil || IsNil(o.Repository) {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MavenMavenArtifact) GetRepositoryOk() (*string, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *MavenMavenArtifact) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *MavenMavenArtifact) SetRepository(v string) {
	o.Repository = &v
}

// GetPulpLabels returns the PulpLabels field value if set, zero value otherwise.
func (o *MavenMavenArtifact) GetPulpLabels() map[string]string {
	if o == nil || IsNil(o.PulpLabels) {
		var ret map[string]string
		return ret
	}
	return *o.PulpLabels
}

// GetPulpLabelsOk returns a tuple with the PulpLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MavenMavenArtifact) GetPulpLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PulpLabels) {
		return nil, false
	}
	return o.PulpLabels, true
}

// HasPulpLabels returns a boolean if a field has been set.
func (o *MavenMavenArtifact) HasPulpLabels() bool {
	if o != nil && !IsNil(o.PulpLabels) {
		return true
	}

	return false
}

// SetPulpLabels gets a reference to the given map[string]string and assigns it to the PulpLabels field.
func (o *MavenMavenArtifact) SetPulpLabels(v map[string]string) {
	o.PulpLabels = &v
}

// GetArtifact returns the Artifact field value
func (o *MavenMavenArtifact) GetArtifact() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Artifact
}

// GetArtifactOk returns a tuple with the Artifact field value
// and a boolean to check if the value has been set.
func (o *MavenMavenArtifact) GetArtifactOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Artifact, true
}

// SetArtifact sets field value
func (o *MavenMavenArtifact) SetArtifact(v string) {
	o.Artifact = v
}

// GetRelativePath returns the RelativePath field value
func (o *MavenMavenArtifact) GetRelativePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RelativePath
}

// GetRelativePathOk returns a tuple with the RelativePath field value
// and a boolean to check if the value has been set.
func (o *MavenMavenArtifact) GetRelativePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelativePath, true
}

// SetRelativePath sets field value
func (o *MavenMavenArtifact) SetRelativePath(v string) {
	o.RelativePath = v
}

func (o MavenMavenArtifact) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MavenMavenArtifact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}
	if !IsNil(o.PulpLabels) {
		toSerialize["pulp_labels"] = o.PulpLabels
	}
	toSerialize["artifact"] = o.Artifact
	toSerialize["relative_path"] = o.RelativePath

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MavenMavenArtifact) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"artifact",
		"relative_path",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varMavenMavenArtifact := _MavenMavenArtifact{}

	err = json.Unmarshal(data, &varMavenMavenArtifact)

	if err != nil {
		return err
	}

	*o = MavenMavenArtifact(varMavenMavenArtifact)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "repository")
		delete(additionalProperties, "pulp_labels")
		delete(additionalProperties, "artifact")
		delete(additionalProperties, "relative_path")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMavenMavenArtifact struct {
	value *MavenMavenArtifact
	isSet bool
}

func (v NullableMavenMavenArtifact) Get() *MavenMavenArtifact {
	return v.value
}

func (v *NullableMavenMavenArtifact) Set(val *MavenMavenArtifact) {
	v.value = val
	v.isSet = true
}

func (v NullableMavenMavenArtifact) IsSet() bool {
	return v.isSet
}

func (v *NullableMavenMavenArtifact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMavenMavenArtifact(val *MavenMavenArtifact) *NullableMavenMavenArtifact {
	return &NullableMavenMavenArtifact{value: val, isSet: true}
}

func (v NullableMavenMavenArtifact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMavenMavenArtifact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


