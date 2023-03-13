/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pulpGoBinding

import (
	"encoding/json"
)

// checks if the UploadChunkResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UploadChunkResponse{}

// UploadChunkResponse A mixin for validating unknown serializers' fields.
type UploadChunkResponse struct {
	Offset *int32 `json:"offset,omitempty"`
	Size *int32 `json:"size,omitempty"`
}

// NewUploadChunkResponse instantiates a new UploadChunkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadChunkResponse() *UploadChunkResponse {
	this := UploadChunkResponse{}
	return &this
}

// NewUploadChunkResponseWithDefaults instantiates a new UploadChunkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadChunkResponseWithDefaults() *UploadChunkResponse {
	this := UploadChunkResponse{}
	return &this
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *UploadChunkResponse) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadChunkResponse) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *UploadChunkResponse) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *UploadChunkResponse) SetOffset(v int32) {
	o.Offset = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *UploadChunkResponse) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UploadChunkResponse) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *UploadChunkResponse) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *UploadChunkResponse) SetSize(v int32) {
	o.Size = &v
}

func (o UploadChunkResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UploadChunkResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: offset is readOnly
	// skip: size is readOnly
	return toSerialize, nil
}

type NullableUploadChunkResponse struct {
	value *UploadChunkResponse
	isSet bool
}

func (v NullableUploadChunkResponse) Get() *UploadChunkResponse {
	return v.value
}

func (v *NullableUploadChunkResponse) Set(val *UploadChunkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadChunkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadChunkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadChunkResponse(val *UploadChunkResponse) *NullableUploadChunkResponse {
	return &NullableUploadChunkResponse{value: val, isSet: true}
}

func (v NullableUploadChunkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadChunkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

