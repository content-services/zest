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

// checks if the PulpImportCheckResponseToc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PulpImportCheckResponseToc{}

// PulpImportCheckResponseToc Evaluation of proposed 'toc' file for PulpImport
type PulpImportCheckResponseToc struct {
	// Parameter value being evaluated.
	Context string `json:"context"`
	// True if evaluation passed, false otherwise.
	IsValid bool `json:"is_valid"`
	// Messages describing results of all evaluations done. May be an empty list.
	Messages []string `json:"messages"`
}

// NewPulpImportCheckResponseToc instantiates a new PulpImportCheckResponseToc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPulpImportCheckResponseToc(context string, isValid bool, messages []string) *PulpImportCheckResponseToc {
	this := PulpImportCheckResponseToc{}
	this.Context = context
	this.IsValid = isValid
	this.Messages = messages
	return &this
}

// NewPulpImportCheckResponseTocWithDefaults instantiates a new PulpImportCheckResponseToc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPulpImportCheckResponseTocWithDefaults() *PulpImportCheckResponseToc {
	this := PulpImportCheckResponseToc{}
	return &this
}

// GetContext returns the Context field value
func (o *PulpImportCheckResponseToc) GetContext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *PulpImportCheckResponseToc) GetContextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *PulpImportCheckResponseToc) SetContext(v string) {
	o.Context = v
}

// GetIsValid returns the IsValid field value
func (o *PulpImportCheckResponseToc) GetIsValid() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsValid
}

// GetIsValidOk returns a tuple with the IsValid field value
// and a boolean to check if the value has been set.
func (o *PulpImportCheckResponseToc) GetIsValidOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsValid, true
}

// SetIsValid sets field value
func (o *PulpImportCheckResponseToc) SetIsValid(v bool) {
	o.IsValid = v
}

// GetMessages returns the Messages field value
func (o *PulpImportCheckResponseToc) GetMessages() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
func (o *PulpImportCheckResponseToc) GetMessagesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Messages, true
}

// SetMessages sets field value
func (o *PulpImportCheckResponseToc) SetMessages(v []string) {
	o.Messages = v
}

func (o PulpImportCheckResponseToc) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PulpImportCheckResponseToc) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["context"] = o.Context
	toSerialize["is_valid"] = o.IsValid
	toSerialize["messages"] = o.Messages
	return toSerialize, nil
}

type NullablePulpImportCheckResponseToc struct {
	value *PulpImportCheckResponseToc
	isSet bool
}

func (v NullablePulpImportCheckResponseToc) Get() *PulpImportCheckResponseToc {
	return v.value
}

func (v *NullablePulpImportCheckResponseToc) Set(val *PulpImportCheckResponseToc) {
	v.value = val
	v.isSet = true
}

func (v NullablePulpImportCheckResponseToc) IsSet() bool {
	return v.isSet
}

func (v *NullablePulpImportCheckResponseToc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePulpImportCheckResponseToc(val *PulpImportCheckResponseToc) *NullablePulpImportCheckResponseToc {
	return &NullablePulpImportCheckResponseToc{value: val, isSet: true}
}

func (v NullablePulpImportCheckResponseToc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePulpImportCheckResponseToc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


