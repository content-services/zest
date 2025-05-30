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
)

// checks if the PatchedUpstreamPulp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedUpstreamPulp{}

// PatchedUpstreamPulp Serializer for a Server.
type PatchedUpstreamPulp struct {
	// A unique name for this Pulp server.
	Name *string `json:"name,omitempty"`
	// The transport, hostname, and an optional port of the Pulp server. e.g. https://example.com
	BaseUrl *string `json:"base_url,omitempty"`
	// The API root. Defaults to '/pulp/'.
	ApiRoot *string `json:"api_root,omitempty"`
	// The domain of the Pulp server if enabled.
	Domain NullableString `json:"domain,omitempty"`
	// A PEM encoded CA certificate used to validate the server certificate presented by the remote server.
	CaCert NullableString `json:"ca_cert,omitempty"`
	// A PEM encoded client certificate used for authentication.
	ClientCert NullableString `json:"client_cert,omitempty"`
	// A PEM encoded private key used for authentication.
	ClientKey NullableString `json:"client_key,omitempty"`
	// If True, TLS peer validation must be performed.
	TlsValidation *bool `json:"tls_validation,omitempty"`
	// The username to be used for authentication when syncing.
	Username NullableString `json:"username,omitempty"`
	// The password to be used for authentication when syncing. Extra leading and trailing whitespace characters are not trimmed.
	Password NullableString `json:"password,omitempty"`
	// Filter distributions on the upstream Pulp using complex filtering. E.g. pulp_label_select=\"foo\" OR pulp_label_select=\"key=val\"
	QSelect NullableString `json:"q_select,omitempty"`
	// Policy for how replicate will manage the local objects within the domain.* `all` - Replicate manages ALL local objects within the domain.* `labeled` - Replicate will only manage the objects created from a previous replication, unlabled local objects will be untouched.* `nodelete` - Replicate will not delete any local object whether they were created by replication or not.
	Policy *Policy357Enum `json:"policy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedUpstreamPulp PatchedUpstreamPulp

// NewPatchedUpstreamPulp instantiates a new PatchedUpstreamPulp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedUpstreamPulp() *PatchedUpstreamPulp {
	this := PatchedUpstreamPulp{}
	return &this
}

// NewPatchedUpstreamPulpWithDefaults instantiates a new PatchedUpstreamPulp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedUpstreamPulpWithDefaults() *PatchedUpstreamPulp {
	this := PatchedUpstreamPulp{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedUpstreamPulp) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUpstreamPulp) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedUpstreamPulp) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedUpstreamPulp) SetName(v string) {
	o.Name = &v
}

// GetBaseUrl returns the BaseUrl field value if set, zero value otherwise.
func (o *PatchedUpstreamPulp) GetBaseUrl() string {
	if o == nil || IsNil(o.BaseUrl) {
		var ret string
		return ret
	}
	return *o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUpstreamPulp) GetBaseUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BaseUrl) {
		return nil, false
	}
	return o.BaseUrl, true
}

// HasBaseUrl returns a boolean if a field has been set.
func (o *PatchedUpstreamPulp) HasBaseUrl() bool {
	if o != nil && !IsNil(o.BaseUrl) {
		return true
	}

	return false
}

// SetBaseUrl gets a reference to the given string and assigns it to the BaseUrl field.
func (o *PatchedUpstreamPulp) SetBaseUrl(v string) {
	o.BaseUrl = &v
}

// GetApiRoot returns the ApiRoot field value if set, zero value otherwise.
func (o *PatchedUpstreamPulp) GetApiRoot() string {
	if o == nil || IsNil(o.ApiRoot) {
		var ret string
		return ret
	}
	return *o.ApiRoot
}

// GetApiRootOk returns a tuple with the ApiRoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUpstreamPulp) GetApiRootOk() (*string, bool) {
	if o == nil || IsNil(o.ApiRoot) {
		return nil, false
	}
	return o.ApiRoot, true
}

// HasApiRoot returns a boolean if a field has been set.
func (o *PatchedUpstreamPulp) HasApiRoot() bool {
	if o != nil && !IsNil(o.ApiRoot) {
		return true
	}

	return false
}

// SetApiRoot gets a reference to the given string and assigns it to the ApiRoot field.
func (o *PatchedUpstreamPulp) SetApiRoot(v string) {
	o.ApiRoot = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedUpstreamPulp) GetDomain() string {
	if o == nil || IsNil(o.Domain.Get()) {
		var ret string
		return ret
	}
	return *o.Domain.Get()
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedUpstreamPulp) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Domain.Get(), o.Domain.IsSet()
}

// HasDomain returns a boolean if a field has been set.
func (o *PatchedUpstreamPulp) HasDomain() bool {
	if o != nil && o.Domain.IsSet() {
		return true
	}

	return false
}

// SetDomain gets a reference to the given NullableString and assigns it to the Domain field.
func (o *PatchedUpstreamPulp) SetDomain(v string) {
	o.Domain.Set(&v)
}
// SetDomainNil sets the value for Domain to be an explicit nil
func (o *PatchedUpstreamPulp) SetDomainNil() {
	o.Domain.Set(nil)
}

// UnsetDomain ensures that no value is present for Domain, not even an explicit nil
func (o *PatchedUpstreamPulp) UnsetDomain() {
	o.Domain.Unset()
}

// GetCaCert returns the CaCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedUpstreamPulp) GetCaCert() string {
	if o == nil || IsNil(o.CaCert.Get()) {
		var ret string
		return ret
	}
	return *o.CaCert.Get()
}

// GetCaCertOk returns a tuple with the CaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedUpstreamPulp) GetCaCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CaCert.Get(), o.CaCert.IsSet()
}

// HasCaCert returns a boolean if a field has been set.
func (o *PatchedUpstreamPulp) HasCaCert() bool {
	if o != nil && o.CaCert.IsSet() {
		return true
	}

	return false
}

// SetCaCert gets a reference to the given NullableString and assigns it to the CaCert field.
func (o *PatchedUpstreamPulp) SetCaCert(v string) {
	o.CaCert.Set(&v)
}
// SetCaCertNil sets the value for CaCert to be an explicit nil
func (o *PatchedUpstreamPulp) SetCaCertNil() {
	o.CaCert.Set(nil)
}

// UnsetCaCert ensures that no value is present for CaCert, not even an explicit nil
func (o *PatchedUpstreamPulp) UnsetCaCert() {
	o.CaCert.Unset()
}

// GetClientCert returns the ClientCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedUpstreamPulp) GetClientCert() string {
	if o == nil || IsNil(o.ClientCert.Get()) {
		var ret string
		return ret
	}
	return *o.ClientCert.Get()
}

// GetClientCertOk returns a tuple with the ClientCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedUpstreamPulp) GetClientCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientCert.Get(), o.ClientCert.IsSet()
}

// HasClientCert returns a boolean if a field has been set.
func (o *PatchedUpstreamPulp) HasClientCert() bool {
	if o != nil && o.ClientCert.IsSet() {
		return true
	}

	return false
}

// SetClientCert gets a reference to the given NullableString and assigns it to the ClientCert field.
func (o *PatchedUpstreamPulp) SetClientCert(v string) {
	o.ClientCert.Set(&v)
}
// SetClientCertNil sets the value for ClientCert to be an explicit nil
func (o *PatchedUpstreamPulp) SetClientCertNil() {
	o.ClientCert.Set(nil)
}

// UnsetClientCert ensures that no value is present for ClientCert, not even an explicit nil
func (o *PatchedUpstreamPulp) UnsetClientCert() {
	o.ClientCert.Unset()
}

// GetClientKey returns the ClientKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedUpstreamPulp) GetClientKey() string {
	if o == nil || IsNil(o.ClientKey.Get()) {
		var ret string
		return ret
	}
	return *o.ClientKey.Get()
}

// GetClientKeyOk returns a tuple with the ClientKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedUpstreamPulp) GetClientKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientKey.Get(), o.ClientKey.IsSet()
}

// HasClientKey returns a boolean if a field has been set.
func (o *PatchedUpstreamPulp) HasClientKey() bool {
	if o != nil && o.ClientKey.IsSet() {
		return true
	}

	return false
}

// SetClientKey gets a reference to the given NullableString and assigns it to the ClientKey field.
func (o *PatchedUpstreamPulp) SetClientKey(v string) {
	o.ClientKey.Set(&v)
}
// SetClientKeyNil sets the value for ClientKey to be an explicit nil
func (o *PatchedUpstreamPulp) SetClientKeyNil() {
	o.ClientKey.Set(nil)
}

// UnsetClientKey ensures that no value is present for ClientKey, not even an explicit nil
func (o *PatchedUpstreamPulp) UnsetClientKey() {
	o.ClientKey.Unset()
}

// GetTlsValidation returns the TlsValidation field value if set, zero value otherwise.
func (o *PatchedUpstreamPulp) GetTlsValidation() bool {
	if o == nil || IsNil(o.TlsValidation) {
		var ret bool
		return ret
	}
	return *o.TlsValidation
}

// GetTlsValidationOk returns a tuple with the TlsValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUpstreamPulp) GetTlsValidationOk() (*bool, bool) {
	if o == nil || IsNil(o.TlsValidation) {
		return nil, false
	}
	return o.TlsValidation, true
}

// HasTlsValidation returns a boolean if a field has been set.
func (o *PatchedUpstreamPulp) HasTlsValidation() bool {
	if o != nil && !IsNil(o.TlsValidation) {
		return true
	}

	return false
}

// SetTlsValidation gets a reference to the given bool and assigns it to the TlsValidation field.
func (o *PatchedUpstreamPulp) SetTlsValidation(v bool) {
	o.TlsValidation = &v
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedUpstreamPulp) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedUpstreamPulp) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *PatchedUpstreamPulp) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *PatchedUpstreamPulp) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *PatchedUpstreamPulp) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *PatchedUpstreamPulp) UnsetUsername() {
	o.Username.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedUpstreamPulp) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedUpstreamPulp) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *PatchedUpstreamPulp) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *PatchedUpstreamPulp) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *PatchedUpstreamPulp) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *PatchedUpstreamPulp) UnsetPassword() {
	o.Password.Unset()
}

// GetQSelect returns the QSelect field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedUpstreamPulp) GetQSelect() string {
	if o == nil || IsNil(o.QSelect.Get()) {
		var ret string
		return ret
	}
	return *o.QSelect.Get()
}

// GetQSelectOk returns a tuple with the QSelect field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedUpstreamPulp) GetQSelectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.QSelect.Get(), o.QSelect.IsSet()
}

// HasQSelect returns a boolean if a field has been set.
func (o *PatchedUpstreamPulp) HasQSelect() bool {
	if o != nil && o.QSelect.IsSet() {
		return true
	}

	return false
}

// SetQSelect gets a reference to the given NullableString and assigns it to the QSelect field.
func (o *PatchedUpstreamPulp) SetQSelect(v string) {
	o.QSelect.Set(&v)
}
// SetQSelectNil sets the value for QSelect to be an explicit nil
func (o *PatchedUpstreamPulp) SetQSelectNil() {
	o.QSelect.Set(nil)
}

// UnsetQSelect ensures that no value is present for QSelect, not even an explicit nil
func (o *PatchedUpstreamPulp) UnsetQSelect() {
	o.QSelect.Unset()
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *PatchedUpstreamPulp) GetPolicy() Policy357Enum {
	if o == nil || IsNil(o.Policy) {
		var ret Policy357Enum
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedUpstreamPulp) GetPolicyOk() (*Policy357Enum, bool) {
	if o == nil || IsNil(o.Policy) {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *PatchedUpstreamPulp) HasPolicy() bool {
	if o != nil && !IsNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given Policy357Enum and assigns it to the Policy field.
func (o *PatchedUpstreamPulp) SetPolicy(v Policy357Enum) {
	o.Policy = &v
}

func (o PatchedUpstreamPulp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedUpstreamPulp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.BaseUrl) {
		toSerialize["base_url"] = o.BaseUrl
	}
	if !IsNil(o.ApiRoot) {
		toSerialize["api_root"] = o.ApiRoot
	}
	if o.Domain.IsSet() {
		toSerialize["domain"] = o.Domain.Get()
	}
	if o.CaCert.IsSet() {
		toSerialize["ca_cert"] = o.CaCert.Get()
	}
	if o.ClientCert.IsSet() {
		toSerialize["client_cert"] = o.ClientCert.Get()
	}
	if o.ClientKey.IsSet() {
		toSerialize["client_key"] = o.ClientKey.Get()
	}
	if !IsNil(o.TlsValidation) {
		toSerialize["tls_validation"] = o.TlsValidation
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.QSelect.IsSet() {
		toSerialize["q_select"] = o.QSelect.Get()
	}
	if !IsNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedUpstreamPulp) UnmarshalJSON(data []byte) (err error) {
	varPatchedUpstreamPulp := _PatchedUpstreamPulp{}

	err = json.Unmarshal(data, &varPatchedUpstreamPulp)

	if err != nil {
		return err
	}

	*o = PatchedUpstreamPulp(varPatchedUpstreamPulp)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "base_url")
		delete(additionalProperties, "api_root")
		delete(additionalProperties, "domain")
		delete(additionalProperties, "ca_cert")
		delete(additionalProperties, "client_cert")
		delete(additionalProperties, "client_key")
		delete(additionalProperties, "tls_validation")
		delete(additionalProperties, "username")
		delete(additionalProperties, "password")
		delete(additionalProperties, "q_select")
		delete(additionalProperties, "policy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedUpstreamPulp struct {
	value *PatchedUpstreamPulp
	isSet bool
}

func (v NullablePatchedUpstreamPulp) Get() *PatchedUpstreamPulp {
	return v.value
}

func (v *NullablePatchedUpstreamPulp) Set(val *PatchedUpstreamPulp) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedUpstreamPulp) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedUpstreamPulp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedUpstreamPulp(val *PatchedUpstreamPulp) *NullablePatchedUpstreamPulp {
	return &NullablePatchedUpstreamPulp{value: val, isSet: true}
}

func (v NullablePatchedUpstreamPulp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedUpstreamPulp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


