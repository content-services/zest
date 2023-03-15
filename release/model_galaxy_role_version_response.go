/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package zest/release/v3

import (
	"encoding/json"
)

// checks if the GalaxyRoleVersionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GalaxyRoleVersionResponse{}

// GalaxyRoleVersionResponse A serializer for Galaxy's representation of Role versions.
type GalaxyRoleVersionResponse struct {
	Name string `json:"name"`
	Source *string `json:"source,omitempty"`
}

// NewGalaxyRoleVersionResponse instantiates a new GalaxyRoleVersionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGalaxyRoleVersionResponse(name string) *GalaxyRoleVersionResponse {
	this := GalaxyRoleVersionResponse{}
	this.Name = name
	return &this
}

// NewGalaxyRoleVersionResponseWithDefaults instantiates a new GalaxyRoleVersionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGalaxyRoleVersionResponseWithDefaults() *GalaxyRoleVersionResponse {
	this := GalaxyRoleVersionResponse{}
	return &this
}

// GetName returns the Name field value
func (o *GalaxyRoleVersionResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GalaxyRoleVersionResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GalaxyRoleVersionResponse) SetName(v string) {
	o.Name = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *GalaxyRoleVersionResponse) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GalaxyRoleVersionResponse) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *GalaxyRoleVersionResponse) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *GalaxyRoleVersionResponse) SetSource(v string) {
	o.Source = &v
}

func (o GalaxyRoleVersionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GalaxyRoleVersionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	// skip: source is readOnly
	return toSerialize, nil
}

type NullableGalaxyRoleVersionResponse struct {
	value *GalaxyRoleVersionResponse
	isSet bool
}

func (v NullableGalaxyRoleVersionResponse) Get() *GalaxyRoleVersionResponse {
	return v.value
}

func (v *NullableGalaxyRoleVersionResponse) Set(val *GalaxyRoleVersionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGalaxyRoleVersionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGalaxyRoleVersionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGalaxyRoleVersionResponse(val *GalaxyRoleVersionResponse) *NullableGalaxyRoleVersionResponse {
	return &NullableGalaxyRoleVersionResponse{value: val, isSet: true}
}

func (v NullableGalaxyRoleVersionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGalaxyRoleVersionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


