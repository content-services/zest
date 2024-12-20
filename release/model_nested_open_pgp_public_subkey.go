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
	"time"
	"fmt"
)

// checks if the NestedOpenPGPPublicSubkey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedOpenPGPPublicSubkey{}

// NestedOpenPGPPublicSubkey Base serializer for use with [pulpcore.app.models.Model][]This ensures that all Serializers provide values for the 'pulp_href` field.The class provides a default for the ``ref_name`` attribute in theModelSerializers's ``Meta`` class. This ensures that the OpenAPI definitionsof plugins are namespaced properly.
type NestedOpenPGPPublicSubkey struct {
	Fingerprint string `json:"fingerprint"`
	Created time.Time `json:"created"`
	AdditionalProperties map[string]interface{}
}

type _NestedOpenPGPPublicSubkey NestedOpenPGPPublicSubkey

// NewNestedOpenPGPPublicSubkey instantiates a new NestedOpenPGPPublicSubkey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedOpenPGPPublicSubkey(fingerprint string, created time.Time) *NestedOpenPGPPublicSubkey {
	this := NestedOpenPGPPublicSubkey{}
	this.Fingerprint = fingerprint
	this.Created = created
	return &this
}

// NewNestedOpenPGPPublicSubkeyWithDefaults instantiates a new NestedOpenPGPPublicSubkey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedOpenPGPPublicSubkeyWithDefaults() *NestedOpenPGPPublicSubkey {
	this := NestedOpenPGPPublicSubkey{}
	return &this
}

// GetFingerprint returns the Fingerprint field value
func (o *NestedOpenPGPPublicSubkey) GetFingerprint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value
// and a boolean to check if the value has been set.
func (o *NestedOpenPGPPublicSubkey) GetFingerprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fingerprint, true
}

// SetFingerprint sets field value
func (o *NestedOpenPGPPublicSubkey) SetFingerprint(v string) {
	o.Fingerprint = v
}

// GetCreated returns the Created field value
func (o *NestedOpenPGPPublicSubkey) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *NestedOpenPGPPublicSubkey) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *NestedOpenPGPPublicSubkey) SetCreated(v time.Time) {
	o.Created = v
}

func (o NestedOpenPGPPublicSubkey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedOpenPGPPublicSubkey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fingerprint"] = o.Fingerprint
	toSerialize["created"] = o.Created

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedOpenPGPPublicSubkey) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fingerprint",
		"created",
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

	varNestedOpenPGPPublicSubkey := _NestedOpenPGPPublicSubkey{}

	err = json.Unmarshal(data, &varNestedOpenPGPPublicSubkey)

	if err != nil {
		return err
	}

	*o = NestedOpenPGPPublicSubkey(varNestedOpenPGPPublicSubkey)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "fingerprint")
		delete(additionalProperties, "created")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedOpenPGPPublicSubkey struct {
	value *NestedOpenPGPPublicSubkey
	isSet bool
}

func (v NullableNestedOpenPGPPublicSubkey) Get() *NestedOpenPGPPublicSubkey {
	return v.value
}

func (v *NullableNestedOpenPGPPublicSubkey) Set(val *NestedOpenPGPPublicSubkey) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedOpenPGPPublicSubkey) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedOpenPGPPublicSubkey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedOpenPGPPublicSubkey(val *NestedOpenPGPPublicSubkey) *NullableNestedOpenPGPPublicSubkey {
	return &NullableNestedOpenPGPPublicSubkey{value: val, isSet: true}
}

func (v NullableNestedOpenPGPPublicSubkey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedOpenPGPPublicSubkey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


