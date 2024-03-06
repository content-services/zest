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

// checks if the RedisConnectionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RedisConnectionResponse{}

// RedisConnectionResponse Serializer for information about the Redis connection
type RedisConnectionResponse struct {
	// Info about whether the app can connect to Redis
	Connected bool `json:"connected"`
	AdditionalProperties map[string]interface{}
}

type _RedisConnectionResponse RedisConnectionResponse

// NewRedisConnectionResponse instantiates a new RedisConnectionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedisConnectionResponse(connected bool) *RedisConnectionResponse {
	this := RedisConnectionResponse{}
	this.Connected = connected
	return &this
}

// NewRedisConnectionResponseWithDefaults instantiates a new RedisConnectionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedisConnectionResponseWithDefaults() *RedisConnectionResponse {
	this := RedisConnectionResponse{}
	return &this
}

// GetConnected returns the Connected field value
func (o *RedisConnectionResponse) GetConnected() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Connected
}

// GetConnectedOk returns a tuple with the Connected field value
// and a boolean to check if the value has been set.
func (o *RedisConnectionResponse) GetConnectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Connected, true
}

// SetConnected sets field value
func (o *RedisConnectionResponse) SetConnected(v bool) {
	o.Connected = v
}

func (o RedisConnectionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RedisConnectionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["connected"] = o.Connected

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RedisConnectionResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"connected",
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

	varRedisConnectionResponse := _RedisConnectionResponse{}

	err = json.Unmarshal(data, &varRedisConnectionResponse)

	if err != nil {
		return err
	}

	*o = RedisConnectionResponse(varRedisConnectionResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "connected")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRedisConnectionResponse struct {
	value *RedisConnectionResponse
	isSet bool
}

func (v NullableRedisConnectionResponse) Get() *RedisConnectionResponse {
	return v.value
}

func (v *NullableRedisConnectionResponse) Set(val *RedisConnectionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRedisConnectionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRedisConnectionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedisConnectionResponse(val *RedisConnectionResponse) *NullableRedisConnectionResponse {
	return &NullableRedisConnectionResponse{value: val, isSet: true}
}

func (v NullableRedisConnectionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedisConnectionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


