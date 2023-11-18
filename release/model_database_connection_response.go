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

// checks if the DatabaseConnectionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseConnectionResponse{}

// DatabaseConnectionResponse Serializer for the database connection information
type DatabaseConnectionResponse struct {
	// Info about whether the app can connect to the database
	Connected bool `json:"connected"`
}

type _DatabaseConnectionResponse DatabaseConnectionResponse

// NewDatabaseConnectionResponse instantiates a new DatabaseConnectionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseConnectionResponse(connected bool) *DatabaseConnectionResponse {
	this := DatabaseConnectionResponse{}
	this.Connected = connected
	return &this
}

// NewDatabaseConnectionResponseWithDefaults instantiates a new DatabaseConnectionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseConnectionResponseWithDefaults() *DatabaseConnectionResponse {
	this := DatabaseConnectionResponse{}
	return &this
}

// GetConnected returns the Connected field value
func (o *DatabaseConnectionResponse) GetConnected() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Connected
}

// GetConnectedOk returns a tuple with the Connected field value
// and a boolean to check if the value has been set.
func (o *DatabaseConnectionResponse) GetConnectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Connected, true
}

// SetConnected sets field value
func (o *DatabaseConnectionResponse) SetConnected(v bool) {
	o.Connected = v
}

func (o DatabaseConnectionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatabaseConnectionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connected"] = o.Connected
	return toSerialize, nil
}

func (o *DatabaseConnectionResponse) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"connected",
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

	varDatabaseConnectionResponse := _DatabaseConnectionResponse{}

	err = json.Unmarshal(bytes, &varDatabaseConnectionResponse)

	if err != nil {
		return err
	}

	*o = DatabaseConnectionResponse(varDatabaseConnectionResponse)

	return err
}

type NullableDatabaseConnectionResponse struct {
	value *DatabaseConnectionResponse
	isSet bool
}

func (v NullableDatabaseConnectionResponse) Get() *DatabaseConnectionResponse {
	return v.value
}

func (v *NullableDatabaseConnectionResponse) Set(val *DatabaseConnectionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseConnectionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseConnectionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseConnectionResponse(val *DatabaseConnectionResponse) *NullableDatabaseConnectionResponse {
	return &NullableDatabaseConnectionResponse{value: val, isSet: true}
}

func (v NullableDatabaseConnectionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseConnectionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


