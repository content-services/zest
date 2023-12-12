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

// checks if the HeaderContentGuard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HeaderContentGuard{}

// HeaderContentGuard A serializer for HeaderContentGuard.
type HeaderContentGuard struct {
	// The unique name.
	Name string `json:"name"`
	// An optional description.
	Description NullableString `json:"description,omitempty"`
	// The header name the guard will check on.
	HeaderName string `json:"header_name"`
	// The value that will authorize the request.
	HeaderValue string `json:"header_value"`
	// A JQ syntax compatible filter. If jq_filter is not set, then the value willonly be Base64 decoded and checked as an explicit string match.
	JqFilter NullableString `json:"jq_filter,omitempty"`
}

type _HeaderContentGuard HeaderContentGuard

// NewHeaderContentGuard instantiates a new HeaderContentGuard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeaderContentGuard(name string, headerName string, headerValue string) *HeaderContentGuard {
	this := HeaderContentGuard{}
	this.Name = name
	this.HeaderName = headerName
	this.HeaderValue = headerValue
	return &this
}

// NewHeaderContentGuardWithDefaults instantiates a new HeaderContentGuard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeaderContentGuardWithDefaults() *HeaderContentGuard {
	this := HeaderContentGuard{}
	return &this
}

// GetName returns the Name field value
func (o *HeaderContentGuard) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *HeaderContentGuard) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *HeaderContentGuard) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeaderContentGuard) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeaderContentGuard) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *HeaderContentGuard) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *HeaderContentGuard) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *HeaderContentGuard) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *HeaderContentGuard) UnsetDescription() {
	o.Description.Unset()
}

// GetHeaderName returns the HeaderName field value
func (o *HeaderContentGuard) GetHeaderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HeaderName
}

// GetHeaderNameOk returns a tuple with the HeaderName field value
// and a boolean to check if the value has been set.
func (o *HeaderContentGuard) GetHeaderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HeaderName, true
}

// SetHeaderName sets field value
func (o *HeaderContentGuard) SetHeaderName(v string) {
	o.HeaderName = v
}

// GetHeaderValue returns the HeaderValue field value
func (o *HeaderContentGuard) GetHeaderValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HeaderValue
}

// GetHeaderValueOk returns a tuple with the HeaderValue field value
// and a boolean to check if the value has been set.
func (o *HeaderContentGuard) GetHeaderValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HeaderValue, true
}

// SetHeaderValue sets field value
func (o *HeaderContentGuard) SetHeaderValue(v string) {
	o.HeaderValue = v
}

// GetJqFilter returns the JqFilter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeaderContentGuard) GetJqFilter() string {
	if o == nil || IsNil(o.JqFilter.Get()) {
		var ret string
		return ret
	}
	return *o.JqFilter.Get()
}

// GetJqFilterOk returns a tuple with the JqFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeaderContentGuard) GetJqFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.JqFilter.Get(), o.JqFilter.IsSet()
}

// HasJqFilter returns a boolean if a field has been set.
func (o *HeaderContentGuard) HasJqFilter() bool {
	if o != nil && o.JqFilter.IsSet() {
		return true
	}

	return false
}

// SetJqFilter gets a reference to the given NullableString and assigns it to the JqFilter field.
func (o *HeaderContentGuard) SetJqFilter(v string) {
	o.JqFilter.Set(&v)
}
// SetJqFilterNil sets the value for JqFilter to be an explicit nil
func (o *HeaderContentGuard) SetJqFilterNil() {
	o.JqFilter.Set(nil)
}

// UnsetJqFilter ensures that no value is present for JqFilter, not even an explicit nil
func (o *HeaderContentGuard) UnsetJqFilter() {
	o.JqFilter.Unset()
}

func (o HeaderContentGuard) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HeaderContentGuard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["header_name"] = o.HeaderName
	toSerialize["header_value"] = o.HeaderValue
	if o.JqFilter.IsSet() {
		toSerialize["jq_filter"] = o.JqFilter.Get()
	}
	return toSerialize, nil
}

func (o *HeaderContentGuard) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"header_name",
		"header_value",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varHeaderContentGuard := _HeaderContentGuard{}

	err = json.Unmarshal(bytes, &varHeaderContentGuard)

	if err != nil {
		return err
	}

	*o = HeaderContentGuard(varHeaderContentGuard)

	return err
}

type NullableHeaderContentGuard struct {
	value *HeaderContentGuard
	isSet bool
}

func (v NullableHeaderContentGuard) Get() *HeaderContentGuard {
	return v.value
}

func (v *NullableHeaderContentGuard) Set(val *HeaderContentGuard) {
	v.value = val
	v.isSet = true
}

func (v NullableHeaderContentGuard) IsSet() bool {
	return v.isSet
}

func (v *NullableHeaderContentGuard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeaderContentGuard(val *HeaderContentGuard) *NullableHeaderContentGuard {
	return &NullableHeaderContentGuard{value: val, isSet: true}
}

func (v NullableHeaderContentGuard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeaderContentGuard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


