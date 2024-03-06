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

// checks if the NestedRoleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedRoleResponse{}

// NestedRoleResponse Serializer to add/remove object roles to/from users/groups.This is used in conjunction with ``pulpcore.app.viewsets.base.RolesMixin`` and requires theunderlying object to be passed as ``content_object`` in the context.
type NestedRoleResponse struct {
	Users []string `json:"users,omitempty"`
	Groups []string `json:"groups,omitempty"`
	Role string `json:"role"`
	AdditionalProperties map[string]interface{}
}

type _NestedRoleResponse NestedRoleResponse

// NewNestedRoleResponse instantiates a new NestedRoleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedRoleResponse(role string) *NestedRoleResponse {
	this := NestedRoleResponse{}
	this.Role = role
	return &this
}

// NewNestedRoleResponseWithDefaults instantiates a new NestedRoleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedRoleResponseWithDefaults() *NestedRoleResponse {
	this := NestedRoleResponse{}
	return &this
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *NestedRoleResponse) GetUsers() []string {
	if o == nil || IsNil(o.Users) {
		var ret []string
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedRoleResponse) GetUsersOk() ([]string, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *NestedRoleResponse) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []string and assigns it to the Users field.
func (o *NestedRoleResponse) SetUsers(v []string) {
	o.Users = v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *NestedRoleResponse) GetGroups() []string {
	if o == nil || IsNil(o.Groups) {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedRoleResponse) GetGroupsOk() ([]string, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *NestedRoleResponse) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []string and assigns it to the Groups field.
func (o *NestedRoleResponse) SetGroups(v []string) {
	o.Groups = v
}

// GetRole returns the Role field value
func (o *NestedRoleResponse) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *NestedRoleResponse) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *NestedRoleResponse) SetRole(v string) {
	o.Role = v
}

func (o NestedRoleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedRoleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	toSerialize["role"] = o.Role

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedRoleResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"role",
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

	varNestedRoleResponse := _NestedRoleResponse{}

	err = json.Unmarshal(data, &varNestedRoleResponse)

	if err != nil {
		return err
	}

	*o = NestedRoleResponse(varNestedRoleResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "users")
		delete(additionalProperties, "groups")
		delete(additionalProperties, "role")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedRoleResponse struct {
	value *NestedRoleResponse
	isSet bool
}

func (v NullableNestedRoleResponse) Get() *NestedRoleResponse {
	return v.value
}

func (v *NullableNestedRoleResponse) Set(val *NestedRoleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedRoleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedRoleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedRoleResponse(val *NestedRoleResponse) *NullableNestedRoleResponse {
	return &NullableNestedRoleResponse{value: val, isSet: true}
}

func (v NullableNestedRoleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedRoleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


