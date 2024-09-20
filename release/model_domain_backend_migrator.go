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

// checks if the DomainBackendMigrator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainBackendMigrator{}

// DomainBackendMigrator Special serializer for performing a storage backend migration on a Domain.
type DomainBackendMigrator struct {
	// The new backend storage class to migrate to.* `pulpcore.app.models.storage.FileSystem` - Use local filesystem as storage* `storages.backends.s3boto3.S3Boto3Storage` - Use Amazon S3 as storage* `storages.backends.azure_storage.AzureStorage` - Use Azure Blob as storage
	StorageClass StorageClassEnum `json:"storage_class"`
	// The settings for the new storage class to migrate to.
	StorageSettings map[string]interface{} `json:"storage_settings"`
	AdditionalProperties map[string]interface{}
}

type _DomainBackendMigrator DomainBackendMigrator

// NewDomainBackendMigrator instantiates a new DomainBackendMigrator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainBackendMigrator(storageClass StorageClassEnum, storageSettings map[string]interface{}) *DomainBackendMigrator {
	this := DomainBackendMigrator{}
	this.StorageClass = storageClass
	this.StorageSettings = storageSettings
	return &this
}

// NewDomainBackendMigratorWithDefaults instantiates a new DomainBackendMigrator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainBackendMigratorWithDefaults() *DomainBackendMigrator {
	this := DomainBackendMigrator{}
	return &this
}

// GetStorageClass returns the StorageClass field value
func (o *DomainBackendMigrator) GetStorageClass() StorageClassEnum {
	if o == nil {
		var ret StorageClassEnum
		return ret
	}

	return o.StorageClass
}

// GetStorageClassOk returns a tuple with the StorageClass field value
// and a boolean to check if the value has been set.
func (o *DomainBackendMigrator) GetStorageClassOk() (*StorageClassEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageClass, true
}

// SetStorageClass sets field value
func (o *DomainBackendMigrator) SetStorageClass(v StorageClassEnum) {
	o.StorageClass = v
}

// GetStorageSettings returns the StorageSettings field value
func (o *DomainBackendMigrator) GetStorageSettings() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.StorageSettings
}

// GetStorageSettingsOk returns a tuple with the StorageSettings field value
// and a boolean to check if the value has been set.
func (o *DomainBackendMigrator) GetStorageSettingsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.StorageSettings, true
}

// SetStorageSettings sets field value
func (o *DomainBackendMigrator) SetStorageSettings(v map[string]interface{}) {
	o.StorageSettings = v
}

func (o DomainBackendMigrator) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainBackendMigrator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["storage_class"] = o.StorageClass
	toSerialize["storage_settings"] = o.StorageSettings

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DomainBackendMigrator) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"storage_class",
		"storage_settings",
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

	varDomainBackendMigrator := _DomainBackendMigrator{}

	err = json.Unmarshal(data, &varDomainBackendMigrator)

	if err != nil {
		return err
	}

	*o = DomainBackendMigrator(varDomainBackendMigrator)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "storage_class")
		delete(additionalProperties, "storage_settings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDomainBackendMigrator struct {
	value *DomainBackendMigrator
	isSet bool
}

func (v NullableDomainBackendMigrator) Get() *DomainBackendMigrator {
	return v.value
}

func (v *NullableDomainBackendMigrator) Set(val *DomainBackendMigrator) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainBackendMigrator) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainBackendMigrator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainBackendMigrator(val *DomainBackendMigrator) *NullableDomainBackendMigrator {
	return &NullableDomainBackendMigrator{value: val, isSet: true}
}

func (v NullableDomainBackendMigrator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainBackendMigrator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

