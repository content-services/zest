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

// checks if the PaginatedGalaxyRoleVersionResponseList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedGalaxyRoleVersionResponseList{}

// PaginatedGalaxyRoleVersionResponseList struct for PaginatedGalaxyRoleVersionResponseList
type PaginatedGalaxyRoleVersionResponseList struct {
	Count *int32 `json:"count,omitempty"`
	Next NullableString `json:"next,omitempty"`
	Previous NullableString `json:"previous,omitempty"`
	Results []GalaxyRoleVersionResponse `json:"results,omitempty"`
}

// NewPaginatedGalaxyRoleVersionResponseList instantiates a new PaginatedGalaxyRoleVersionResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedGalaxyRoleVersionResponseList() *PaginatedGalaxyRoleVersionResponseList {
	this := PaginatedGalaxyRoleVersionResponseList{}
	return &this
}

// NewPaginatedGalaxyRoleVersionResponseListWithDefaults instantiates a new PaginatedGalaxyRoleVersionResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedGalaxyRoleVersionResponseListWithDefaults() *PaginatedGalaxyRoleVersionResponseList {
	this := PaginatedGalaxyRoleVersionResponseList{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PaginatedGalaxyRoleVersionResponseList) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedGalaxyRoleVersionResponseList) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PaginatedGalaxyRoleVersionResponseList) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PaginatedGalaxyRoleVersionResponseList) SetCount(v int32) {
	o.Count = &v
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedGalaxyRoleVersionResponseList) GetNext() string {
	if o == nil || IsNil(o.Next.Get()) {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedGalaxyRoleVersionResponseList) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedGalaxyRoleVersionResponseList) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *PaginatedGalaxyRoleVersionResponseList) SetNext(v string) {
	o.Next.Set(&v)
}
// SetNextNil sets the value for Next to be an explicit nil
func (o *PaginatedGalaxyRoleVersionResponseList) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *PaginatedGalaxyRoleVersionResponseList) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedGalaxyRoleVersionResponseList) GetPrevious() string {
	if o == nil || IsNil(o.Previous.Get()) {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedGalaxyRoleVersionResponseList) GetPreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedGalaxyRoleVersionResponseList) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *PaginatedGalaxyRoleVersionResponseList) SetPrevious(v string) {
	o.Previous.Set(&v)
}
// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *PaginatedGalaxyRoleVersionResponseList) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *PaginatedGalaxyRoleVersionResponseList) UnsetPrevious() {
	o.Previous.Unset()
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedGalaxyRoleVersionResponseList) GetResults() []GalaxyRoleVersionResponse {
	if o == nil || IsNil(o.Results) {
		var ret []GalaxyRoleVersionResponse
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedGalaxyRoleVersionResponseList) GetResultsOk() ([]GalaxyRoleVersionResponse, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedGalaxyRoleVersionResponseList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []GalaxyRoleVersionResponse and assigns it to the Results field.
func (o *PaginatedGalaxyRoleVersionResponseList) SetResults(v []GalaxyRoleVersionResponse) {
	o.Results = v
}

func (o PaginatedGalaxyRoleVersionResponseList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedGalaxyRoleVersionResponseList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if o.Next.IsSet() {
		toSerialize["next"] = o.Next.Get()
	}
	if o.Previous.IsSet() {
		toSerialize["previous"] = o.Previous.Get()
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullablePaginatedGalaxyRoleVersionResponseList struct {
	value *PaginatedGalaxyRoleVersionResponseList
	isSet bool
}

func (v NullablePaginatedGalaxyRoleVersionResponseList) Get() *PaginatedGalaxyRoleVersionResponseList {
	return v.value
}

func (v *NullablePaginatedGalaxyRoleVersionResponseList) Set(val *PaginatedGalaxyRoleVersionResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedGalaxyRoleVersionResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedGalaxyRoleVersionResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedGalaxyRoleVersionResponseList(val *PaginatedGalaxyRoleVersionResponseList) *NullablePaginatedGalaxyRoleVersionResponseList {
	return &NullablePaginatedGalaxyRoleVersionResponseList{value: val, isSet: true}
}

func (v NullablePaginatedGalaxyRoleVersionResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedGalaxyRoleVersionResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


