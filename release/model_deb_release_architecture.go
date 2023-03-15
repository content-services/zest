/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DebReleaseArchitecture type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DebReleaseArchitecture{}

// DebReleaseArchitecture A Serializer for ReleaseArchitecture.
type DebReleaseArchitecture struct {
	// Name of the architecture.
	Architecture string `json:"architecture"`
	// Release this architecture is contained in.
	Release string `json:"release"`
}

// NewDebReleaseArchitecture instantiates a new DebReleaseArchitecture object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDebReleaseArchitecture(architecture string, release string) *DebReleaseArchitecture {
	this := DebReleaseArchitecture{}
	this.Architecture = architecture
	this.Release = release
	return &this
}

// NewDebReleaseArchitectureWithDefaults instantiates a new DebReleaseArchitecture object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDebReleaseArchitectureWithDefaults() *DebReleaseArchitecture {
	this := DebReleaseArchitecture{}
	return &this
}

// GetArchitecture returns the Architecture field value
func (o *DebReleaseArchitecture) GetArchitecture() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Architecture
}

// GetArchitectureOk returns a tuple with the Architecture field value
// and a boolean to check if the value has been set.
func (o *DebReleaseArchitecture) GetArchitectureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Architecture, true
}

// SetArchitecture sets field value
func (o *DebReleaseArchitecture) SetArchitecture(v string) {
	o.Architecture = v
}

// GetRelease returns the Release field value
func (o *DebReleaseArchitecture) GetRelease() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Release
}

// GetReleaseOk returns a tuple with the Release field value
// and a boolean to check if the value has been set.
func (o *DebReleaseArchitecture) GetReleaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Release, true
}

// SetRelease sets field value
func (o *DebReleaseArchitecture) SetRelease(v string) {
	o.Release = v
}

func (o DebReleaseArchitecture) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DebReleaseArchitecture) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["architecture"] = o.Architecture
	toSerialize["release"] = o.Release
	return toSerialize, nil
}

type NullableDebReleaseArchitecture struct {
	value *DebReleaseArchitecture
	isSet bool
}

func (v NullableDebReleaseArchitecture) Get() *DebReleaseArchitecture {
	return v.value
}

func (v *NullableDebReleaseArchitecture) Set(val *DebReleaseArchitecture) {
	v.value = val
	v.isSet = true
}

func (v NullableDebReleaseArchitecture) IsSet() bool {
	return v.isSet
}

func (v *NullableDebReleaseArchitecture) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDebReleaseArchitecture(val *DebReleaseArchitecture) *NullableDebReleaseArchitecture {
	return &NullableDebReleaseArchitecture{value: val, isSet: true}
}

func (v NullableDebReleaseArchitecture) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDebReleaseArchitecture) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


