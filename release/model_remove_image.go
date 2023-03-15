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

// checks if the RemoveImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoveImage{}

// RemoveImage A serializer for parsing and validating data associated with the image removal.
type RemoveImage struct {
	// sha256 of the Manifest file
	Digest string `json:"digest"`
}

// NewRemoveImage instantiates a new RemoveImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveImage(digest string) *RemoveImage {
	this := RemoveImage{}
	this.Digest = digest
	return &this
}

// NewRemoveImageWithDefaults instantiates a new RemoveImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveImageWithDefaults() *RemoveImage {
	this := RemoveImage{}
	return &this
}

// GetDigest returns the Digest field value
func (o *RemoveImage) GetDigest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Digest
}

// GetDigestOk returns a tuple with the Digest field value
// and a boolean to check if the value has been set.
func (o *RemoveImage) GetDigestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Digest, true
}

// SetDigest sets field value
func (o *RemoveImage) SetDigest(v string) {
	o.Digest = v
}

func (o RemoveImage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoveImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["digest"] = o.Digest
	return toSerialize, nil
}

type NullableRemoveImage struct {
	value *RemoveImage
	isSet bool
}

func (v NullableRemoveImage) Get() *RemoveImage {
	return v.value
}

func (v *NullableRemoveImage) Set(val *RemoveImage) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveImage) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveImage(val *RemoveImage) *NullableRemoveImage {
	return &NullableRemoveImage{value: val, isSet: true}
}

func (v NullableRemoveImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


