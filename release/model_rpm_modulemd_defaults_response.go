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

// checks if the RpmModulemdDefaultsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RpmModulemdDefaultsResponse{}

// RpmModulemdDefaultsResponse ModulemdDefaults serializer.
type RpmModulemdDefaultsResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// The Pulp Resource Name (PRN).
	Prn *string `json:"prn,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// Timestamp of the last time this resource was updated. Note: for immutable resources - like content, repository versions, and publication - pulp_created and pulp_last_updated dates will be the same.
	PulpLastUpdated *time.Time `json:"pulp_last_updated,omitempty"`
	// A dictionary of arbitrary key/value pairs used to describe a specific Content instance.
	PulpLabels *map[string]string `json:"pulp_labels,omitempty"`
	// Modulemd name.
	Module string `json:"module"`
	// Modulemd default stream.
	Stream string `json:"stream"`
	// Default profiles for modulemd streams.
	Profiles interface{} `json:"profiles"`
	AdditionalProperties map[string]interface{}
}

type _RpmModulemdDefaultsResponse RpmModulemdDefaultsResponse

// NewRpmModulemdDefaultsResponse instantiates a new RpmModulemdDefaultsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRpmModulemdDefaultsResponse(module string, stream string, profiles interface{}) *RpmModulemdDefaultsResponse {
	this := RpmModulemdDefaultsResponse{}
	this.Module = module
	this.Stream = stream
	this.Profiles = profiles
	return &this
}

// NewRpmModulemdDefaultsResponseWithDefaults instantiates a new RpmModulemdDefaultsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRpmModulemdDefaultsResponseWithDefaults() *RpmModulemdDefaultsResponse {
	this := RpmModulemdDefaultsResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *RpmModulemdDefaultsResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmModulemdDefaultsResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *RpmModulemdDefaultsResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *RpmModulemdDefaultsResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPrn returns the Prn field value if set, zero value otherwise.
func (o *RpmModulemdDefaultsResponse) GetPrn() string {
	if o == nil || IsNil(o.Prn) {
		var ret string
		return ret
	}
	return *o.Prn
}

// GetPrnOk returns a tuple with the Prn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmModulemdDefaultsResponse) GetPrnOk() (*string, bool) {
	if o == nil || IsNil(o.Prn) {
		return nil, false
	}
	return o.Prn, true
}

// HasPrn returns a boolean if a field has been set.
func (o *RpmModulemdDefaultsResponse) HasPrn() bool {
	if o != nil && !IsNil(o.Prn) {
		return true
	}

	return false
}

// SetPrn gets a reference to the given string and assigns it to the Prn field.
func (o *RpmModulemdDefaultsResponse) SetPrn(v string) {
	o.Prn = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *RpmModulemdDefaultsResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmModulemdDefaultsResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *RpmModulemdDefaultsResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *RpmModulemdDefaultsResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetPulpLastUpdated returns the PulpLastUpdated field value if set, zero value otherwise.
func (o *RpmModulemdDefaultsResponse) GetPulpLastUpdated() time.Time {
	if o == nil || IsNil(o.PulpLastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.PulpLastUpdated
}

// GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmModulemdDefaultsResponse) GetPulpLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpLastUpdated) {
		return nil, false
	}
	return o.PulpLastUpdated, true
}

// HasPulpLastUpdated returns a boolean if a field has been set.
func (o *RpmModulemdDefaultsResponse) HasPulpLastUpdated() bool {
	if o != nil && !IsNil(o.PulpLastUpdated) {
		return true
	}

	return false
}

// SetPulpLastUpdated gets a reference to the given time.Time and assigns it to the PulpLastUpdated field.
func (o *RpmModulemdDefaultsResponse) SetPulpLastUpdated(v time.Time) {
	o.PulpLastUpdated = &v
}

// GetPulpLabels returns the PulpLabels field value if set, zero value otherwise.
func (o *RpmModulemdDefaultsResponse) GetPulpLabels() map[string]string {
	if o == nil || IsNil(o.PulpLabels) {
		var ret map[string]string
		return ret
	}
	return *o.PulpLabels
}

// GetPulpLabelsOk returns a tuple with the PulpLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmModulemdDefaultsResponse) GetPulpLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PulpLabels) {
		return nil, false
	}
	return o.PulpLabels, true
}

// HasPulpLabels returns a boolean if a field has been set.
func (o *RpmModulemdDefaultsResponse) HasPulpLabels() bool {
	if o != nil && !IsNil(o.PulpLabels) {
		return true
	}

	return false
}

// SetPulpLabels gets a reference to the given map[string]string and assigns it to the PulpLabels field.
func (o *RpmModulemdDefaultsResponse) SetPulpLabels(v map[string]string) {
	o.PulpLabels = &v
}

// GetModule returns the Module field value
func (o *RpmModulemdDefaultsResponse) GetModule() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Module
}

// GetModuleOk returns a tuple with the Module field value
// and a boolean to check if the value has been set.
func (o *RpmModulemdDefaultsResponse) GetModuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Module, true
}

// SetModule sets field value
func (o *RpmModulemdDefaultsResponse) SetModule(v string) {
	o.Module = v
}

// GetStream returns the Stream field value
func (o *RpmModulemdDefaultsResponse) GetStream() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Stream
}

// GetStreamOk returns a tuple with the Stream field value
// and a boolean to check if the value has been set.
func (o *RpmModulemdDefaultsResponse) GetStreamOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stream, true
}

// SetStream sets field value
func (o *RpmModulemdDefaultsResponse) SetStream(v string) {
	o.Stream = v
}

// GetProfiles returns the Profiles field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *RpmModulemdDefaultsResponse) GetProfiles() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmModulemdDefaultsResponse) GetProfilesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Profiles) {
		return nil, false
	}
	return &o.Profiles, true
}

// SetProfiles sets field value
func (o *RpmModulemdDefaultsResponse) SetProfiles(v interface{}) {
	o.Profiles = v
}

func (o RpmModulemdDefaultsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RpmModulemdDefaultsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PulpHref) {
		toSerialize["pulp_href"] = o.PulpHref
	}
	if !IsNil(o.Prn) {
		toSerialize["prn"] = o.Prn
	}
	if !IsNil(o.PulpCreated) {
		toSerialize["pulp_created"] = o.PulpCreated
	}
	if !IsNil(o.PulpLastUpdated) {
		toSerialize["pulp_last_updated"] = o.PulpLastUpdated
	}
	if !IsNil(o.PulpLabels) {
		toSerialize["pulp_labels"] = o.PulpLabels
	}
	toSerialize["module"] = o.Module
	toSerialize["stream"] = o.Stream
	if o.Profiles != nil {
		toSerialize["profiles"] = o.Profiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RpmModulemdDefaultsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"module",
		"stream",
		"profiles",
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

	varRpmModulemdDefaultsResponse := _RpmModulemdDefaultsResponse{}

	err = json.Unmarshal(data, &varRpmModulemdDefaultsResponse)

	if err != nil {
		return err
	}

	*o = RpmModulemdDefaultsResponse(varRpmModulemdDefaultsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pulp_href")
		delete(additionalProperties, "prn")
		delete(additionalProperties, "pulp_created")
		delete(additionalProperties, "pulp_last_updated")
		delete(additionalProperties, "pulp_labels")
		delete(additionalProperties, "module")
		delete(additionalProperties, "stream")
		delete(additionalProperties, "profiles")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRpmModulemdDefaultsResponse struct {
	value *RpmModulemdDefaultsResponse
	isSet bool
}

func (v NullableRpmModulemdDefaultsResponse) Get() *RpmModulemdDefaultsResponse {
	return v.value
}

func (v *NullableRpmModulemdDefaultsResponse) Set(val *RpmModulemdDefaultsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRpmModulemdDefaultsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRpmModulemdDefaultsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRpmModulemdDefaultsResponse(val *RpmModulemdDefaultsResponse) *NullableRpmModulemdDefaultsResponse {
	return &NullableRpmModulemdDefaultsResponse{value: val, isSet: true}
}

func (v NullableRpmModulemdDefaultsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRpmModulemdDefaultsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


