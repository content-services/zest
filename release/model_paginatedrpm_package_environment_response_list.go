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

// checks if the PaginatedrpmPackageEnvironmentResponseList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedrpmPackageEnvironmentResponseList{}

// PaginatedrpmPackageEnvironmentResponseList struct for PaginatedrpmPackageEnvironmentResponseList
type PaginatedrpmPackageEnvironmentResponseList struct {
	Count *int32 `json:"count,omitempty"`
	Next NullableString `json:"next,omitempty"`
	Previous NullableString `json:"previous,omitempty"`
	Results []RpmPackageEnvironmentResponse `json:"results,omitempty"`
}

// NewPaginatedrpmPackageEnvironmentResponseList instantiates a new PaginatedrpmPackageEnvironmentResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedrpmPackageEnvironmentResponseList() *PaginatedrpmPackageEnvironmentResponseList {
	this := PaginatedrpmPackageEnvironmentResponseList{}
	return &this
}

// NewPaginatedrpmPackageEnvironmentResponseListWithDefaults instantiates a new PaginatedrpmPackageEnvironmentResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedrpmPackageEnvironmentResponseListWithDefaults() *PaginatedrpmPackageEnvironmentResponseList {
	this := PaginatedrpmPackageEnvironmentResponseList{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PaginatedrpmPackageEnvironmentResponseList) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedrpmPackageEnvironmentResponseList) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PaginatedrpmPackageEnvironmentResponseList) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PaginatedrpmPackageEnvironmentResponseList) SetCount(v int32) {
	o.Count = &v
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedrpmPackageEnvironmentResponseList) GetNext() string {
	if o == nil || IsNil(o.Next.Get()) {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedrpmPackageEnvironmentResponseList) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedrpmPackageEnvironmentResponseList) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *PaginatedrpmPackageEnvironmentResponseList) SetNext(v string) {
	o.Next.Set(&v)
}
// SetNextNil sets the value for Next to be an explicit nil
func (o *PaginatedrpmPackageEnvironmentResponseList) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *PaginatedrpmPackageEnvironmentResponseList) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedrpmPackageEnvironmentResponseList) GetPrevious() string {
	if o == nil || IsNil(o.Previous.Get()) {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedrpmPackageEnvironmentResponseList) GetPreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedrpmPackageEnvironmentResponseList) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *PaginatedrpmPackageEnvironmentResponseList) SetPrevious(v string) {
	o.Previous.Set(&v)
}
// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *PaginatedrpmPackageEnvironmentResponseList) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *PaginatedrpmPackageEnvironmentResponseList) UnsetPrevious() {
	o.Previous.Unset()
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedrpmPackageEnvironmentResponseList) GetResults() []RpmPackageEnvironmentResponse {
	if o == nil || IsNil(o.Results) {
		var ret []RpmPackageEnvironmentResponse
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedrpmPackageEnvironmentResponseList) GetResultsOk() ([]RpmPackageEnvironmentResponse, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedrpmPackageEnvironmentResponseList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []RpmPackageEnvironmentResponse and assigns it to the Results field.
func (o *PaginatedrpmPackageEnvironmentResponseList) SetResults(v []RpmPackageEnvironmentResponse) {
	o.Results = v
}

func (o PaginatedrpmPackageEnvironmentResponseList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedrpmPackageEnvironmentResponseList) ToMap() (map[string]interface{}, error) {
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

type NullablePaginatedrpmPackageEnvironmentResponseList struct {
	value *PaginatedrpmPackageEnvironmentResponseList
	isSet bool
}

func (v NullablePaginatedrpmPackageEnvironmentResponseList) Get() *PaginatedrpmPackageEnvironmentResponseList {
	return v.value
}

func (v *NullablePaginatedrpmPackageEnvironmentResponseList) Set(val *PaginatedrpmPackageEnvironmentResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedrpmPackageEnvironmentResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedrpmPackageEnvironmentResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedrpmPackageEnvironmentResponseList(val *PaginatedrpmPackageEnvironmentResponseList) *NullablePaginatedrpmPackageEnvironmentResponseList {
	return &NullablePaginatedrpmPackageEnvironmentResponseList{value: val, isSet: true}
}

func (v NullablePaginatedrpmPackageEnvironmentResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedrpmPackageEnvironmentResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


