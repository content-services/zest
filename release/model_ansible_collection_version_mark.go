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
)

// checks if the AnsibleCollectionVersionMark type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnsibleCollectionVersionMark{}

// AnsibleCollectionVersionMark A serializer for mark models.
type AnsibleCollectionVersionMark struct {
	// The content this mark is pointing to.
	MarkedCollection string `json:"marked_collection"`
	// The string value of this mark.
	Value string `json:"value"`
}

// NewAnsibleCollectionVersionMark instantiates a new AnsibleCollectionVersionMark object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnsibleCollectionVersionMark(markedCollection string, value string) *AnsibleCollectionVersionMark {
	this := AnsibleCollectionVersionMark{}
	this.MarkedCollection = markedCollection
	this.Value = value
	return &this
}

// NewAnsibleCollectionVersionMarkWithDefaults instantiates a new AnsibleCollectionVersionMark object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnsibleCollectionVersionMarkWithDefaults() *AnsibleCollectionVersionMark {
	this := AnsibleCollectionVersionMark{}
	return &this
}

// GetMarkedCollection returns the MarkedCollection field value
func (o *AnsibleCollectionVersionMark) GetMarkedCollection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarkedCollection
}

// GetMarkedCollectionOk returns a tuple with the MarkedCollection field value
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionVersionMark) GetMarkedCollectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarkedCollection, true
}

// SetMarkedCollection sets field value
func (o *AnsibleCollectionVersionMark) SetMarkedCollection(v string) {
	o.MarkedCollection = v
}

// GetValue returns the Value field value
func (o *AnsibleCollectionVersionMark) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *AnsibleCollectionVersionMark) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *AnsibleCollectionVersionMark) SetValue(v string) {
	o.Value = v
}

func (o AnsibleCollectionVersionMark) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnsibleCollectionVersionMark) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["marked_collection"] = o.MarkedCollection
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableAnsibleCollectionVersionMark struct {
	value *AnsibleCollectionVersionMark
	isSet bool
}

func (v NullableAnsibleCollectionVersionMark) Get() *AnsibleCollectionVersionMark {
	return v.value
}

func (v *NullableAnsibleCollectionVersionMark) Set(val *AnsibleCollectionVersionMark) {
	v.value = val
	v.isSet = true
}

func (v NullableAnsibleCollectionVersionMark) IsSet() bool {
	return v.isSet
}

func (v *NullableAnsibleCollectionVersionMark) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnsibleCollectionVersionMark(val *AnsibleCollectionVersionMark) *NullableAnsibleCollectionVersionMark {
	return &NullableAnsibleCollectionVersionMark{value: val, isSet: true}
}

func (v NullableAnsibleCollectionVersionMark) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnsibleCollectionVersionMark) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


