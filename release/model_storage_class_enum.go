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

// StorageClassEnum the model 'StorageClassEnum'
type StorageClassEnum string

// List of StorageClassEnum
const (
	STORAGECLASSENUM_PULPCORE_APP_MODELS_STORAGE_FILE_SYSTEM StorageClassEnum = "pulpcore.app.models.storage.FileSystem"
	STORAGECLASSENUM_STORAGES_BACKENDS_S3BOTO3_S3_BOTO3_STORAGE StorageClassEnum = "storages.backends.s3boto3.S3Boto3Storage"
	STORAGECLASSENUM_STORAGES_BACKENDS_AZURE_STORAGE_AZURE_STORAGE StorageClassEnum = "storages.backends.azure_storage.AzureStorage"
)

// All allowed values of StorageClassEnum enum
var AllowedStorageClassEnumEnumValues = []StorageClassEnum{
	"pulpcore.app.models.storage.FileSystem",
	"storages.backends.s3boto3.S3Boto3Storage",
	"storages.backends.azure_storage.AzureStorage",
}

func (v *StorageClassEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StorageClassEnum(value)
	for _, existing := range AllowedStorageClassEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StorageClassEnum", value)
}

// NewStorageClassEnumFromValue returns a pointer to a valid StorageClassEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStorageClassEnumFromValue(v string) (*StorageClassEnum, error) {
	ev := StorageClassEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StorageClassEnum: valid values are %v", v, AllowedStorageClassEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StorageClassEnum) IsValid() bool {
	for _, existing := range AllowedStorageClassEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StorageClassEnum value
func (v StorageClassEnum) Ptr() *StorageClassEnum {
	return &v
}

type NullableStorageClassEnum struct {
	value *StorageClassEnum
	isSet bool
}

func (v NullableStorageClassEnum) Get() *StorageClassEnum {
	return v.value
}

func (v *NullableStorageClassEnum) Set(val *StorageClassEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageClassEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageClassEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageClassEnum(val *StorageClassEnum) *NullableStorageClassEnum {
	return &NullableStorageClassEnum{value: val, isSet: true}
}

func (v NullableStorageClassEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageClassEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

