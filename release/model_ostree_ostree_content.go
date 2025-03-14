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

// checks if the OstreeOstreeContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OstreeOstreeContent{}

// OstreeOstreeContent A Serializer class for uncategorized content units (e.g., static deltas).
type OstreeOstreeContent struct {
	// A URI of a repository the new content unit should be associated with.
	Repository *string `json:"repository,omitempty"`
	// A dictionary of arbitrary key/value pairs used to describe a specific Content instance.
	PulpLabels *map[string]string `json:"pulp_labels,omitempty"`
	// Artifact file representing the physical content
	Artifact string `json:"artifact"`
	RelativePath string `json:"relative_path"`
	Digest string `json:"digest"`
	AdditionalProperties map[string]interface{}
}

type _OstreeOstreeContent OstreeOstreeContent

// NewOstreeOstreeContent instantiates a new OstreeOstreeContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOstreeOstreeContent(artifact string, relativePath string, digest string) *OstreeOstreeContent {
	this := OstreeOstreeContent{}
	this.Artifact = artifact
	this.RelativePath = relativePath
	this.Digest = digest
	return &this
}

// NewOstreeOstreeContentWithDefaults instantiates a new OstreeOstreeContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOstreeOstreeContentWithDefaults() *OstreeOstreeContent {
	this := OstreeOstreeContent{}
	return &this
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *OstreeOstreeContent) GetRepository() string {
	if o == nil || IsNil(o.Repository) {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OstreeOstreeContent) GetRepositoryOk() (*string, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *OstreeOstreeContent) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *OstreeOstreeContent) SetRepository(v string) {
	o.Repository = &v
}

// GetPulpLabels returns the PulpLabels field value if set, zero value otherwise.
func (o *OstreeOstreeContent) GetPulpLabels() map[string]string {
	if o == nil || IsNil(o.PulpLabels) {
		var ret map[string]string
		return ret
	}
	return *o.PulpLabels
}

// GetPulpLabelsOk returns a tuple with the PulpLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OstreeOstreeContent) GetPulpLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PulpLabels) {
		return nil, false
	}
	return o.PulpLabels, true
}

// HasPulpLabels returns a boolean if a field has been set.
func (o *OstreeOstreeContent) HasPulpLabels() bool {
	if o != nil && !IsNil(o.PulpLabels) {
		return true
	}

	return false
}

// SetPulpLabels gets a reference to the given map[string]string and assigns it to the PulpLabels field.
func (o *OstreeOstreeContent) SetPulpLabels(v map[string]string) {
	o.PulpLabels = &v
}

// GetArtifact returns the Artifact field value
func (o *OstreeOstreeContent) GetArtifact() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Artifact
}

// GetArtifactOk returns a tuple with the Artifact field value
// and a boolean to check if the value has been set.
func (o *OstreeOstreeContent) GetArtifactOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Artifact, true
}

// SetArtifact sets field value
func (o *OstreeOstreeContent) SetArtifact(v string) {
	o.Artifact = v
}

// GetRelativePath returns the RelativePath field value
func (o *OstreeOstreeContent) GetRelativePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RelativePath
}

// GetRelativePathOk returns a tuple with the RelativePath field value
// and a boolean to check if the value has been set.
func (o *OstreeOstreeContent) GetRelativePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelativePath, true
}

// SetRelativePath sets field value
func (o *OstreeOstreeContent) SetRelativePath(v string) {
	o.RelativePath = v
}

// GetDigest returns the Digest field value
func (o *OstreeOstreeContent) GetDigest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Digest
}

// GetDigestOk returns a tuple with the Digest field value
// and a boolean to check if the value has been set.
func (o *OstreeOstreeContent) GetDigestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Digest, true
}

// SetDigest sets field value
func (o *OstreeOstreeContent) SetDigest(v string) {
	o.Digest = v
}

func (o OstreeOstreeContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OstreeOstreeContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}
	if !IsNil(o.PulpLabels) {
		toSerialize["pulp_labels"] = o.PulpLabels
	}
	toSerialize["artifact"] = o.Artifact
	toSerialize["relative_path"] = o.RelativePath
	toSerialize["digest"] = o.Digest

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OstreeOstreeContent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"artifact",
		"relative_path",
		"digest",
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

	varOstreeOstreeContent := _OstreeOstreeContent{}

	err = json.Unmarshal(data, &varOstreeOstreeContent)

	if err != nil {
		return err
	}

	*o = OstreeOstreeContent(varOstreeOstreeContent)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "repository")
		delete(additionalProperties, "pulp_labels")
		delete(additionalProperties, "artifact")
		delete(additionalProperties, "relative_path")
		delete(additionalProperties, "digest")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOstreeOstreeContent struct {
	value *OstreeOstreeContent
	isSet bool
}

func (v NullableOstreeOstreeContent) Get() *OstreeOstreeContent {
	return v.value
}

func (v *NullableOstreeOstreeContent) Set(val *OstreeOstreeContent) {
	v.value = val
	v.isSet = true
}

func (v NullableOstreeOstreeContent) IsSet() bool {
	return v.isSet
}

func (v *NullableOstreeOstreeContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOstreeOstreeContent(val *OstreeOstreeContent) *NullableOstreeOstreeContent {
	return &NullableOstreeOstreeContent{value: val, isSet: true}
}

func (v NullableOstreeOstreeContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOstreeOstreeContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


