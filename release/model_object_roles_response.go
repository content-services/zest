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

// checks if the ObjectRolesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectRolesResponse{}

// ObjectRolesResponse struct for ObjectRolesResponse
type ObjectRolesResponse struct {
	Roles []NestedRoleResponse `json:"roles"`
}

// NewObjectRolesResponse instantiates a new ObjectRolesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectRolesResponse(roles []NestedRoleResponse) *ObjectRolesResponse {
	this := ObjectRolesResponse{}
	this.Roles = roles
	return &this
}

// NewObjectRolesResponseWithDefaults instantiates a new ObjectRolesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectRolesResponseWithDefaults() *ObjectRolesResponse {
	this := ObjectRolesResponse{}
	return &this
}

// GetRoles returns the Roles field value
func (o *ObjectRolesResponse) GetRoles() []NestedRoleResponse {
	if o == nil {
		var ret []NestedRoleResponse
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *ObjectRolesResponse) GetRolesOk() ([]NestedRoleResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *ObjectRolesResponse) SetRoles(v []NestedRoleResponse) {
	o.Roles = v
}

func (o ObjectRolesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectRolesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["roles"] = o.Roles
	return toSerialize, nil
}

type NullableObjectRolesResponse struct {
	value *ObjectRolesResponse
	isSet bool
}

func (v NullableObjectRolesResponse) Get() *ObjectRolesResponse {
	return v.value
}

func (v *NullableObjectRolesResponse) Set(val *ObjectRolesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectRolesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectRolesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectRolesResponse(val *ObjectRolesResponse) *NullableObjectRolesResponse {
	return &NullableObjectRolesResponse{value: val, isSet: true}
}

func (v NullableObjectRolesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectRolesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


