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

// checks if the UserRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserRole{}

// UserRole Serializer for UserRole.
type UserRole struct {
	Role string `json:"role"`
	// pulp_href of the object for which role permissions should be asserted. If set to 'null', permissions will act on either domain or model-level.
	ContentObject NullableString `json:"content_object"`
	// Domain this role should be applied on, mutually exclusive with content_object.
	Domain NullableString `json:"domain,omitempty"`
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

// GetDomain returns the Domain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserRole) GetDomain() string {
	if o == nil || IsNil(o.Domain.Get()) {
		var ret string
		return ret
	}
	return *o.Domain.Get()
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserRole) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Domain.Get(), o.Domain.IsSet()
}

// HasDomain returns a boolean if a field has been set.
func (o *UserRole) HasDomain() bool {
	if o != nil && o.Domain.IsSet() {
		return true
	}

	return false
}

// SetDomain gets a reference to the given NullableString and assigns it to the Domain field.
func (o *UserRole) SetDomain(v string) {
	o.Domain.Set(&v)
}
// SetDomainNil sets the value for Domain to be an explicit nil
func (o *UserRole) SetDomainNil() {
	o.Domain.Set(nil)
}

// UnsetDomain ensures that no value is present for Domain, not even an explicit nil
func (o *UserRole) UnsetDomain() {
	o.Domain.Unset()
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
	if o.Domain.IsSet() {
		toSerialize["domain"] = o.Domain.Get()
	}
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


