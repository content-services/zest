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

// RemoteResponsePolicyEnum * `immediate` - immediate * `When syncing, download all metadata and content now.` - When syncing, download all metadata and content now.
type RemoteResponsePolicyEnum string

// List of RemoteResponsePolicyEnum
const (
	REMOTERESPONSEPOLICYENUM_IMMEDIATE RemoteResponsePolicyEnum = "immediate"
	REMOTERESPONSEPOLICYENUM_WHEN_SYNCING_DOWNLOAD_ALL_METADATA_AND_CONTENT_NOW RemoteResponsePolicyEnum = "When syncing, download all metadata and content now."
)

// All allowed values of RemoteResponsePolicyEnum enum
var AllowedRemoteResponsePolicyEnumEnumValues = []RemoteResponsePolicyEnum{
	"immediate",
	"When syncing, download all metadata and content now.",
}

func (v *RemoteResponsePolicyEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RemoteResponsePolicyEnum(value)
	for _, existing := range AllowedRemoteResponsePolicyEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RemoteResponsePolicyEnum", value)
}

// NewRemoteResponsePolicyEnumFromValue returns a pointer to a valid RemoteResponsePolicyEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRemoteResponsePolicyEnumFromValue(v string) (*RemoteResponsePolicyEnum, error) {
	ev := RemoteResponsePolicyEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RemoteResponsePolicyEnum: valid values are %v", v, AllowedRemoteResponsePolicyEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RemoteResponsePolicyEnum) IsValid() bool {
	for _, existing := range AllowedRemoteResponsePolicyEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RemoteResponsePolicyEnum value
func (v RemoteResponsePolicyEnum) Ptr() *RemoteResponsePolicyEnum {
	return &v
}

type NullableRemoteResponsePolicyEnum struct {
	value *RemoteResponsePolicyEnum
	isSet bool
}

func (v NullableRemoteResponsePolicyEnum) Get() *RemoteResponsePolicyEnum {
	return v.value
}

func (v *NullableRemoteResponsePolicyEnum) Set(val *RemoteResponsePolicyEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteResponsePolicyEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteResponsePolicyEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteResponsePolicyEnum(val *RemoteResponsePolicyEnum) *NullableRemoteResponsePolicyEnum {
	return &NullableRemoteResponsePolicyEnum{value: val, isSet: true}
}

func (v NullableRemoteResponsePolicyEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteResponsePolicyEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
