/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// StatesEnum the model 'StatesEnum'
type StatesEnum string

// List of StatesEnum
const (
	STATESENUM_SKIPPED StatesEnum = "skipped"
	STATESENUM_COMPLETED StatesEnum = "completed"
	STATESENUM_FAILED StatesEnum = "failed"
	STATESENUM_CANCELED StatesEnum = "canceled"
)

// All allowed values of StatesEnum enum
var AllowedStatesEnumEnumValues = []StatesEnum{
	"skipped",
	"completed",
	"failed",
	"canceled",
}

func (v *StatesEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StatesEnum(value)
	for _, existing := range AllowedStatesEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StatesEnum", value)
}

// NewStatesEnumFromValue returns a pointer to a valid StatesEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStatesEnumFromValue(v string) (*StatesEnum, error) {
	ev := StatesEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StatesEnum: valid values are %v", v, AllowedStatesEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StatesEnum) IsValid() bool {
	for _, existing := range AllowedStatesEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StatesEnum value
func (v StatesEnum) Ptr() *StatesEnum {
	return &v
}

type NullableStatesEnum struct {
	value *StatesEnum
	isSet bool
}

func (v NullableStatesEnum) Get() *StatesEnum {
	return v.value
}

func (v *NullableStatesEnum) Set(val *StatesEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableStatesEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableStatesEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatesEnum(val *StatesEnum) *NullableStatesEnum {
	return &NullableStatesEnum{value: val, isSet: true}
}

func (v NullableStatesEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatesEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

