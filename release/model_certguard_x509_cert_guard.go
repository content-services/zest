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

// checks if the CertguardX509CertGuard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertguardX509CertGuard{}

// CertguardX509CertGuard X.509 Content Guard Serializer.
type CertguardX509CertGuard struct {
	// The unique name.
	Name string `json:"name"`
	// An optional description.
	Description NullableString `json:"description,omitempty"`
	// A Certificate Authority (CA) certificate (or a bundle thereof) used to verify client-certificate authenticity.
	CaCertificate string `json:"ca_certificate"`
}

type _CertguardX509CertGuard CertguardX509CertGuard

// NewCertguardX509CertGuard instantiates a new CertguardX509CertGuard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertguardX509CertGuard(name string, caCertificate string) *CertguardX509CertGuard {
	this := CertguardX509CertGuard{}
	this.Name = name
	this.CaCertificate = caCertificate
	return &this
}

// NewCertguardX509CertGuardWithDefaults instantiates a new CertguardX509CertGuard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertguardX509CertGuardWithDefaults() *CertguardX509CertGuard {
	this := CertguardX509CertGuard{}
	return &this
}

// GetName returns the Name field value
func (o *CertguardX509CertGuard) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CertguardX509CertGuard) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CertguardX509CertGuard) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertguardX509CertGuard) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertguardX509CertGuard) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CertguardX509CertGuard) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CertguardX509CertGuard) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CertguardX509CertGuard) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CertguardX509CertGuard) UnsetDescription() {
	o.Description.Unset()
}

// GetCaCertificate returns the CaCertificate field value
func (o *CertguardX509CertGuard) GetCaCertificate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CaCertificate
}

// GetCaCertificateOk returns a tuple with the CaCertificate field value
// and a boolean to check if the value has been set.
func (o *CertguardX509CertGuard) GetCaCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CaCertificate, true
}

// SetCaCertificate sets field value
func (o *CertguardX509CertGuard) SetCaCertificate(v string) {
	o.CaCertificate = v
}

func (o CertguardX509CertGuard) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertguardX509CertGuard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["ca_certificate"] = o.CaCertificate
	return toSerialize, nil
}

func (o *CertguardX509CertGuard) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"ca_certificate",
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

	varCertguardX509CertGuard := _CertguardX509CertGuard{}

	err = json.Unmarshal(bytes, &varCertguardX509CertGuard)

	if err != nil {
		return err
	}

	*o = CertguardX509CertGuard(varCertguardX509CertGuard)

	return err
}

type NullableCertguardX509CertGuard struct {
	value *CertguardX509CertGuard
	isSet bool
}

func (v NullableCertguardX509CertGuard) Get() *CertguardX509CertGuard {
	return v.value
}

func (v *NullableCertguardX509CertGuard) Set(val *CertguardX509CertGuard) {
	v.value = val
	v.isSet = true
}

func (v NullableCertguardX509CertGuard) IsSet() bool {
	return v.isSet
}

func (v *NullableCertguardX509CertGuard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertguardX509CertGuard(val *CertguardX509CertGuard) *NullableCertguardX509CertGuard {
	return &NullableCertguardX509CertGuard{value: val, isSet: true}
}

func (v NullableCertguardX509CertGuard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertguardX509CertGuard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


