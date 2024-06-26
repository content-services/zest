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

// checks if the PaginatedpythonPythonRemoteResponseList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedpythonPythonRemoteResponseList{}

// PaginatedpythonPythonRemoteResponseList struct for PaginatedpythonPythonRemoteResponseList
type PaginatedpythonPythonRemoteResponseList struct {
	Count int32 `json:"count"`
	Next NullableString `json:"next,omitempty"`
	Previous NullableString `json:"previous,omitempty"`
	Results []PythonPythonRemoteResponse `json:"results"`
	AdditionalProperties map[string]interface{}
}

type _PaginatedpythonPythonRemoteResponseList PaginatedpythonPythonRemoteResponseList

// NewPaginatedpythonPythonRemoteResponseList instantiates a new PaginatedpythonPythonRemoteResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedpythonPythonRemoteResponseList(count int32, results []PythonPythonRemoteResponse) *PaginatedpythonPythonRemoteResponseList {
	this := PaginatedpythonPythonRemoteResponseList{}
	this.Count = count
	this.Results = results
	return &this
}

// NewPaginatedpythonPythonRemoteResponseListWithDefaults instantiates a new PaginatedpythonPythonRemoteResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedpythonPythonRemoteResponseListWithDefaults() *PaginatedpythonPythonRemoteResponseList {
	this := PaginatedpythonPythonRemoteResponseList{}
	return &this
}

// GetCount returns the Count field value
func (o *PaginatedpythonPythonRemoteResponseList) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *PaginatedpythonPythonRemoteResponseList) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *PaginatedpythonPythonRemoteResponseList) SetCount(v int32) {
	o.Count = v
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedpythonPythonRemoteResponseList) GetNext() string {
	if o == nil || IsNil(o.Next.Get()) {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedpythonPythonRemoteResponseList) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedpythonPythonRemoteResponseList) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *PaginatedpythonPythonRemoteResponseList) SetNext(v string) {
	o.Next.Set(&v)
}
// SetNextNil sets the value for Next to be an explicit nil
func (o *PaginatedpythonPythonRemoteResponseList) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *PaginatedpythonPythonRemoteResponseList) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedpythonPythonRemoteResponseList) GetPrevious() string {
	if o == nil || IsNil(o.Previous.Get()) {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedpythonPythonRemoteResponseList) GetPreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedpythonPythonRemoteResponseList) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *PaginatedpythonPythonRemoteResponseList) SetPrevious(v string) {
	o.Previous.Set(&v)
}
// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *PaginatedpythonPythonRemoteResponseList) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *PaginatedpythonPythonRemoteResponseList) UnsetPrevious() {
	o.Previous.Unset()
}

// GetResults returns the Results field value
func (o *PaginatedpythonPythonRemoteResponseList) GetResults() []PythonPythonRemoteResponse {
	if o == nil {
		var ret []PythonPythonRemoteResponse
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *PaginatedpythonPythonRemoteResponseList) GetResultsOk() ([]PythonPythonRemoteResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *PaginatedpythonPythonRemoteResponseList) SetResults(v []PythonPythonRemoteResponse) {
	o.Results = v
}

func (o PaginatedpythonPythonRemoteResponseList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedpythonPythonRemoteResponseList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	if o.Next.IsSet() {
		toSerialize["next"] = o.Next.Get()
	}
	if o.Previous.IsSet() {
		toSerialize["previous"] = o.Previous.Get()
	}
	toSerialize["results"] = o.Results

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaginatedpythonPythonRemoteResponseList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"count",
		"results",
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

	varPaginatedpythonPythonRemoteResponseList := _PaginatedpythonPythonRemoteResponseList{}

	err = json.Unmarshal(data, &varPaginatedpythonPythonRemoteResponseList)

	if err != nil {
		return err
	}

	*o = PaginatedpythonPythonRemoteResponseList(varPaginatedpythonPythonRemoteResponseList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "count")
		delete(additionalProperties, "next")
		delete(additionalProperties, "previous")
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaginatedpythonPythonRemoteResponseList struct {
	value *PaginatedpythonPythonRemoteResponseList
	isSet bool
}

func (v NullablePaginatedpythonPythonRemoteResponseList) Get() *PaginatedpythonPythonRemoteResponseList {
	return v.value
}

func (v *NullablePaginatedpythonPythonRemoteResponseList) Set(val *PaginatedpythonPythonRemoteResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedpythonPythonRemoteResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedpythonPythonRemoteResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedpythonPythonRemoteResponseList(val *PaginatedpythonPythonRemoteResponseList) *NullablePaginatedpythonPythonRemoteResponseList {
	return &NullablePaginatedpythonPythonRemoteResponseList{value: val, isSet: true}
}

func (v NullablePaginatedpythonPythonRemoteResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedpythonPythonRemoteResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


