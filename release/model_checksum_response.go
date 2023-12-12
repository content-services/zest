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

// checks if the ChecksumResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChecksumResponse{}

// ChecksumResponse Checksum serializer.
type ChecksumResponse struct {
	// File path.
	Path string `json:"path"`
	// Checksum for the file.
	Checksum string `json:"checksum"`
}

type _ChecksumResponse ChecksumResponse

// NewChecksumResponse instantiates a new ChecksumResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChecksumResponse(path string, checksum string) *ChecksumResponse {
	this := ChecksumResponse{}
	this.Path = path
	this.Checksum = checksum
	return &this
}

// NewChecksumResponseWithDefaults instantiates a new ChecksumResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChecksumResponseWithDefaults() *ChecksumResponse {
	this := ChecksumResponse{}
	return &this
}

// GetPath returns the Path field value
func (o *ChecksumResponse) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *ChecksumResponse) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *ChecksumResponse) SetPath(v string) {
	o.Path = v
}

// GetChecksum returns the Checksum field value
func (o *ChecksumResponse) GetChecksum() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value
// and a boolean to check if the value has been set.
func (o *ChecksumResponse) GetChecksumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Checksum, true
}

// SetChecksum sets field value
func (o *ChecksumResponse) SetChecksum(v string) {
	o.Checksum = v
}

func (o ChecksumResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChecksumResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["path"] = o.Path
	toSerialize["checksum"] = o.Checksum
	return toSerialize, nil
}

func (o *ChecksumResponse) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"path",
		"checksum",
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

	varChecksumResponse := _ChecksumResponse{}

	err = json.Unmarshal(bytes, &varChecksumResponse)

	if err != nil {
		return err
	}

	*o = ChecksumResponse(varChecksumResponse)

	return err
}

type NullableChecksumResponse struct {
	value *ChecksumResponse
	isSet bool
}

func (v NullableChecksumResponse) Get() *ChecksumResponse {
	return v.value
}

func (v *NullableChecksumResponse) Set(val *ChecksumResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableChecksumResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableChecksumResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChecksumResponse(val *ChecksumResponse) *NullableChecksumResponse {
	return &NullableChecksumResponse{value: val, isSet: true}
}

func (v NullableChecksumResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChecksumResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


