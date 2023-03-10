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

// checks if the StatusResponseRedisConnection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusResponseRedisConnection{}

// StatusResponseRedisConnection Redis connection information
type StatusResponseRedisConnection struct {
	// Info about whether the app can connect to Redis
	Connected bool `json:"connected"`
}

// NewStatusResponseRedisConnection instantiates a new StatusResponseRedisConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusResponseRedisConnection(connected bool) *StatusResponseRedisConnection {
	this := StatusResponseRedisConnection{}
	this.Connected = connected
	return &this
}

// NewStatusResponseRedisConnectionWithDefaults instantiates a new StatusResponseRedisConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusResponseRedisConnectionWithDefaults() *StatusResponseRedisConnection {
	this := StatusResponseRedisConnection{}
	return &this
}

// GetConnected returns the Connected field value
func (o *StatusResponseRedisConnection) GetConnected() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Connected
}

// GetConnectedOk returns a tuple with the Connected field value
// and a boolean to check if the value has been set.
func (o *StatusResponseRedisConnection) GetConnectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Connected, true
}

// SetConnected sets field value
func (o *StatusResponseRedisConnection) SetConnected(v bool) {
	o.Connected = v
}

func (o StatusResponseRedisConnection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatusResponseRedisConnection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connected"] = o.Connected
	return toSerialize, nil
}

type NullableStatusResponseRedisConnection struct {
	value *StatusResponseRedisConnection
	isSet bool
}

func (v NullableStatusResponseRedisConnection) Get() *StatusResponseRedisConnection {
	return v.value
}

func (v *NullableStatusResponseRedisConnection) Set(val *StatusResponseRedisConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusResponseRedisConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusResponseRedisConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusResponseRedisConnection(val *StatusResponseRedisConnection) *NullableStatusResponseRedisConnection {
	return &NullableStatusResponseRedisConnection{value: val, isSet: true}
}

func (v NullableStatusResponseRedisConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusResponseRedisConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


