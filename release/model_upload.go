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
	"bytes"
	"fmt"
)

// checks if the Upload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Upload{}

// Upload Serializer for chunked uploads.
type Upload struct {
	// The size of the upload in bytes.
	Size int64 `json:"size"`
}

type _Upload Upload

// NewUpload instantiates a new Upload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpload(size int64) *Upload {
	this := Upload{}
	this.Size = size
	return &this
}

// NewUploadWithDefaults instantiates a new Upload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadWithDefaults() *Upload {
	this := Upload{}
	return &this
}

// GetSize returns the Size field value
func (o *Upload) GetSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *Upload) GetSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *Upload) SetSize(v int64) {
	o.Size = v
}

func (o Upload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Upload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["size"] = o.Size
	return toSerialize, nil
}

func (o *Upload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"size",
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

	varUpload := _Upload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpload)

	if err != nil {
		return err
	}

	*o = Upload(varUpload)

	return err
}

type NullableUpload struct {
	value *Upload
	isSet bool
}

func (v NullableUpload) Get() *Upload {
	return v.value
}

func (v *NullableUpload) Set(val *Upload) {
	v.value = val
	v.isSet = true
}

func (v NullableUpload) IsSet() bool {
	return v.isSet
}

func (v *NullableUpload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpload(val *Upload) *NullableUpload {
	return &NullableUpload{value: val, isSet: true}
}

func (v NullableUpload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


