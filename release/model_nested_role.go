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

// checks if the NestedRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedRole{}

// NestedRole Serializer to add/remove object roles to/from users/groups.  This is used in conjunction with ``pulpcore.app.viewsets.base.RolesMixin`` and requires the underlying object to be passed as ``content_object`` in the context.
type NestedRole struct {
	Users []string `json:"users,omitempty"`
	Groups []string `json:"groups,omitempty"`
	Role string `json:"role"`
}

// NewNestedRole instantiates a new NestedRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedRole(role string) *NestedRole {
	this := NestedRole{}
	this.Role = role
	return &this
}

// NewNestedRoleWithDefaults instantiates a new NestedRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedRoleWithDefaults() *NestedRole {
	this := NestedRole{}
	return &this
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *NestedRole) GetUsers() []string {
	if o == nil || IsNil(o.Users) {
		var ret []string
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedRole) GetUsersOk() ([]string, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *NestedRole) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []string and assigns it to the Users field.
func (o *NestedRole) SetUsers(v []string) {
	o.Users = v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *NestedRole) GetGroups() []string {
	if o == nil || IsNil(o.Groups) {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedRole) GetGroupsOk() ([]string, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *NestedRole) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []string and assigns it to the Groups field.
func (o *NestedRole) SetGroups(v []string) {
	o.Groups = v
}

// GetRole returns the Role field value
func (o *NestedRole) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *NestedRole) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *NestedRole) SetRole(v string) {
	o.Role = v
}

func (o NestedRole) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	toSerialize["role"] = o.Role
	return toSerialize, nil
}

type NullableNestedRole struct {
	value *NestedRole
	isSet bool
}

func (v NullableNestedRole) Get() *NestedRole {
	return v.value
}

func (v *NullableNestedRole) Set(val *NestedRole) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedRole) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedRole(val *NestedRole) *NullableNestedRole {
	return &NullableNestedRole{value: val, isSet: true}
}

func (v NullableNestedRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


