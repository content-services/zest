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

// checks if the PatchedRBACContentGuard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedRBACContentGuard{}

// PatchedRBACContentGuard Base serializer for use with :class:`pulpcore.app.models.Model`  This ensures that all Serializers provide values for the 'pulp_href` field.  The class provides a default for the ``ref_name`` attribute in the ModelSerializers's ``Meta`` class. This ensures that the OpenAPI definitions of plugins are namespaced properly.
type PatchedRBACContentGuard struct {
	// The unique name.
	Name *string `json:"name,omitempty"`
	// An optional description.
	Description NullableString `json:"description,omitempty"`
}

// NewPatchedRBACContentGuard instantiates a new PatchedRBACContentGuard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedRBACContentGuard() *PatchedRBACContentGuard {
	this := PatchedRBACContentGuard{}
	return &this
}

// NewPatchedRBACContentGuardWithDefaults instantiates a new PatchedRBACContentGuard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedRBACContentGuardWithDefaults() *PatchedRBACContentGuard {
	this := PatchedRBACContentGuard{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedRBACContentGuard) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRBACContentGuard) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedRBACContentGuard) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedRBACContentGuard) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedRBACContentGuard) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedRBACContentGuard) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedRBACContentGuard) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PatchedRBACContentGuard) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PatchedRBACContentGuard) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PatchedRBACContentGuard) UnsetDescription() {
	o.Description.Unset()
}

func (o PatchedRBACContentGuard) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedRBACContentGuard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	return toSerialize, nil
}

type NullablePatchedRBACContentGuard struct {
	value *PatchedRBACContentGuard
	isSet bool
}

func (v NullablePatchedRBACContentGuard) Get() *PatchedRBACContentGuard {
	return v.value
}

func (v *NullablePatchedRBACContentGuard) Set(val *PatchedRBACContentGuard) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedRBACContentGuard) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedRBACContentGuard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedRBACContentGuard(val *PatchedRBACContentGuard) *NullablePatchedRBACContentGuard {
	return &NullablePatchedRBACContentGuard{value: val, isSet: true}
}

func (v NullablePatchedRBACContentGuard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedRBACContentGuard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


