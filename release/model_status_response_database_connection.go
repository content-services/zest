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

// checks if the StatusResponseDatabaseConnection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusResponseDatabaseConnection{}

// StatusResponseDatabaseConnection Database connection information
type StatusResponseDatabaseConnection struct {
	// Info about whether the app can connect to the database
	Connected bool `json:"connected"`
}

// NewStatusResponseDatabaseConnection instantiates a new StatusResponseDatabaseConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusResponseDatabaseConnection(connected bool) *StatusResponseDatabaseConnection {
	this := StatusResponseDatabaseConnection{}
	this.Connected = connected
	return &this
}

// NewStatusResponseDatabaseConnectionWithDefaults instantiates a new StatusResponseDatabaseConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusResponseDatabaseConnectionWithDefaults() *StatusResponseDatabaseConnection {
	this := StatusResponseDatabaseConnection{}
	return &this
}

// GetConnected returns the Connected field value
func (o *StatusResponseDatabaseConnection) GetConnected() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Connected
}

// GetConnectedOk returns a tuple with the Connected field value
// and a boolean to check if the value has been set.
func (o *StatusResponseDatabaseConnection) GetConnectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Connected, true
}

// SetConnected sets field value
func (o *StatusResponseDatabaseConnection) SetConnected(v bool) {
	o.Connected = v
}

func (o StatusResponseDatabaseConnection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatusResponseDatabaseConnection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connected"] = o.Connected
	return toSerialize, nil
}

type NullableStatusResponseDatabaseConnection struct {
	value *StatusResponseDatabaseConnection
	isSet bool
}

func (v NullableStatusResponseDatabaseConnection) Get() *StatusResponseDatabaseConnection {
	return v.value
}

func (v *NullableStatusResponseDatabaseConnection) Set(val *StatusResponseDatabaseConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusResponseDatabaseConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusResponseDatabaseConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusResponseDatabaseConnection(val *StatusResponseDatabaseConnection) *NullableStatusResponseDatabaseConnection {
	return &NullableStatusResponseDatabaseConnection{value: val, isSet: true}
}

func (v NullableStatusResponseDatabaseConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusResponseDatabaseConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


