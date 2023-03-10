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
)

// checks if the PulpImportCheckResponseRepoMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PulpImportCheckResponseRepoMapping{}

// PulpImportCheckResponseRepoMapping Evaluation of proposed 'repo_mapping' file for PulpImport
type PulpImportCheckResponseRepoMapping struct {
	// Parameter value being evaluated.
	Context string `json:"context"`
	// True if evaluation passed, false otherwise.
	IsValid bool `json:"is_valid"`
	// Messages describing results of all evaluations done. May be an empty list.
	Messages []string `json:"messages"`
}

// NewPulpImportCheckResponseRepoMapping instantiates a new PulpImportCheckResponseRepoMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPulpImportCheckResponseRepoMapping(context string, isValid bool, messages []string) *PulpImportCheckResponseRepoMapping {
	this := PulpImportCheckResponseRepoMapping{}
	this.Context = context
	this.IsValid = isValid
	this.Messages = messages
	return &this
}

// NewPulpImportCheckResponseRepoMappingWithDefaults instantiates a new PulpImportCheckResponseRepoMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPulpImportCheckResponseRepoMappingWithDefaults() *PulpImportCheckResponseRepoMapping {
	this := PulpImportCheckResponseRepoMapping{}
	return &this
}

// GetContext returns the Context field value
func (o *PulpImportCheckResponseRepoMapping) GetContext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *PulpImportCheckResponseRepoMapping) GetContextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *PulpImportCheckResponseRepoMapping) SetContext(v string) {
	o.Context = v
}

// GetIsValid returns the IsValid field value
func (o *PulpImportCheckResponseRepoMapping) GetIsValid() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsValid
}

// GetIsValidOk returns a tuple with the IsValid field value
// and a boolean to check if the value has been set.
func (o *PulpImportCheckResponseRepoMapping) GetIsValidOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsValid, true
}

// SetIsValid sets field value
func (o *PulpImportCheckResponseRepoMapping) SetIsValid(v bool) {
	o.IsValid = v
}

// GetMessages returns the Messages field value
func (o *PulpImportCheckResponseRepoMapping) GetMessages() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
func (o *PulpImportCheckResponseRepoMapping) GetMessagesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Messages, true
}

// SetMessages sets field value
func (o *PulpImportCheckResponseRepoMapping) SetMessages(v []string) {
	o.Messages = v
}

func (o PulpImportCheckResponseRepoMapping) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PulpImportCheckResponseRepoMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["context"] = o.Context
	toSerialize["is_valid"] = o.IsValid
	toSerialize["messages"] = o.Messages
	return toSerialize, nil
}

type NullablePulpImportCheckResponseRepoMapping struct {
	value *PulpImportCheckResponseRepoMapping
	isSet bool
}

func (v NullablePulpImportCheckResponseRepoMapping) Get() *PulpImportCheckResponseRepoMapping {
	return v.value
}

func (v *NullablePulpImportCheckResponseRepoMapping) Set(val *PulpImportCheckResponseRepoMapping) {
	v.value = val
	v.isSet = true
}

func (v NullablePulpImportCheckResponseRepoMapping) IsSet() bool {
	return v.isSet
}

func (v *NullablePulpImportCheckResponseRepoMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePulpImportCheckResponseRepoMapping(val *PulpImportCheckResponseRepoMapping) *NullablePulpImportCheckResponseRepoMapping {
	return &NullablePulpImportCheckResponseRepoMapping{value: val, isSet: true}
}

func (v NullablePulpImportCheckResponseRepoMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePulpImportCheckResponseRepoMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


