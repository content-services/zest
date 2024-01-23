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

// checks if the PatchedHeaderContentGuard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedHeaderContentGuard{}

// PatchedHeaderContentGuard A serializer for HeaderContentGuard.
type PatchedHeaderContentGuard struct {
	// The unique name.
	Name *string `json:"name,omitempty"`
	// An optional description.
	Description NullableString `json:"description,omitempty"`
	// The header name the guard will check on.
	HeaderName *string `json:"header_name,omitempty"`
	// The value that will authorize the request.
	HeaderValue *string `json:"header_value,omitempty"`
	// A JQ syntax compatible filter. If jq_filter is not set, then the value willonly be Base64 decoded and checked as an explicit string match.
	JqFilter NullableString `json:"jq_filter,omitempty"`
}

// NewPatchedHeaderContentGuard instantiates a new PatchedHeaderContentGuard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedHeaderContentGuard() *PatchedHeaderContentGuard {
	this := PatchedHeaderContentGuard{}
	return &this
}

// NewPatchedHeaderContentGuardWithDefaults instantiates a new PatchedHeaderContentGuard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedHeaderContentGuardWithDefaults() *PatchedHeaderContentGuard {
	this := PatchedHeaderContentGuard{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedHeaderContentGuard) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedHeaderContentGuard) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedHeaderContentGuard) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedHeaderContentGuard) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedHeaderContentGuard) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedHeaderContentGuard) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedHeaderContentGuard) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PatchedHeaderContentGuard) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PatchedHeaderContentGuard) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PatchedHeaderContentGuard) UnsetDescription() {
	o.Description.Unset()
}

// GetHeaderName returns the HeaderName field value if set, zero value otherwise.
func (o *PatchedHeaderContentGuard) GetHeaderName() string {
	if o == nil || IsNil(o.HeaderName) {
		var ret string
		return ret
	}
	return *o.HeaderName
}

// GetHeaderNameOk returns a tuple with the HeaderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedHeaderContentGuard) GetHeaderNameOk() (*string, bool) {
	if o == nil || IsNil(o.HeaderName) {
		return nil, false
	}
	return o.HeaderName, true
}

// HasHeaderName returns a boolean if a field has been set.
func (o *PatchedHeaderContentGuard) HasHeaderName() bool {
	if o != nil && !IsNil(o.HeaderName) {
		return true
	}

	return false
}

// SetHeaderName gets a reference to the given string and assigns it to the HeaderName field.
func (o *PatchedHeaderContentGuard) SetHeaderName(v string) {
	o.HeaderName = &v
}

// GetHeaderValue returns the HeaderValue field value if set, zero value otherwise.
func (o *PatchedHeaderContentGuard) GetHeaderValue() string {
	if o == nil || IsNil(o.HeaderValue) {
		var ret string
		return ret
	}
	return *o.HeaderValue
}

// GetHeaderValueOk returns a tuple with the HeaderValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedHeaderContentGuard) GetHeaderValueOk() (*string, bool) {
	if o == nil || IsNil(o.HeaderValue) {
		return nil, false
	}
	return o.HeaderValue, true
}

// HasHeaderValue returns a boolean if a field has been set.
func (o *PatchedHeaderContentGuard) HasHeaderValue() bool {
	if o != nil && !IsNil(o.HeaderValue) {
		return true
	}

	return false
}

// SetHeaderValue gets a reference to the given string and assigns it to the HeaderValue field.
func (o *PatchedHeaderContentGuard) SetHeaderValue(v string) {
	o.HeaderValue = &v
}

// GetJqFilter returns the JqFilter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedHeaderContentGuard) GetJqFilter() string {
	if o == nil || IsNil(o.JqFilter.Get()) {
		var ret string
		return ret
	}
	return *o.JqFilter.Get()
}

// GetJqFilterOk returns a tuple with the JqFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedHeaderContentGuard) GetJqFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.JqFilter.Get(), o.JqFilter.IsSet()
}

// HasJqFilter returns a boolean if a field has been set.
func (o *PatchedHeaderContentGuard) HasJqFilter() bool {
	if o != nil && o.JqFilter.IsSet() {
		return true
	}

	return false
}

// SetJqFilter gets a reference to the given NullableString and assigns it to the JqFilter field.
func (o *PatchedHeaderContentGuard) SetJqFilter(v string) {
	o.JqFilter.Set(&v)
}
// SetJqFilterNil sets the value for JqFilter to be an explicit nil
func (o *PatchedHeaderContentGuard) SetJqFilterNil() {
	o.JqFilter.Set(nil)
}

// UnsetJqFilter ensures that no value is present for JqFilter, not even an explicit nil
func (o *PatchedHeaderContentGuard) UnsetJqFilter() {
	o.JqFilter.Unset()
}

func (o PatchedHeaderContentGuard) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedHeaderContentGuard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.HeaderName) {
		toSerialize["header_name"] = o.HeaderName
	}
	if !IsNil(o.HeaderValue) {
		toSerialize["header_value"] = o.HeaderValue
	}
	if o.JqFilter.IsSet() {
		toSerialize["jq_filter"] = o.JqFilter.Get()
	}
	return toSerialize, nil
}

type NullablePatchedHeaderContentGuard struct {
	value *PatchedHeaderContentGuard
	isSet bool
}

func (v NullablePatchedHeaderContentGuard) Get() *PatchedHeaderContentGuard {
	return v.value
}

func (v *NullablePatchedHeaderContentGuard) Set(val *PatchedHeaderContentGuard) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedHeaderContentGuard) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedHeaderContentGuard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedHeaderContentGuard(val *PatchedHeaderContentGuard) *NullablePatchedHeaderContentGuard {
	return &NullablePatchedHeaderContentGuard{value: val, isSet: true}
}

func (v NullablePatchedHeaderContentGuard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedHeaderContentGuard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

