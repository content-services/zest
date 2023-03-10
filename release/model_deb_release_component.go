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

// checks if the DebReleaseComponent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DebReleaseComponent{}

// DebReleaseComponent A Serializer for ReleaseComponent.
type DebReleaseComponent struct {
	// Name of the component.
	Component string `json:"component"`
	// Release this component is contained in.
	Release string `json:"release"`
}

// NewDebReleaseComponent instantiates a new DebReleaseComponent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDebReleaseComponent(component string, release string) *DebReleaseComponent {
	this := DebReleaseComponent{}
	this.Component = component
	this.Release = release
	return &this
}

// NewDebReleaseComponentWithDefaults instantiates a new DebReleaseComponent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDebReleaseComponentWithDefaults() *DebReleaseComponent {
	this := DebReleaseComponent{}
	return &this
}

// GetComponent returns the Component field value
func (o *DebReleaseComponent) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *DebReleaseComponent) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Component, true
}

// SetComponent sets field value
func (o *DebReleaseComponent) SetComponent(v string) {
	o.Component = v
}

// GetRelease returns the Release field value
func (o *DebReleaseComponent) GetRelease() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Release
}

// GetReleaseOk returns a tuple with the Release field value
// and a boolean to check if the value has been set.
func (o *DebReleaseComponent) GetReleaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Release, true
}

// SetRelease sets field value
func (o *DebReleaseComponent) SetRelease(v string) {
	o.Release = v
}

func (o DebReleaseComponent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DebReleaseComponent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["component"] = o.Component
	toSerialize["release"] = o.Release
	return toSerialize, nil
}

type NullableDebReleaseComponent struct {
	value *DebReleaseComponent
	isSet bool
}

func (v NullableDebReleaseComponent) Get() *DebReleaseComponent {
	return v.value
}

func (v *NullableDebReleaseComponent) Set(val *DebReleaseComponent) {
	v.value = val
	v.isSet = true
}

func (v NullableDebReleaseComponent) IsSet() bool {
	return v.isSet
}

func (v *NullableDebReleaseComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDebReleaseComponent(val *DebReleaseComponent) *NullableDebReleaseComponent {
	return &NullableDebReleaseComponent{value: val, isSet: true}
}

func (v NullableDebReleaseComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDebReleaseComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


