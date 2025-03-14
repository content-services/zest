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

// checks if the NestedOpenPGPSignatureResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedOpenPGPSignatureResponse{}

// NestedOpenPGPSignatureResponse Base serializer for use with [pulpcore.app.models.Model][]This ensures that all Serializers provide values for the 'pulp_href` field.The class provides a default for the ``ref_name`` attribute in theModelSerializers's ``Meta`` class. This ensures that the OpenAPI definitionsof plugins are namespaced properly.
type NestedOpenPGPSignatureResponse struct {
	Issuer NullableString `json:"issuer,omitempty"`
	Created time.Time `json:"created"`
	ExpirationTime NullableString `json:"expiration_time,omitempty"`
	SignersUserId NullableString `json:"signers_user_id,omitempty"`
	KeyExpirationTime NullableString `json:"key_expiration_time,omitempty"`
	Expired bool `json:"expired"`
	KeyExpired *string `json:"key_expired,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NestedOpenPGPSignatureResponse NestedOpenPGPSignatureResponse

// NewNestedOpenPGPSignatureResponse instantiates a new NestedOpenPGPSignatureResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedOpenPGPSignatureResponse(created time.Time, expired bool) *NestedOpenPGPSignatureResponse {
	this := NestedOpenPGPSignatureResponse{}
	this.Created = created
	this.Expired = expired
	return &this
}

// NewNestedOpenPGPSignatureResponseWithDefaults instantiates a new NestedOpenPGPSignatureResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedOpenPGPSignatureResponseWithDefaults() *NestedOpenPGPSignatureResponse {
	this := NestedOpenPGPSignatureResponse{}
	return &this
}

// GetIssuer returns the Issuer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NestedOpenPGPSignatureResponse) GetIssuer() string {
	if o == nil || IsNil(o.Issuer.Get()) {
		var ret string
		return ret
	}
	return *o.Issuer.Get()
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NestedOpenPGPSignatureResponse) GetIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Issuer.Get(), o.Issuer.IsSet()
}

// HasIssuer returns a boolean if a field has been set.
func (o *NestedOpenPGPSignatureResponse) HasIssuer() bool {
	if o != nil && o.Issuer.IsSet() {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given NullableString and assigns it to the Issuer field.
func (o *NestedOpenPGPSignatureResponse) SetIssuer(v string) {
	o.Issuer.Set(&v)
}
// SetIssuerNil sets the value for Issuer to be an explicit nil
func (o *NestedOpenPGPSignatureResponse) SetIssuerNil() {
	o.Issuer.Set(nil)
}

// UnsetIssuer ensures that no value is present for Issuer, not even an explicit nil
func (o *NestedOpenPGPSignatureResponse) UnsetIssuer() {
	o.Issuer.Unset()
}

// GetCreated returns the Created field value
func (o *NestedOpenPGPSignatureResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *NestedOpenPGPSignatureResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *NestedOpenPGPSignatureResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetExpirationTime returns the ExpirationTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NestedOpenPGPSignatureResponse) GetExpirationTime() string {
	if o == nil || IsNil(o.ExpirationTime.Get()) {
		var ret string
		return ret
	}
	return *o.ExpirationTime.Get()
}

// GetExpirationTimeOk returns a tuple with the ExpirationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NestedOpenPGPSignatureResponse) GetExpirationTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpirationTime.Get(), o.ExpirationTime.IsSet()
}

// HasExpirationTime returns a boolean if a field has been set.
func (o *NestedOpenPGPSignatureResponse) HasExpirationTime() bool {
	if o != nil && o.ExpirationTime.IsSet() {
		return true
	}

	return false
}

// SetExpirationTime gets a reference to the given NullableString and assigns it to the ExpirationTime field.
func (o *NestedOpenPGPSignatureResponse) SetExpirationTime(v string) {
	o.ExpirationTime.Set(&v)
}
// SetExpirationTimeNil sets the value for ExpirationTime to be an explicit nil
func (o *NestedOpenPGPSignatureResponse) SetExpirationTimeNil() {
	o.ExpirationTime.Set(nil)
}

// UnsetExpirationTime ensures that no value is present for ExpirationTime, not even an explicit nil
func (o *NestedOpenPGPSignatureResponse) UnsetExpirationTime() {
	o.ExpirationTime.Unset()
}

// GetSignersUserId returns the SignersUserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NestedOpenPGPSignatureResponse) GetSignersUserId() string {
	if o == nil || IsNil(o.SignersUserId.Get()) {
		var ret string
		return ret
	}
	return *o.SignersUserId.Get()
}

// GetSignersUserIdOk returns a tuple with the SignersUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NestedOpenPGPSignatureResponse) GetSignersUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SignersUserId.Get(), o.SignersUserId.IsSet()
}

// HasSignersUserId returns a boolean if a field has been set.
func (o *NestedOpenPGPSignatureResponse) HasSignersUserId() bool {
	if o != nil && o.SignersUserId.IsSet() {
		return true
	}

	return false
}

// SetSignersUserId gets a reference to the given NullableString and assigns it to the SignersUserId field.
func (o *NestedOpenPGPSignatureResponse) SetSignersUserId(v string) {
	o.SignersUserId.Set(&v)
}
// SetSignersUserIdNil sets the value for SignersUserId to be an explicit nil
func (o *NestedOpenPGPSignatureResponse) SetSignersUserIdNil() {
	o.SignersUserId.Set(nil)
}

// UnsetSignersUserId ensures that no value is present for SignersUserId, not even an explicit nil
func (o *NestedOpenPGPSignatureResponse) UnsetSignersUserId() {
	o.SignersUserId.Unset()
}

// GetKeyExpirationTime returns the KeyExpirationTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NestedOpenPGPSignatureResponse) GetKeyExpirationTime() string {
	if o == nil || IsNil(o.KeyExpirationTime.Get()) {
		var ret string
		return ret
	}
	return *o.KeyExpirationTime.Get()
}

// GetKeyExpirationTimeOk returns a tuple with the KeyExpirationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NestedOpenPGPSignatureResponse) GetKeyExpirationTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.KeyExpirationTime.Get(), o.KeyExpirationTime.IsSet()
}

// HasKeyExpirationTime returns a boolean if a field has been set.
func (o *NestedOpenPGPSignatureResponse) HasKeyExpirationTime() bool {
	if o != nil && o.KeyExpirationTime.IsSet() {
		return true
	}

	return false
}

// SetKeyExpirationTime gets a reference to the given NullableString and assigns it to the KeyExpirationTime field.
func (o *NestedOpenPGPSignatureResponse) SetKeyExpirationTime(v string) {
	o.KeyExpirationTime.Set(&v)
}
// SetKeyExpirationTimeNil sets the value for KeyExpirationTime to be an explicit nil
func (o *NestedOpenPGPSignatureResponse) SetKeyExpirationTimeNil() {
	o.KeyExpirationTime.Set(nil)
}

// UnsetKeyExpirationTime ensures that no value is present for KeyExpirationTime, not even an explicit nil
func (o *NestedOpenPGPSignatureResponse) UnsetKeyExpirationTime() {
	o.KeyExpirationTime.Unset()
}

// GetExpired returns the Expired field value
func (o *NestedOpenPGPSignatureResponse) GetExpired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Expired
}

// GetExpiredOk returns a tuple with the Expired field value
// and a boolean to check if the value has been set.
func (o *NestedOpenPGPSignatureResponse) GetExpiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expired, true
}

// SetExpired sets field value
func (o *NestedOpenPGPSignatureResponse) SetExpired(v bool) {
	o.Expired = v
}

// GetKeyExpired returns the KeyExpired field value if set, zero value otherwise.
func (o *NestedOpenPGPSignatureResponse) GetKeyExpired() string {
	if o == nil || IsNil(o.KeyExpired) {
		var ret string
		return ret
	}
	return *o.KeyExpired
}

// GetKeyExpiredOk returns a tuple with the KeyExpired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NestedOpenPGPSignatureResponse) GetKeyExpiredOk() (*string, bool) {
	if o == nil || IsNil(o.KeyExpired) {
		return nil, false
	}
	return o.KeyExpired, true
}

// HasKeyExpired returns a boolean if a field has been set.
func (o *NestedOpenPGPSignatureResponse) HasKeyExpired() bool {
	if o != nil && !IsNil(o.KeyExpired) {
		return true
	}

	return false
}

// SetKeyExpired gets a reference to the given string and assigns it to the KeyExpired field.
func (o *NestedOpenPGPSignatureResponse) SetKeyExpired(v string) {
	o.KeyExpired = &v
}

func (o NestedOpenPGPSignatureResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedOpenPGPSignatureResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Issuer.IsSet() {
		toSerialize["issuer"] = o.Issuer.Get()
	}
	toSerialize["created"] = o.Created
	if o.ExpirationTime.IsSet() {
		toSerialize["expiration_time"] = o.ExpirationTime.Get()
	}
	if o.SignersUserId.IsSet() {
		toSerialize["signers_user_id"] = o.SignersUserId.Get()
	}
	if o.KeyExpirationTime.IsSet() {
		toSerialize["key_expiration_time"] = o.KeyExpirationTime.Get()
	}
	toSerialize["expired"] = o.Expired
	if !IsNil(o.KeyExpired) {
		toSerialize["key_expired"] = o.KeyExpired
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedOpenPGPSignatureResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created",
		"expired",
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

	varNestedOpenPGPSignatureResponse := _NestedOpenPGPSignatureResponse{}

	err = json.Unmarshal(data, &varNestedOpenPGPSignatureResponse)

	if err != nil {
		return err
	}

	*o = NestedOpenPGPSignatureResponse(varNestedOpenPGPSignatureResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "issuer")
		delete(additionalProperties, "created")
		delete(additionalProperties, "expiration_time")
		delete(additionalProperties, "signers_user_id")
		delete(additionalProperties, "key_expiration_time")
		delete(additionalProperties, "expired")
		delete(additionalProperties, "key_expired")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedOpenPGPSignatureResponse struct {
	value *NestedOpenPGPSignatureResponse
	isSet bool
}

func (v NullableNestedOpenPGPSignatureResponse) Get() *NestedOpenPGPSignatureResponse {
	return v.value
}

func (v *NullableNestedOpenPGPSignatureResponse) Set(val *NestedOpenPGPSignatureResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedOpenPGPSignatureResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedOpenPGPSignatureResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedOpenPGPSignatureResponse(val *NestedOpenPGPSignatureResponse) *NullableNestedOpenPGPSignatureResponse {
	return &NullableNestedOpenPGPSignatureResponse{value: val, isSet: true}
}

func (v NullableNestedOpenPGPSignatureResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedOpenPGPSignatureResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


