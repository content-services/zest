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

// checks if the PaginatedArtifactResponseList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedArtifactResponseList{}

// PaginatedArtifactResponseList struct for PaginatedArtifactResponseList
type PaginatedArtifactResponseList struct {
	Count *int32 `json:"count,omitempty"`
	Next NullableString `json:"next,omitempty"`
	Previous NullableString `json:"previous,omitempty"`
	Results []ArtifactResponse `json:"results,omitempty"`
}

// NewPaginatedArtifactResponseList instantiates a new PaginatedArtifactResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedArtifactResponseList() *PaginatedArtifactResponseList {
	this := PaginatedArtifactResponseList{}
	return &this
}

// NewPaginatedArtifactResponseListWithDefaults instantiates a new PaginatedArtifactResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedArtifactResponseListWithDefaults() *PaginatedArtifactResponseList {
	this := PaginatedArtifactResponseList{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PaginatedArtifactResponseList) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedArtifactResponseList) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PaginatedArtifactResponseList) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PaginatedArtifactResponseList) SetCount(v int32) {
	o.Count = &v
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedArtifactResponseList) GetNext() string {
	if o == nil || IsNil(o.Next.Get()) {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedArtifactResponseList) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedArtifactResponseList) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *PaginatedArtifactResponseList) SetNext(v string) {
	o.Next.Set(&v)
}
// SetNextNil sets the value for Next to be an explicit nil
func (o *PaginatedArtifactResponseList) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *PaginatedArtifactResponseList) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedArtifactResponseList) GetPrevious() string {
	if o == nil || IsNil(o.Previous.Get()) {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedArtifactResponseList) GetPreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedArtifactResponseList) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *PaginatedArtifactResponseList) SetPrevious(v string) {
	o.Previous.Set(&v)
}
// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *PaginatedArtifactResponseList) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *PaginatedArtifactResponseList) UnsetPrevious() {
	o.Previous.Unset()
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedArtifactResponseList) GetResults() []ArtifactResponse {
	if o == nil || IsNil(o.Results) {
		var ret []ArtifactResponse
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedArtifactResponseList) GetResultsOk() ([]ArtifactResponse, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedArtifactResponseList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ArtifactResponse and assigns it to the Results field.
func (o *PaginatedArtifactResponseList) SetResults(v []ArtifactResponse) {
	o.Results = v
}

func (o PaginatedArtifactResponseList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedArtifactResponseList) ToMap() (map[string]interface{}, error) {
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

type NullablePaginatedArtifactResponseList struct {
	value *PaginatedArtifactResponseList
	isSet bool
}

func (v NullablePaginatedArtifactResponseList) Get() *PaginatedArtifactResponseList {
	return v.value
}

func (v *NullablePaginatedArtifactResponseList) Set(val *PaginatedArtifactResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedArtifactResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedArtifactResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedArtifactResponseList(val *PaginatedArtifactResponseList) *NullablePaginatedArtifactResponseList {
	return &NullablePaginatedArtifactResponseList{value: val, isSet: true}
}

func (v NullablePaginatedArtifactResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedArtifactResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


