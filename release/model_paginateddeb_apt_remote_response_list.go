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

// checks if the PaginateddebAptRemoteResponseList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginateddebAptRemoteResponseList{}

// PaginateddebAptRemoteResponseList struct for PaginateddebAptRemoteResponseList
type PaginateddebAptRemoteResponseList struct {
	Count *int32 `json:"count,omitempty"`
	Next NullableString `json:"next,omitempty"`
	Previous NullableString `json:"previous,omitempty"`
	Results []DebAptRemoteResponse `json:"results,omitempty"`
}

// NewPaginateddebAptRemoteResponseList instantiates a new PaginateddebAptRemoteResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginateddebAptRemoteResponseList() *PaginateddebAptRemoteResponseList {
	this := PaginateddebAptRemoteResponseList{}
	return &this
}

// NewPaginateddebAptRemoteResponseListWithDefaults instantiates a new PaginateddebAptRemoteResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginateddebAptRemoteResponseListWithDefaults() *PaginateddebAptRemoteResponseList {
	this := PaginateddebAptRemoteResponseList{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PaginateddebAptRemoteResponseList) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginateddebAptRemoteResponseList) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PaginateddebAptRemoteResponseList) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PaginateddebAptRemoteResponseList) SetCount(v int32) {
	o.Count = &v
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginateddebAptRemoteResponseList) GetNext() string {
	if o == nil || IsNil(o.Next.Get()) {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginateddebAptRemoteResponseList) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *PaginateddebAptRemoteResponseList) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *PaginateddebAptRemoteResponseList) SetNext(v string) {
	o.Next.Set(&v)
}
// SetNextNil sets the value for Next to be an explicit nil
func (o *PaginateddebAptRemoteResponseList) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *PaginateddebAptRemoteResponseList) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginateddebAptRemoteResponseList) GetPrevious() string {
	if o == nil || IsNil(o.Previous.Get()) {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginateddebAptRemoteResponseList) GetPreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginateddebAptRemoteResponseList) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *PaginateddebAptRemoteResponseList) SetPrevious(v string) {
	o.Previous.Set(&v)
}
// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *PaginateddebAptRemoteResponseList) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *PaginateddebAptRemoteResponseList) UnsetPrevious() {
	o.Previous.Unset()
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginateddebAptRemoteResponseList) GetResults() []DebAptRemoteResponse {
	if o == nil || IsNil(o.Results) {
		var ret []DebAptRemoteResponse
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginateddebAptRemoteResponseList) GetResultsOk() ([]DebAptRemoteResponse, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginateddebAptRemoteResponseList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []DebAptRemoteResponse and assigns it to the Results field.
func (o *PaginateddebAptRemoteResponseList) SetResults(v []DebAptRemoteResponse) {
	o.Results = v
}

func (o PaginateddebAptRemoteResponseList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginateddebAptRemoteResponseList) ToMap() (map[string]interface{}, error) {
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

type NullablePaginateddebAptRemoteResponseList struct {
	value *PaginateddebAptRemoteResponseList
	isSet bool
}

func (v NullablePaginateddebAptRemoteResponseList) Get() *PaginateddebAptRemoteResponseList {
	return v.value
}

func (v *NullablePaginateddebAptRemoteResponseList) Set(val *PaginateddebAptRemoteResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginateddebAptRemoteResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginateddebAptRemoteResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginateddebAptRemoteResponseList(val *PaginateddebAptRemoteResponseList) *NullablePaginateddebAptRemoteResponseList {
	return &NullablePaginateddebAptRemoteResponseList{value: val, isSet: true}
}

func (v NullablePaginateddebAptRemoteResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginateddebAptRemoteResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


