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

// checks if the RecursiveManage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecursiveManage{}

// RecursiveManage Serializer for adding and removing content to/from a Container repository.
type RecursiveManage struct {
	// A list of content units to operate on.
	ContentUnits []interface{} `json:"content_units,omitempty"`
}

// NewRecursiveManage instantiates a new RecursiveManage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecursiveManage() *RecursiveManage {
	this := RecursiveManage{}
	return &this
}

// NewRecursiveManageWithDefaults instantiates a new RecursiveManage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecursiveManageWithDefaults() *RecursiveManage {
	this := RecursiveManage{}
	return &this
}

// GetContentUnits returns the ContentUnits field value if set, zero value otherwise.
func (o *RecursiveManage) GetContentUnits() []interface{} {
	if o == nil || IsNil(o.ContentUnits) {
		var ret []interface{}
		return ret
	}
	return o.ContentUnits
}

// GetContentUnitsOk returns a tuple with the ContentUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecursiveManage) GetContentUnitsOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.ContentUnits) {
		return nil, false
	}
	return o.ContentUnits, true
}

// HasContentUnits returns a boolean if a field has been set.
func (o *RecursiveManage) HasContentUnits() bool {
	if o != nil && !IsNil(o.ContentUnits) {
		return true
	}

	return false
}

// SetContentUnits gets a reference to the given []interface{} and assigns it to the ContentUnits field.
func (o *RecursiveManage) SetContentUnits(v []interface{}) {
	o.ContentUnits = v
}

func (o RecursiveManage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecursiveManage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContentUnits) {
		toSerialize["content_units"] = o.ContentUnits
	}
	return toSerialize, nil
}

type NullableRecursiveManage struct {
	value *RecursiveManage
	isSet bool
}

func (v NullableRecursiveManage) Get() *RecursiveManage {
	return v.value
}

func (v *NullableRecursiveManage) Set(val *RecursiveManage) {
	v.value = val
	v.isSet = true
}

func (v NullableRecursiveManage) IsSet() bool {
	return v.isSet
}

func (v *NullableRecursiveManage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecursiveManage(val *RecursiveManage) *NullableRecursiveManage {
	return &NullableRecursiveManage{value: val, isSet: true}
}

func (v NullableRecursiveManage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecursiveManage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


