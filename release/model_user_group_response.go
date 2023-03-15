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

// checks if the UserGroupResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserGroupResponse{}

// UserGroupResponse Serializer for Groups that belong to an User.
type UserGroupResponse struct {
	// Name.
	Name string `json:"name"`
	PulpHref *string `json:"pulp_href,omitempty"`
}

// NewUserGroupResponse instantiates a new UserGroupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGroupResponse(name string) *UserGroupResponse {
	this := UserGroupResponse{}
	this.Name = name
	return &this
}

// NewUserGroupResponseWithDefaults instantiates a new UserGroupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGroupResponseWithDefaults() *UserGroupResponse {
	this := UserGroupResponse{}
	return &this
}

// GetName returns the Name field value
func (o *UserGroupResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserGroupResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserGroupResponse) SetName(v string) {
	o.Name = v
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *UserGroupResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *UserGroupResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *UserGroupResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

func (o UserGroupResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserGroupResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	// skip: pulp_href is readOnly
	return toSerialize, nil
}

type NullableUserGroupResponse struct {
	value *UserGroupResponse
	isSet bool
}

func (v NullableUserGroupResponse) Get() *UserGroupResponse {
	return v.value
}

func (v *NullableUserGroupResponse) Set(val *UserGroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGroupResponse(val *UserGroupResponse) *NullableUserGroupResponse {
	return &NullableUserGroupResponse{value: val, isSet: true}
}

func (v NullableUserGroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


