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

// checks if the Domain type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Domain{}

// Domain Serializer for Domain.
type Domain struct {
	// A name for this domain.
	Name string `json:"name"`
	// An optional description.
	Description NullableString `json:"description,omitempty"`
	StorageClass StorageClassEnum `json:"storage_class"`
	// Settings for storage class.
	StorageSettings map[string]interface{} `json:"storage_settings"`
	// Boolean to have the content app redirect to object storage.
	RedirectToObjectStorage *bool `json:"redirect_to_object_storage,omitempty"`
	// Boolean to hide distributions with a content guard in the content app.
	HideGuardedDistributions *bool `json:"hide_guarded_distributions,omitempty"`
}

type _Domain Domain

// NewDomain instantiates a new Domain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomain(name string, storageClass StorageClassEnum, storageSettings map[string]interface{}) *Domain {
	this := Domain{}
	this.Name = name
	this.StorageClass = storageClass
	this.StorageSettings = storageSettings
	var redirectToObjectStorage bool = true
	this.RedirectToObjectStorage = &redirectToObjectStorage
	var hideGuardedDistributions bool = false
	this.HideGuardedDistributions = &hideGuardedDistributions
	return &this
}

// NewDomainWithDefaults instantiates a new Domain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainWithDefaults() *Domain {
	this := Domain{}
	var redirectToObjectStorage bool = true
	this.RedirectToObjectStorage = &redirectToObjectStorage
	var hideGuardedDistributions bool = false
	this.HideGuardedDistributions = &hideGuardedDistributions
	return &this
}

// GetName returns the Name field value
func (o *Domain) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Domain) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Domain) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Domain) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Domain) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Domain) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Domain) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Domain) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Domain) UnsetDescription() {
	o.Description.Unset()
}

// GetStorageClass returns the StorageClass field value
func (o *Domain) GetStorageClass() StorageClassEnum {
	if o == nil {
		var ret StorageClassEnum
		return ret
	}

	return o.StorageClass
}

// GetStorageClassOk returns a tuple with the StorageClass field value
// and a boolean to check if the value has been set.
func (o *Domain) GetStorageClassOk() (*StorageClassEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageClass, true
}

// SetStorageClass sets field value
func (o *Domain) SetStorageClass(v StorageClassEnum) {
	o.StorageClass = v
}

// GetStorageSettings returns the StorageSettings field value
func (o *Domain) GetStorageSettings() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.StorageSettings
}

// GetStorageSettingsOk returns a tuple with the StorageSettings field value
// and a boolean to check if the value has been set.
func (o *Domain) GetStorageSettingsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.StorageSettings, true
}

// SetStorageSettings sets field value
func (o *Domain) SetStorageSettings(v map[string]interface{}) {
	o.StorageSettings = v
}

// GetRedirectToObjectStorage returns the RedirectToObjectStorage field value if set, zero value otherwise.
func (o *Domain) GetRedirectToObjectStorage() bool {
	if o == nil || IsNil(o.RedirectToObjectStorage) {
		var ret bool
		return ret
	}
	return *o.RedirectToObjectStorage
}

// GetRedirectToObjectStorageOk returns a tuple with the RedirectToObjectStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetRedirectToObjectStorageOk() (*bool, bool) {
	if o == nil || IsNil(o.RedirectToObjectStorage) {
		return nil, false
	}
	return o.RedirectToObjectStorage, true
}

// HasRedirectToObjectStorage returns a boolean if a field has been set.
func (o *Domain) HasRedirectToObjectStorage() bool {
	if o != nil && !IsNil(o.RedirectToObjectStorage) {
		return true
	}

	return false
}

// SetRedirectToObjectStorage gets a reference to the given bool and assigns it to the RedirectToObjectStorage field.
func (o *Domain) SetRedirectToObjectStorage(v bool) {
	o.RedirectToObjectStorage = &v
}

// GetHideGuardedDistributions returns the HideGuardedDistributions field value if set, zero value otherwise.
func (o *Domain) GetHideGuardedDistributions() bool {
	if o == nil || IsNil(o.HideGuardedDistributions) {
		var ret bool
		return ret
	}
	return *o.HideGuardedDistributions
}

// GetHideGuardedDistributionsOk returns a tuple with the HideGuardedDistributions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetHideGuardedDistributionsOk() (*bool, bool) {
	if o == nil || IsNil(o.HideGuardedDistributions) {
		return nil, false
	}
	return o.HideGuardedDistributions, true
}

// HasHideGuardedDistributions returns a boolean if a field has been set.
func (o *Domain) HasHideGuardedDistributions() bool {
	if o != nil && !IsNil(o.HideGuardedDistributions) {
		return true
	}

	return false
}

// SetHideGuardedDistributions gets a reference to the given bool and assigns it to the HideGuardedDistributions field.
func (o *Domain) SetHideGuardedDistributions(v bool) {
	o.HideGuardedDistributions = &v
}

func (o Domain) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Domain) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["storage_class"] = o.StorageClass
	toSerialize["storage_settings"] = o.StorageSettings
	if !IsNil(o.RedirectToObjectStorage) {
		toSerialize["redirect_to_object_storage"] = o.RedirectToObjectStorage
	}
	if !IsNil(o.HideGuardedDistributions) {
		toSerialize["hide_guarded_distributions"] = o.HideGuardedDistributions
	}
	return toSerialize, nil
}

func (o *Domain) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"storage_class",
		"storage_settings",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDomain := _Domain{}

	err = json.Unmarshal(bytes, &varDomain)

	if err != nil {
		return err
	}

	*o = Domain(varDomain)

	return err
}

type NullableDomain struct {
	value *Domain
	isSet bool
}

func (v NullableDomain) Get() *Domain {
	return v.value
}

func (v *NullableDomain) Set(val *Domain) {
	v.value = val
	v.isSet = true
}

func (v NullableDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomain(val *Domain) *NullableDomain {
	return &NullableDomain{value: val, isSet: true}
}

func (v NullableDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


