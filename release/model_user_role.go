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

// checks if the UserRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserRole{}

// UserRole Serializer for UserRole.
type UserRole struct {
	Role string `json:"role"`
	// pulp_href of the object for which role permissions should be asserted. If set to 'null', permissions will act on the model-level.
	ContentObject NullableString `json:"content_object"`
}

// NewUserRole instantiates a new UserRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserRole(role string, contentObject NullableString) *UserRole {
	this := UserRole{}
	this.Role = role
	this.ContentObject = contentObject
	return &this
}

// NewUserRoleWithDefaults instantiates a new UserRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserRoleWithDefaults() *UserRole {
	this := UserRole{}
	return &this
}

// GetRole returns the Role field value
func (o *UserRole) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *UserRole) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *UserRole) SetRole(v string) {
	o.Role = v
}

// GetContentObject returns the ContentObject field value
// If the value is explicit nil, the zero value for string will be returned
func (o *UserRole) GetContentObject() string {
	if o == nil || o.ContentObject.Get() == nil {
		var ret string
		return ret
	}

	return *o.ContentObject.Get()
}

// GetContentObjectOk returns a tuple with the ContentObject field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserRole) GetContentObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentObject.Get(), o.ContentObject.IsSet()
}

// SetContentObject sets field value
func (o *UserRole) SetContentObject(v string) {
	o.ContentObject.Set(&v)
}

func (o UserRole) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["role"] = o.Role
	toSerialize["content_object"] = o.ContentObject.Get()
	return toSerialize, nil
}

type NullableUserRole struct {
	value *UserRole
	isSet bool
}

func (v NullableUserRole) Get() *UserRole {
	return v.value
}

func (v *NullableUserRole) Set(val *UserRole) {
	v.value = val
	v.isSet = true
}

func (v NullableUserRole) IsSet() bool {
	return v.isSet
}

func (v *NullableUserRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserRole(val *UserRole) *NullableUserRole {
	return &NullableUserRole{value: val, isSet: true}
}

func (v NullableUserRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

