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

// checks if the HeaderContentGuardResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HeaderContentGuardResponse{}

// HeaderContentGuardResponse A serializer for HeaderContentGuard.
type HeaderContentGuardResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// The unique name.
	Name string `json:"name"`
	// An optional description.
	Description NullableString `json:"description,omitempty"`
	// The header name the guard will check on.
	HeaderName string `json:"header_name"`
	// The value that will authorize the request.
	HeaderValue string `json:"header_value"`
	// A JQ syntax compatible filter. If jq_filter is not set, then the value willonly be Base64 decoded and checked as an explicit string match.
	JqFilter NullableString `json:"jq_filter,omitempty"`
}

type _HeaderContentGuardResponse HeaderContentGuardResponse

// NewHeaderContentGuardResponse instantiates a new HeaderContentGuardResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeaderContentGuardResponse(name string, headerName string, headerValue string) *HeaderContentGuardResponse {
	this := HeaderContentGuardResponse{}
	this.Name = name
	this.HeaderName = headerName
	this.HeaderValue = headerValue
	return &this
}

// NewHeaderContentGuardResponseWithDefaults instantiates a new HeaderContentGuardResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeaderContentGuardResponseWithDefaults() *HeaderContentGuardResponse {
	this := HeaderContentGuardResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *HeaderContentGuardResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeaderContentGuardResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *HeaderContentGuardResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *HeaderContentGuardResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *HeaderContentGuardResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeaderContentGuardResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *HeaderContentGuardResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *HeaderContentGuardResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetName returns the Name field value
func (o *HeaderContentGuardResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *HeaderContentGuardResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *HeaderContentGuardResponse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeaderContentGuardResponse) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeaderContentGuardResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *HeaderContentGuardResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *HeaderContentGuardResponse) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *HeaderContentGuardResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *HeaderContentGuardResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetHeaderName returns the HeaderName field value
func (o *HeaderContentGuardResponse) GetHeaderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HeaderName
}

// GetHeaderNameOk returns a tuple with the HeaderName field value
// and a boolean to check if the value has been set.
func (o *HeaderContentGuardResponse) GetHeaderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HeaderName, true
}

// SetHeaderName sets field value
func (o *HeaderContentGuardResponse) SetHeaderName(v string) {
	o.HeaderName = v
}

// GetHeaderValue returns the HeaderValue field value
func (o *HeaderContentGuardResponse) GetHeaderValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HeaderValue
}

// GetHeaderValueOk returns a tuple with the HeaderValue field value
// and a boolean to check if the value has been set.
func (o *HeaderContentGuardResponse) GetHeaderValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HeaderValue, true
}

// SetHeaderValue sets field value
func (o *HeaderContentGuardResponse) SetHeaderValue(v string) {
	o.HeaderValue = v
}

// GetJqFilter returns the JqFilter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HeaderContentGuardResponse) GetJqFilter() string {
	if o == nil || IsNil(o.JqFilter.Get()) {
		var ret string
		return ret
	}
	return *o.JqFilter.Get()
}

// GetJqFilterOk returns a tuple with the JqFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HeaderContentGuardResponse) GetJqFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.JqFilter.Get(), o.JqFilter.IsSet()
}

// HasJqFilter returns a boolean if a field has been set.
func (o *HeaderContentGuardResponse) HasJqFilter() bool {
	if o != nil && o.JqFilter.IsSet() {
		return true
	}

	return false
}

// SetJqFilter gets a reference to the given NullableString and assigns it to the JqFilter field.
func (o *HeaderContentGuardResponse) SetJqFilter(v string) {
	o.JqFilter.Set(&v)
}
// SetJqFilterNil sets the value for JqFilter to be an explicit nil
func (o *HeaderContentGuardResponse) SetJqFilterNil() {
	o.JqFilter.Set(nil)
}

// UnsetJqFilter ensures that no value is present for JqFilter, not even an explicit nil
func (o *HeaderContentGuardResponse) UnsetJqFilter() {
	o.JqFilter.Unset()
}

func (o HeaderContentGuardResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HeaderContentGuardResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PulpHref) {
		toSerialize["pulp_href"] = o.PulpHref
	}
	if !IsNil(o.PulpCreated) {
		toSerialize["pulp_created"] = o.PulpCreated
	}
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["header_name"] = o.HeaderName
	toSerialize["header_value"] = o.HeaderValue
	if o.JqFilter.IsSet() {
		toSerialize["jq_filter"] = o.JqFilter.Get()
	}
	return toSerialize, nil
}

func (o *HeaderContentGuardResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"header_name",
		"header_value",
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

	varHeaderContentGuardResponse := _HeaderContentGuardResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHeaderContentGuardResponse)

	if err != nil {
		return err
	}

	*o = HeaderContentGuardResponse(varHeaderContentGuardResponse)

	return err
}

type NullableHeaderContentGuardResponse struct {
	value *HeaderContentGuardResponse
	isSet bool
}

func (v NullableHeaderContentGuardResponse) Get() *HeaderContentGuardResponse {
	return v.value
}

func (v *NullableHeaderContentGuardResponse) Set(val *HeaderContentGuardResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHeaderContentGuardResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHeaderContentGuardResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeaderContentGuardResponse(val *HeaderContentGuardResponse) *NullableHeaderContentGuardResponse {
	return &NullableHeaderContentGuardResponse{value: val, isSet: true}
}

func (v NullableHeaderContentGuardResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeaderContentGuardResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

