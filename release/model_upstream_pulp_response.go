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

// checks if the UpstreamPulpResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpstreamPulpResponse{}

// UpstreamPulpResponse Serializer for a Server.
type UpstreamPulpResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// Timestamp of the most recent update of the remote.
	PulpLastUpdated *time.Time `json:"pulp_last_updated,omitempty"`
	// A unique name for this Pulp server.
	Name string `json:"name"`
	// The transport, hostname, and an optional port of the Pulp server. e.g. https://example.com
	BaseUrl string `json:"base_url"`
	// The API root. Defaults to '/pulp/'.
	ApiRoot string `json:"api_root"`
	// The domain of the Pulp server if enabled.
	Domain NullableString `json:"domain,omitempty"`
	// A PEM encoded CA certificate used to validate the server certificate presented by the remote server.
	CaCert NullableString `json:"ca_cert,omitempty"`
	// A PEM encoded client certificate used for authentication.
	ClientCert NullableString `json:"client_cert,omitempty"`
	// If True, TLS peer validation must be performed.
	TlsValidation *bool `json:"tls_validation,omitempty"`
	// List of hidden (write only) fields
	HiddenFields []RemoteResponseHiddenFieldsInner `json:"hidden_fields,omitempty"`
	// One or more comma separated labels that will be used to filter distributions on the upstream Pulp. E.g. \"foo=bar,key=val\" or \"foo,key\"
	PulpLabelSelect NullableString `json:"pulp_label_select,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpstreamPulpResponse UpstreamPulpResponse

// NewUpstreamPulpResponse instantiates a new UpstreamPulpResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpstreamPulpResponse(name string, baseUrl string, apiRoot string) *UpstreamPulpResponse {
	this := UpstreamPulpResponse{}
	this.Name = name
	this.BaseUrl = baseUrl
	this.ApiRoot = apiRoot
	return &this
}

// NewUpstreamPulpResponseWithDefaults instantiates a new UpstreamPulpResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpstreamPulpResponseWithDefaults() *UpstreamPulpResponse {
	this := UpstreamPulpResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *UpstreamPulpResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpstreamPulpResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *UpstreamPulpResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *UpstreamPulpResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *UpstreamPulpResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpstreamPulpResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *UpstreamPulpResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *UpstreamPulpResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetPulpLastUpdated returns the PulpLastUpdated field value if set, zero value otherwise.
func (o *UpstreamPulpResponse) GetPulpLastUpdated() time.Time {
	if o == nil || IsNil(o.PulpLastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.PulpLastUpdated
}

// GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpstreamPulpResponse) GetPulpLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpLastUpdated) {
		return nil, false
	}
	return o.PulpLastUpdated, true
}

// HasPulpLastUpdated returns a boolean if a field has been set.
func (o *UpstreamPulpResponse) HasPulpLastUpdated() bool {
	if o != nil && !IsNil(o.PulpLastUpdated) {
		return true
	}

	return false
}

// SetPulpLastUpdated gets a reference to the given time.Time and assigns it to the PulpLastUpdated field.
func (o *UpstreamPulpResponse) SetPulpLastUpdated(v time.Time) {
	o.PulpLastUpdated = &v
}

// GetName returns the Name field value
func (o *UpstreamPulpResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpstreamPulpResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpstreamPulpResponse) SetName(v string) {
	o.Name = v
}

// GetBaseUrl returns the BaseUrl field value
func (o *UpstreamPulpResponse) GetBaseUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value
// and a boolean to check if the value has been set.
func (o *UpstreamPulpResponse) GetBaseUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseUrl, true
}

// SetBaseUrl sets field value
func (o *UpstreamPulpResponse) SetBaseUrl(v string) {
	o.BaseUrl = v
}

// GetApiRoot returns the ApiRoot field value
func (o *UpstreamPulpResponse) GetApiRoot() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiRoot
}

// GetApiRootOk returns a tuple with the ApiRoot field value
// and a boolean to check if the value has been set.
func (o *UpstreamPulpResponse) GetApiRootOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiRoot, true
}

// SetApiRoot sets field value
func (o *UpstreamPulpResponse) SetApiRoot(v string) {
	o.ApiRoot = v
}

// GetDomain returns the Domain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpstreamPulpResponse) GetDomain() string {
	if o == nil || IsNil(o.Domain.Get()) {
		var ret string
		return ret
	}
	return *o.Domain.Get()
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpstreamPulpResponse) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Domain.Get(), o.Domain.IsSet()
}

// HasDomain returns a boolean if a field has been set.
func (o *UpstreamPulpResponse) HasDomain() bool {
	if o != nil && o.Domain.IsSet() {
		return true
	}

	return false
}

// SetDomain gets a reference to the given NullableString and assigns it to the Domain field.
func (o *UpstreamPulpResponse) SetDomain(v string) {
	o.Domain.Set(&v)
}
// SetDomainNil sets the value for Domain to be an explicit nil
func (o *UpstreamPulpResponse) SetDomainNil() {
	o.Domain.Set(nil)
}

// UnsetDomain ensures that no value is present for Domain, not even an explicit nil
func (o *UpstreamPulpResponse) UnsetDomain() {
	o.Domain.Unset()
}

// GetCaCert returns the CaCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpstreamPulpResponse) GetCaCert() string {
	if o == nil || IsNil(o.CaCert.Get()) {
		var ret string
		return ret
	}
	return *o.CaCert.Get()
}

// GetCaCertOk returns a tuple with the CaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpstreamPulpResponse) GetCaCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CaCert.Get(), o.CaCert.IsSet()
}

// HasCaCert returns a boolean if a field has been set.
func (o *UpstreamPulpResponse) HasCaCert() bool {
	if o != nil && o.CaCert.IsSet() {
		return true
	}

	return false
}

// SetCaCert gets a reference to the given NullableString and assigns it to the CaCert field.
func (o *UpstreamPulpResponse) SetCaCert(v string) {
	o.CaCert.Set(&v)
}
// SetCaCertNil sets the value for CaCert to be an explicit nil
func (o *UpstreamPulpResponse) SetCaCertNil() {
	o.CaCert.Set(nil)
}

// UnsetCaCert ensures that no value is present for CaCert, not even an explicit nil
func (o *UpstreamPulpResponse) UnsetCaCert() {
	o.CaCert.Unset()
}

// GetClientCert returns the ClientCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpstreamPulpResponse) GetClientCert() string {
	if o == nil || IsNil(o.ClientCert.Get()) {
		var ret string
		return ret
	}
	return *o.ClientCert.Get()
}

// GetClientCertOk returns a tuple with the ClientCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpstreamPulpResponse) GetClientCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientCert.Get(), o.ClientCert.IsSet()
}

// HasClientCert returns a boolean if a field has been set.
func (o *UpstreamPulpResponse) HasClientCert() bool {
	if o != nil && o.ClientCert.IsSet() {
		return true
	}

	return false
}

// SetClientCert gets a reference to the given NullableString and assigns it to the ClientCert field.
func (o *UpstreamPulpResponse) SetClientCert(v string) {
	o.ClientCert.Set(&v)
}
// SetClientCertNil sets the value for ClientCert to be an explicit nil
func (o *UpstreamPulpResponse) SetClientCertNil() {
	o.ClientCert.Set(nil)
}

// UnsetClientCert ensures that no value is present for ClientCert, not even an explicit nil
func (o *UpstreamPulpResponse) UnsetClientCert() {
	o.ClientCert.Unset()
}

// GetTlsValidation returns the TlsValidation field value if set, zero value otherwise.
func (o *UpstreamPulpResponse) GetTlsValidation() bool {
	if o == nil || IsNil(o.TlsValidation) {
		var ret bool
		return ret
	}
	return *o.TlsValidation
}

// GetTlsValidationOk returns a tuple with the TlsValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpstreamPulpResponse) GetTlsValidationOk() (*bool, bool) {
	if o == nil || IsNil(o.TlsValidation) {
		return nil, false
	}
	return o.TlsValidation, true
}

// HasTlsValidation returns a boolean if a field has been set.
func (o *UpstreamPulpResponse) HasTlsValidation() bool {
	if o != nil && !IsNil(o.TlsValidation) {
		return true
	}

	return false
}

// SetTlsValidation gets a reference to the given bool and assigns it to the TlsValidation field.
func (o *UpstreamPulpResponse) SetTlsValidation(v bool) {
	o.TlsValidation = &v
}

// GetHiddenFields returns the HiddenFields field value if set, zero value otherwise.
func (o *UpstreamPulpResponse) GetHiddenFields() []RemoteResponseHiddenFieldsInner {
	if o == nil || IsNil(o.HiddenFields) {
		var ret []RemoteResponseHiddenFieldsInner
		return ret
	}
	return o.HiddenFields
}

// GetHiddenFieldsOk returns a tuple with the HiddenFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpstreamPulpResponse) GetHiddenFieldsOk() ([]RemoteResponseHiddenFieldsInner, bool) {
	if o == nil || IsNil(o.HiddenFields) {
		return nil, false
	}
	return o.HiddenFields, true
}

// HasHiddenFields returns a boolean if a field has been set.
func (o *UpstreamPulpResponse) HasHiddenFields() bool {
	if o != nil && !IsNil(o.HiddenFields) {
		return true
	}

	return false
}

// SetHiddenFields gets a reference to the given []RemoteResponseHiddenFieldsInner and assigns it to the HiddenFields field.
func (o *UpstreamPulpResponse) SetHiddenFields(v []RemoteResponseHiddenFieldsInner) {
	o.HiddenFields = v
}

// GetPulpLabelSelect returns the PulpLabelSelect field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpstreamPulpResponse) GetPulpLabelSelect() string {
	if o == nil || IsNil(o.PulpLabelSelect.Get()) {
		var ret string
		return ret
	}
	return *o.PulpLabelSelect.Get()
}

// GetPulpLabelSelectOk returns a tuple with the PulpLabelSelect field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpstreamPulpResponse) GetPulpLabelSelectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PulpLabelSelect.Get(), o.PulpLabelSelect.IsSet()
}

// HasPulpLabelSelect returns a boolean if a field has been set.
func (o *UpstreamPulpResponse) HasPulpLabelSelect() bool {
	if o != nil && o.PulpLabelSelect.IsSet() {
		return true
	}

	return false
}

// SetPulpLabelSelect gets a reference to the given NullableString and assigns it to the PulpLabelSelect field.
func (o *UpstreamPulpResponse) SetPulpLabelSelect(v string) {
	o.PulpLabelSelect.Set(&v)
}
// SetPulpLabelSelectNil sets the value for PulpLabelSelect to be an explicit nil
func (o *UpstreamPulpResponse) SetPulpLabelSelectNil() {
	o.PulpLabelSelect.Set(nil)
}

// UnsetPulpLabelSelect ensures that no value is present for PulpLabelSelect, not even an explicit nil
func (o *UpstreamPulpResponse) UnsetPulpLabelSelect() {
	o.PulpLabelSelect.Unset()
}

func (o UpstreamPulpResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpstreamPulpResponse) ToMap() (map[string]interface{}, error) {
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
	toSerialize["base_url"] = o.BaseUrl
	toSerialize["api_root"] = o.ApiRoot
	if o.Domain.IsSet() {
		toSerialize["domain"] = o.Domain.Get()
	}
	if o.CaCert.IsSet() {
		toSerialize["ca_cert"] = o.CaCert.Get()
	}
	if o.ClientCert.IsSet() {
		toSerialize["client_cert"] = o.ClientCert.Get()
	}
	if !IsNil(o.TlsValidation) {
		toSerialize["tls_validation"] = o.TlsValidation
	}
	if !IsNil(o.HiddenFields) {
		toSerialize["hidden_fields"] = o.HiddenFields
	}
	if o.PulpLabelSelect.IsSet() {
		toSerialize["pulp_label_select"] = o.PulpLabelSelect.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpstreamPulpResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"base_url",
		"api_root",
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

	varUpstreamPulpResponse := _UpstreamPulpResponse{}

	err = json.Unmarshal(data, &varUpstreamPulpResponse)

	if err != nil {
		return err
	}

	*o = UpstreamPulpResponse(varUpstreamPulpResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pulp_href")
		delete(additionalProperties, "pulp_created")
		delete(additionalProperties, "pulp_last_updated")
		delete(additionalProperties, "name")
		delete(additionalProperties, "base_url")
		delete(additionalProperties, "api_root")
		delete(additionalProperties, "domain")
		delete(additionalProperties, "ca_cert")
		delete(additionalProperties, "client_cert")
		delete(additionalProperties, "tls_validation")
		delete(additionalProperties, "hidden_fields")
		delete(additionalProperties, "pulp_label_select")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpstreamPulpResponse struct {
	value *UpstreamPulpResponse
	isSet bool
}

func (v NullableUpstreamPulpResponse) Get() *UpstreamPulpResponse {
	return v.value
}

func (v *NullableUpstreamPulpResponse) Set(val *UpstreamPulpResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpstreamPulpResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpstreamPulpResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpstreamPulpResponse(val *UpstreamPulpResponse) *NullableUpstreamPulpResponse {
	return &NullableUpstreamPulpResponse{value: val, isSet: true}
}

func (v NullableUpstreamPulpResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpstreamPulpResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


