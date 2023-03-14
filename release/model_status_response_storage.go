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

// checks if the StatusResponseStorage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusResponseStorage{}

// StatusResponseStorage Storage information
type StatusResponseStorage struct {
	// Total number of bytes
	Total int32 `json:"total"`
	// Number of bytes in use
	Used int32 `json:"used"`
	// Number of free bytes
	Free int32 `json:"free"`
}

// NewStatusResponseStorage instantiates a new StatusResponseStorage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusResponseStorage(total int32, used int32, free int32) *StatusResponseStorage {
	this := StatusResponseStorage{}
	this.Total = total
	this.Used = used
	this.Free = free
	return &this
}

// NewStatusResponseStorageWithDefaults instantiates a new StatusResponseStorage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusResponseStorageWithDefaults() *StatusResponseStorage {
	this := StatusResponseStorage{}
	return &this
}

// GetTotal returns the Total field value
func (o *StatusResponseStorage) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *StatusResponseStorage) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *StatusResponseStorage) SetTotal(v int32) {
	o.Total = v
}

// GetUsed returns the Used field value
func (o *StatusResponseStorage) GetUsed() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Used
}

// GetUsedOk returns a tuple with the Used field value
// and a boolean to check if the value has been set.
func (o *StatusResponseStorage) GetUsedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Used, true
}

// SetUsed sets field value
func (o *StatusResponseStorage) SetUsed(v int32) {
	o.Used = v
}

// GetFree returns the Free field value
func (o *StatusResponseStorage) GetFree() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Free
}

// GetFreeOk returns a tuple with the Free field value
// and a boolean to check if the value has been set.
func (o *StatusResponseStorage) GetFreeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Free, true
}

// SetFree sets field value
func (o *StatusResponseStorage) SetFree(v int32) {
	o.Free = v
}

func (o StatusResponseStorage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatusResponseStorage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["total"] = o.Total
	toSerialize["used"] = o.Used
	toSerialize["free"] = o.Free
	return toSerialize, nil
}

type NullableStatusResponseStorage struct {
	value *StatusResponseStorage
	isSet bool
}

func (v NullableStatusResponseStorage) Get() *StatusResponseStorage {
	return v.value
}

func (v *NullableStatusResponseStorage) Set(val *StatusResponseStorage) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusResponseStorage) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusResponseStorage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusResponseStorage(val *StatusResponseStorage) *NullableStatusResponseStorage {
	return &NullableStatusResponseStorage{value: val, isSet: true}
}

func (v NullableStatusResponseStorage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusResponseStorage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

