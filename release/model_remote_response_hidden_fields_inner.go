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

// checks if the RemoteResponseHiddenFieldsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoteResponseHiddenFieldsInner{}

// RemoteResponseHiddenFieldsInner struct for RemoteResponseHiddenFieldsInner
type RemoteResponseHiddenFieldsInner struct {
	Name *string `json:"name,omitempty"`
	IsSet *bool `json:"is_set,omitempty"`
}

// NewRemoteResponseHiddenFieldsInner instantiates a new RemoteResponseHiddenFieldsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteResponseHiddenFieldsInner() *RemoteResponseHiddenFieldsInner {
	this := RemoteResponseHiddenFieldsInner{}
	return &this
}

// NewRemoteResponseHiddenFieldsInnerWithDefaults instantiates a new RemoteResponseHiddenFieldsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteResponseHiddenFieldsInnerWithDefaults() *RemoteResponseHiddenFieldsInner {
	this := RemoteResponseHiddenFieldsInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RemoteResponseHiddenFieldsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteResponseHiddenFieldsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RemoteResponseHiddenFieldsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RemoteResponseHiddenFieldsInner) SetName(v string) {
	o.Name = &v
}

// GetIsSet returns the IsSet field value if set, zero value otherwise.
func (o *RemoteResponseHiddenFieldsInner) GetIsSet() bool {
	if o == nil || IsNil(o.IsSet) {
		var ret bool
		return ret
	}
	return *o.IsSet
}

// GetIsSetOk returns a tuple with the IsSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteResponseHiddenFieldsInner) GetIsSetOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSet) {
		return nil, false
	}
	return o.IsSet, true
}

// HasIsSet returns a boolean if a field has been set.
func (o *RemoteResponseHiddenFieldsInner) HasIsSet() bool {
	if o != nil && !IsNil(o.IsSet) {
		return true
	}

	return false
}

// SetIsSet gets a reference to the given bool and assigns it to the IsSet field.
func (o *RemoteResponseHiddenFieldsInner) SetIsSet(v bool) {
	o.IsSet = &v
}

func (o RemoteResponseHiddenFieldsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoteResponseHiddenFieldsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.IsSet) {
		toSerialize["is_set"] = o.IsSet
	}
	return toSerialize, nil
}

type NullableRemoteResponseHiddenFieldsInner struct {
	value *RemoteResponseHiddenFieldsInner
	isSet bool
}

func (v NullableRemoteResponseHiddenFieldsInner) Get() *RemoteResponseHiddenFieldsInner {
	return v.value
}

func (v *NullableRemoteResponseHiddenFieldsInner) Set(val *RemoteResponseHiddenFieldsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteResponseHiddenFieldsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteResponseHiddenFieldsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteResponseHiddenFieldsInner(val *RemoteResponseHiddenFieldsInner) *NullableRemoteResponseHiddenFieldsInner {
	return &NullableRemoteResponseHiddenFieldsInner{value: val, isSet: true}
}

func (v NullableRemoteResponseHiddenFieldsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteResponseHiddenFieldsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


