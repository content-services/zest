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

// CompressionTypeEnum * `zstd` - zstd* `gz` - gz
type CompressionTypeEnum string

// List of CompressionTypeEnum
const (
	COMPRESSIONTYPEENUM_ZSTD CompressionTypeEnum = "zstd"
	COMPRESSIONTYPEENUM_GZ CompressionTypeEnum = "gz"
)

// All allowed values of CompressionTypeEnum enum
var AllowedCompressionTypeEnumEnumValues = []CompressionTypeEnum{
	"zstd",
	"gz",
}

func (v *CompressionTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CompressionTypeEnum(value)
	for _, existing := range AllowedCompressionTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CompressionTypeEnum", value)
}

// NewCompressionTypeEnumFromValue returns a pointer to a valid CompressionTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCompressionTypeEnumFromValue(v string) (*CompressionTypeEnum, error) {
	ev := CompressionTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CompressionTypeEnum: valid values are %v", v, AllowedCompressionTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CompressionTypeEnum) IsValid() bool {
	for _, existing := range AllowedCompressionTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CompressionTypeEnum value
func (v CompressionTypeEnum) Ptr() *CompressionTypeEnum {
	return &v
}

type NullableCompressionTypeEnum struct {
	value *CompressionTypeEnum
	isSet bool
}

func (v NullableCompressionTypeEnum) Get() *CompressionTypeEnum {
	return v.value
}

func (v *NullableCompressionTypeEnum) Set(val *CompressionTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableCompressionTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableCompressionTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompressionTypeEnum(val *CompressionTypeEnum) *NullableCompressionTypeEnum {
	return &NullableCompressionTypeEnum{value: val, isSet: true}
}

func (v NullableCompressionTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompressionTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

