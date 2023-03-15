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

// checks if the TagResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagResponse{}

// TagResponse A serializer for the Tag model.
type TagResponse struct {
	Name *string `json:"name,omitempty"`
}

// NewTagResponse instantiates a new TagResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagResponse() *TagResponse {
	this := TagResponse{}
	return &this
}

// NewTagResponseWithDefaults instantiates a new TagResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagResponseWithDefaults() *TagResponse {
	this := TagResponse{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TagResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TagResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TagResponse) SetName(v string) {
	o.Name = &v
}

func (o TagResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: name is readOnly
	return toSerialize, nil
}

type NullableTagResponse struct {
	value *TagResponse
	isSet bool
}

func (v NullableTagResponse) Get() *TagResponse {
	return v.value
}

func (v *NullableTagResponse) Set(val *TagResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTagResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTagResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagResponse(val *TagResponse) *NullableTagResponse {
	return &NullableTagResponse{value: val, isSet: true}
}

func (v NullableTagResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


