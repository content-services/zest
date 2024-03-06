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

// checks if the GroupUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupUser{}

// GroupUser Serializer for Users that belong to a Group.
type GroupUser struct {
	// Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.
	Username string `json:"username"`
	AdditionalProperties map[string]interface{}
}

type _GroupUser GroupUser

// NewGroupUser instantiates a new GroupUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupUser(username string) *GroupUser {
	this := GroupUser{}
	this.Username = username
	return &this
}

// NewGroupUserWithDefaults instantiates a new GroupUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupUserWithDefaults() *GroupUser {
	this := GroupUser{}
	return &this
}

// GetUsername returns the Username field value
func (o *GroupUser) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *GroupUser) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *GroupUser) SetUsername(v string) {
	o.Username = v
}

func (o GroupUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["username"] = o.Username

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupUser) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"username",
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

	varGroupUser := _GroupUser{}

	err = json.Unmarshal(data, &varGroupUser)

	if err != nil {
		return err
	}

	*o = GroupUser(varGroupUser)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "username")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupUser struct {
	value *GroupUser
	isSet bool
}

func (v NullableGroupUser) Get() *GroupUser {
	return v.value
}

func (v *NullableGroupUser) Set(val *GroupUser) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupUser) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupUser(val *GroupUser) *NullableGroupUser {
	return &NullableGroupUser{value: val, isSet: true}
}

func (v NullableGroupUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


