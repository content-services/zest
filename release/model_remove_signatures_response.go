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

// checks if the RemoveSignaturesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoveSignaturesResponse{}

// RemoveSignaturesResponse A serializer for parsing and validating data associated with the signatures removal.
type RemoveSignaturesResponse struct {
	// key_id of the key the signatures were produced with
	SignedWithKeyId string `json:"signed_with_key_id"`
	AdditionalProperties map[string]interface{}
}

type _RemoveSignaturesResponse RemoveSignaturesResponse

// NewRemoveSignaturesResponse instantiates a new RemoveSignaturesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveSignaturesResponse(signedWithKeyId string) *RemoveSignaturesResponse {
	this := RemoveSignaturesResponse{}
	this.SignedWithKeyId = signedWithKeyId
	return &this
}

// NewRemoveSignaturesResponseWithDefaults instantiates a new RemoveSignaturesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveSignaturesResponseWithDefaults() *RemoveSignaturesResponse {
	this := RemoveSignaturesResponse{}
	return &this
}

// GetSignedWithKeyId returns the SignedWithKeyId field value
func (o *RemoveSignaturesResponse) GetSignedWithKeyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignedWithKeyId
}

// GetSignedWithKeyIdOk returns a tuple with the SignedWithKeyId field value
// and a boolean to check if the value has been set.
func (o *RemoveSignaturesResponse) GetSignedWithKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignedWithKeyId, true
}

// SetSignedWithKeyId sets field value
func (o *RemoveSignaturesResponse) SetSignedWithKeyId(v string) {
	o.SignedWithKeyId = v
}

func (o RemoveSignaturesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoveSignaturesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["signed_with_key_id"] = o.SignedWithKeyId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RemoveSignaturesResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"signed_with_key_id",
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

	varRemoveSignaturesResponse := _RemoveSignaturesResponse{}

	err = json.Unmarshal(data, &varRemoveSignaturesResponse)

	if err != nil {
		return err
	}

	*o = RemoveSignaturesResponse(varRemoveSignaturesResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "signed_with_key_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRemoveSignaturesResponse struct {
	value *RemoveSignaturesResponse
	isSet bool
}

func (v NullableRemoveSignaturesResponse) Get() *RemoveSignaturesResponse {
	return v.value
}

func (v *NullableRemoveSignaturesResponse) Set(val *RemoveSignaturesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveSignaturesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveSignaturesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveSignaturesResponse(val *RemoveSignaturesResponse) *NullableRemoveSignaturesResponse {
	return &NullableRemoveSignaturesResponse{value: val, isSet: true}
}

func (v NullableRemoveSignaturesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveSignaturesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


