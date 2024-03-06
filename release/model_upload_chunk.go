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
	"os"
	"fmt"
)

// checks if the UploadChunk type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UploadChunk{}

// UploadChunk A mixin for validating unknown serializers' fields.
type UploadChunk struct {
	// A chunk of the uploaded file.
	File *os.File `json:"file"`
	// The SHA-256 checksum of the chunk if available.
	Sha256 NullableString `json:"sha256,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UploadChunk UploadChunk

// NewUploadChunk instantiates a new UploadChunk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadChunk(file *os.File) *UploadChunk {
	this := UploadChunk{}
	this.File = file
	return &this
}

// NewUploadChunkWithDefaults instantiates a new UploadChunk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadChunkWithDefaults() *UploadChunk {
	this := UploadChunk{}
	return &this
}

// GetFile returns the File field value
func (o *UploadChunk) GetFile() *os.File {
	if o == nil {
		var ret *os.File
		return ret
	}

	return o.File
}

// GetFileOk returns a tuple with the File field value
// and a boolean to check if the value has been set.
func (o *UploadChunk) GetFileOk() (**os.File, bool) {
	if o == nil {
		return nil, false
	}
	return &o.File, true
}

// SetFile sets field value
func (o *UploadChunk) SetFile(v *os.File) {
	o.File = v
}

// GetSha256 returns the Sha256 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UploadChunk) GetSha256() string {
	if o == nil || IsNil(o.Sha256.Get()) {
		var ret string
		return ret
	}
	return *o.Sha256.Get()
}

// GetSha256Ok returns a tuple with the Sha256 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UploadChunk) GetSha256Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sha256.Get(), o.Sha256.IsSet()
}

// HasSha256 returns a boolean if a field has been set.
func (o *UploadChunk) HasSha256() bool {
	if o != nil && o.Sha256.IsSet() {
		return true
	}

	return false
}

// SetSha256 gets a reference to the given NullableString and assigns it to the Sha256 field.
func (o *UploadChunk) SetSha256(v string) {
	o.Sha256.Set(&v)
}
// SetSha256Nil sets the value for Sha256 to be an explicit nil
func (o *UploadChunk) SetSha256Nil() {
	o.Sha256.Set(nil)
}

// UnsetSha256 ensures that no value is present for Sha256, not even an explicit nil
func (o *UploadChunk) UnsetSha256() {
	o.Sha256.Unset()
}

func (o UploadChunk) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UploadChunk) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["file"] = o.File
	if o.Sha256.IsSet() {
		toSerialize["sha256"] = o.Sha256.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UploadChunk) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"file",
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

	varUploadChunk := _UploadChunk{}

	err = json.Unmarshal(data, &varUploadChunk)

	if err != nil {
		return err
	}

	*o = UploadChunk(varUploadChunk)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "file")
		delete(additionalProperties, "sha256")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUploadChunk struct {
	value *UploadChunk
	isSet bool
}

func (v NullableUploadChunk) Get() *UploadChunk {
	return v.value
}

func (v *NullableUploadChunk) Set(val *UploadChunk) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadChunk) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadChunk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadChunk(val *UploadChunk) *NullableUploadChunk {
	return &NullableUploadChunk{value: val, isSet: true}
}

func (v NullableUploadChunk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadChunk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


