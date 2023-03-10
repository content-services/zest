/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pulpGoBinding

import (
	"encoding/json"
)

// checks if the PatchedcertguardX509CertGuard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedcertguardX509CertGuard{}

// PatchedcertguardX509CertGuard X.509 Content Guard Serializer.
type PatchedcertguardX509CertGuard struct {
	// The unique name.
	Name *string `json:"name,omitempty"`
	// An optional description.
	Description NullableString `json:"description,omitempty"`
	// A Certificate Authority (CA) certificate (or a bundle thereof) used to verify client-certificate authenticity.
	CaCertificate *string `json:"ca_certificate,omitempty"`
}

// NewPatchedcertguardX509CertGuard instantiates a new PatchedcertguardX509CertGuard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedcertguardX509CertGuard() *PatchedcertguardX509CertGuard {
	this := PatchedcertguardX509CertGuard{}
	return &this
}

// NewPatchedcertguardX509CertGuardWithDefaults instantiates a new PatchedcertguardX509CertGuard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedcertguardX509CertGuardWithDefaults() *PatchedcertguardX509CertGuard {
	this := PatchedcertguardX509CertGuard{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedcertguardX509CertGuard) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedcertguardX509CertGuard) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedcertguardX509CertGuard) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedcertguardX509CertGuard) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedcertguardX509CertGuard) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedcertguardX509CertGuard) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedcertguardX509CertGuard) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PatchedcertguardX509CertGuard) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PatchedcertguardX509CertGuard) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PatchedcertguardX509CertGuard) UnsetDescription() {
	o.Description.Unset()
}

// GetCaCertificate returns the CaCertificate field value if set, zero value otherwise.
func (o *PatchedcertguardX509CertGuard) GetCaCertificate() string {
	if o == nil || IsNil(o.CaCertificate) {
		var ret string
		return ret
	}
	return *o.CaCertificate
}

// GetCaCertificateOk returns a tuple with the CaCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedcertguardX509CertGuard) GetCaCertificateOk() (*string, bool) {
	if o == nil || IsNil(o.CaCertificate) {
		return nil, false
	}
	return o.CaCertificate, true
}

// HasCaCertificate returns a boolean if a field has been set.
func (o *PatchedcertguardX509CertGuard) HasCaCertificate() bool {
	if o != nil && !IsNil(o.CaCertificate) {
		return true
	}

	return false
}

// SetCaCertificate gets a reference to the given string and assigns it to the CaCertificate field.
func (o *PatchedcertguardX509CertGuard) SetCaCertificate(v string) {
	o.CaCertificate = &v
}

func (o PatchedcertguardX509CertGuard) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedcertguardX509CertGuard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.CaCertificate) {
		toSerialize["ca_certificate"] = o.CaCertificate
	}
	return toSerialize, nil
}

type NullablePatchedcertguardX509CertGuard struct {
	value *PatchedcertguardX509CertGuard
	isSet bool
}

func (v NullablePatchedcertguardX509CertGuard) Get() *PatchedcertguardX509CertGuard {
	return v.value
}

func (v *NullablePatchedcertguardX509CertGuard) Set(val *PatchedcertguardX509CertGuard) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedcertguardX509CertGuard) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedcertguardX509CertGuard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedcertguardX509CertGuard(val *PatchedcertguardX509CertGuard) *NullablePatchedcertguardX509CertGuard {
	return &NullablePatchedcertguardX509CertGuard{value: val, isSet: true}
}

func (v NullablePatchedcertguardX509CertGuard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedcertguardX509CertGuard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


