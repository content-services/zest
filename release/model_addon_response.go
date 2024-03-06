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

// checks if the AddonResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddonResponse{}

// AddonResponse Addon serializer.
type AddonResponse struct {
	// Addon id.
	AddonId string `json:"addon_id"`
	// Addon uid.
	Uid string `json:"uid"`
	// Addon name.
	Name string `json:"name"`
	// Addon type.
	Type string `json:"type"`
	// Relative path to directory with binary RPMs.
	Packages string `json:"packages"`
	AdditionalProperties map[string]interface{}
}

type _AddonResponse AddonResponse

// NewAddonResponse instantiates a new AddonResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddonResponse(addonId string, uid string, name string, type_ string, packages string) *AddonResponse {
	this := AddonResponse{}
	this.AddonId = addonId
	this.Uid = uid
	this.Name = name
	this.Type = type_
	this.Packages = packages
	return &this
}

// NewAddonResponseWithDefaults instantiates a new AddonResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddonResponseWithDefaults() *AddonResponse {
	this := AddonResponse{}
	return &this
}

// GetAddonId returns the AddonId field value
func (o *AddonResponse) GetAddonId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddonId
}

// GetAddonIdOk returns a tuple with the AddonId field value
// and a boolean to check if the value has been set.
func (o *AddonResponse) GetAddonIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddonId, true
}

// SetAddonId sets field value
func (o *AddonResponse) SetAddonId(v string) {
	o.AddonId = v
}

// GetUid returns the Uid field value
func (o *AddonResponse) GetUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uid
}

// GetUidOk returns a tuple with the Uid field value
// and a boolean to check if the value has been set.
func (o *AddonResponse) GetUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uid, true
}

// SetUid sets field value
func (o *AddonResponse) SetUid(v string) {
	o.Uid = v
}

// GetName returns the Name field value
func (o *AddonResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AddonResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AddonResponse) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *AddonResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AddonResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AddonResponse) SetType(v string) {
	o.Type = v
}

// GetPackages returns the Packages field value
func (o *AddonResponse) GetPackages() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Packages
}

// GetPackagesOk returns a tuple with the Packages field value
// and a boolean to check if the value has been set.
func (o *AddonResponse) GetPackagesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Packages, true
}

// SetPackages sets field value
func (o *AddonResponse) SetPackages(v string) {
	o.Packages = v
}

func (o AddonResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddonResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["addon_id"] = o.AddonId
	toSerialize["uid"] = o.Uid
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	toSerialize["packages"] = o.Packages

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AddonResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"addon_id",
		"uid",
		"name",
		"type",
		"packages",
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

	varAddonResponse := _AddonResponse{}

	err = json.Unmarshal(data, &varAddonResponse)

	if err != nil {
		return err
	}

	*o = AddonResponse(varAddonResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "addon_id")
		delete(additionalProperties, "uid")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "packages")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAddonResponse struct {
	value *AddonResponse
	isSet bool
}

func (v NullableAddonResponse) Get() *AddonResponse {
	return v.value
}

func (v *NullableAddonResponse) Set(val *AddonResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAddonResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAddonResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddonResponse(val *AddonResponse) *NullableAddonResponse {
	return &NullableAddonResponse{value: val, isSet: true}
}

func (v NullableAddonResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddonResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


