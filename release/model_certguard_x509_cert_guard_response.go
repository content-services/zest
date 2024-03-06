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
	"bytes"
	"fmt"
)

// checks if the CertguardX509CertGuardResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertguardX509CertGuardResponse{}

// CertguardX509CertGuardResponse X.509 Content Guard Serializer.
type CertguardX509CertGuardResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same.
	PulpLastUpdated *time.Time `json:"pulp_last_updated,omitempty"`
	// The unique name.
	Name string `json:"name"`
	// An optional description.
	Description NullableString `json:"description,omitempty"`
	// A Certificate Authority (CA) certificate (or a bundle thereof) used to verify client-certificate authenticity.
	CaCertificate string `json:"ca_certificate"`
}

type _CertguardX509CertGuardResponse CertguardX509CertGuardResponse

// NewCertguardX509CertGuardResponse instantiates a new CertguardX509CertGuardResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertguardX509CertGuardResponse(name string, caCertificate string) *CertguardX509CertGuardResponse {
	this := CertguardX509CertGuardResponse{}
	this.Name = name
	this.CaCertificate = caCertificate
	return &this
}

// NewCertguardX509CertGuardResponseWithDefaults instantiates a new CertguardX509CertGuardResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertguardX509CertGuardResponseWithDefaults() *CertguardX509CertGuardResponse {
	this := CertguardX509CertGuardResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *CertguardX509CertGuardResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertguardX509CertGuardResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *CertguardX509CertGuardResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *CertguardX509CertGuardResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *CertguardX509CertGuardResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertguardX509CertGuardResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *CertguardX509CertGuardResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *CertguardX509CertGuardResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetPulpLastUpdated returns the PulpLastUpdated field value if set, zero value otherwise.
func (o *CertguardX509CertGuardResponse) GetPulpLastUpdated() time.Time {
	if o == nil || IsNil(o.PulpLastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.PulpLastUpdated
}

// GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertguardX509CertGuardResponse) GetPulpLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpLastUpdated) {
		return nil, false
	}
	return o.PulpLastUpdated, true
}

// HasPulpLastUpdated returns a boolean if a field has been set.
func (o *CertguardX509CertGuardResponse) HasPulpLastUpdated() bool {
	if o != nil && !IsNil(o.PulpLastUpdated) {
		return true
	}

	return false
}

// SetPulpLastUpdated gets a reference to the given time.Time and assigns it to the PulpLastUpdated field.
func (o *CertguardX509CertGuardResponse) SetPulpLastUpdated(v time.Time) {
	o.PulpLastUpdated = &v
}

// GetName returns the Name field value
func (o *CertguardX509CertGuardResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CertguardX509CertGuardResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CertguardX509CertGuardResponse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CertguardX509CertGuardResponse) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CertguardX509CertGuardResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CertguardX509CertGuardResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CertguardX509CertGuardResponse) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CertguardX509CertGuardResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CertguardX509CertGuardResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetCaCertificate returns the CaCertificate field value
func (o *CertguardX509CertGuardResponse) GetCaCertificate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CaCertificate
}

// GetCaCertificateOk returns a tuple with the CaCertificate field value
// and a boolean to check if the value has been set.
func (o *CertguardX509CertGuardResponse) GetCaCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CaCertificate, true
}

// SetCaCertificate sets field value
func (o *CertguardX509CertGuardResponse) SetCaCertificate(v string) {
	o.CaCertificate = v
}

func (o CertguardX509CertGuardResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertguardX509CertGuardResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PulpHref) {
		toSerialize["pulp_href"] = o.PulpHref
	}
	if !IsNil(o.PulpCreated) {
		toSerialize["pulp_created"] = o.PulpCreated
	}
	if !IsNil(o.PulpLastUpdated) {
		toSerialize["pulp_last_updated"] = o.PulpLastUpdated
	}
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["ca_certificate"] = o.CaCertificate
	return toSerialize, nil
}

func (o *CertguardX509CertGuardResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"ca_certificate",
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

	varCertguardX509CertGuardResponse := _CertguardX509CertGuardResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCertguardX509CertGuardResponse)

	if err != nil {
		return err
	}

	*o = CertguardX509CertGuardResponse(varCertguardX509CertGuardResponse)

	return err
}

type NullableCertguardX509CertGuardResponse struct {
	value *CertguardX509CertGuardResponse
	isSet bool
}

func (v NullableCertguardX509CertGuardResponse) Get() *CertguardX509CertGuardResponse {
	return v.value
}

func (v *NullableCertguardX509CertGuardResponse) Set(val *CertguardX509CertGuardResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCertguardX509CertGuardResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCertguardX509CertGuardResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertguardX509CertGuardResponse(val *CertguardX509CertGuardResponse) *NullableCertguardX509CertGuardResponse {
	return &NullableCertguardX509CertGuardResponse{value: val, isSet: true}
}

func (v NullableCertguardX509CertGuardResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertguardX509CertGuardResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


