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

// checks if the SigningServiceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SigningServiceResponse{}

// SigningServiceResponse A serializer for the model declaring a signing service.
type SigningServiceResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// A unique name used to recognize a script.
	Name string `json:"name"`
	// The value of a public key used for the repository verification.
	PublicKey string `json:"public_key"`
	// The fingerprint of the public key.
	PubkeyFingerprint string `json:"pubkey_fingerprint"`
	// An absolute path to a script which is going to be used for the signing.
	Script string `json:"script"`
}

type _SigningServiceResponse SigningServiceResponse

// NewSigningServiceResponse instantiates a new SigningServiceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSigningServiceResponse(name string, publicKey string, pubkeyFingerprint string, script string) *SigningServiceResponse {
	this := SigningServiceResponse{}
	this.Name = name
	this.PublicKey = publicKey
	this.PubkeyFingerprint = pubkeyFingerprint
	this.Script = script
	return &this
}

// NewSigningServiceResponseWithDefaults instantiates a new SigningServiceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSigningServiceResponseWithDefaults() *SigningServiceResponse {
	this := SigningServiceResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *SigningServiceResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigningServiceResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *SigningServiceResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *SigningServiceResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *SigningServiceResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigningServiceResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *SigningServiceResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *SigningServiceResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetName returns the Name field value
func (o *SigningServiceResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SigningServiceResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SigningServiceResponse) SetName(v string) {
	o.Name = v
}

// GetPublicKey returns the PublicKey field value
func (o *SigningServiceResponse) GetPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
func (o *SigningServiceResponse) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKey, true
}

// SetPublicKey sets field value
func (o *SigningServiceResponse) SetPublicKey(v string) {
	o.PublicKey = v
}

// GetPubkeyFingerprint returns the PubkeyFingerprint field value
func (o *SigningServiceResponse) GetPubkeyFingerprint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PubkeyFingerprint
}

// GetPubkeyFingerprintOk returns a tuple with the PubkeyFingerprint field value
// and a boolean to check if the value has been set.
func (o *SigningServiceResponse) GetPubkeyFingerprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PubkeyFingerprint, true
}

// SetPubkeyFingerprint sets field value
func (o *SigningServiceResponse) SetPubkeyFingerprint(v string) {
	o.PubkeyFingerprint = v
}

// GetScript returns the Script field value
func (o *SigningServiceResponse) GetScript() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Script
}

// GetScriptOk returns a tuple with the Script field value
// and a boolean to check if the value has been set.
func (o *SigningServiceResponse) GetScriptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Script, true
}

// SetScript sets field value
func (o *SigningServiceResponse) SetScript(v string) {
	o.Script = v
}

func (o SigningServiceResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SigningServiceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PulpHref) {
		toSerialize["pulp_href"] = o.PulpHref
	}
	if !IsNil(o.PulpCreated) {
		toSerialize["pulp_created"] = o.PulpCreated
	}
	toSerialize["name"] = o.Name
	toSerialize["public_key"] = o.PublicKey
	toSerialize["pubkey_fingerprint"] = o.PubkeyFingerprint
	toSerialize["script"] = o.Script
	return toSerialize, nil
}

func (o *SigningServiceResponse) UnmarshalJSON(bytes []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"public_key",
		"pubkey_fingerprint",
		"script",
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

	varSigningServiceResponse := _SigningServiceResponse{}

	err = json.Unmarshal(bytes, &varSigningServiceResponse)

	if err != nil {
		return err
	}

	*o = SigningServiceResponse(varSigningServiceResponse)

	return err
}

type NullableSigningServiceResponse struct {
	value *SigningServiceResponse
	isSet bool
}

func (v NullableSigningServiceResponse) Get() *SigningServiceResponse {
	return v.value
}

func (v *NullableSigningServiceResponse) Set(val *SigningServiceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSigningServiceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSigningServiceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSigningServiceResponse(val *SigningServiceResponse) *NullableSigningServiceResponse {
	return &NullableSigningServiceResponse{value: val, isSet: true}
}

func (v NullableSigningServiceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSigningServiceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


