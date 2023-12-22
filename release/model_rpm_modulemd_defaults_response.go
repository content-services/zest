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

// checks if the RpmModulemdDefaultsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RpmModulemdDefaultsResponse{}

// RpmModulemdDefaultsResponse ModulemdDefaults serializer.
type RpmModulemdDefaultsResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// Modulemd name.
	Module string `json:"module"`
	// Modulemd default stream.
	Stream string `json:"stream"`
	// Default profiles for modulemd streams.
	Profiles map[string]interface{} `json:"profiles"`
}

type _RpmModulemdDefaultsResponse RpmModulemdDefaultsResponse

// NewRpmModulemdDefaultsResponse instantiates a new RpmModulemdDefaultsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRpmModulemdDefaultsResponse(module string, stream string, profiles map[string]interface{}) *RpmModulemdDefaultsResponse {
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
func (o *RpmModulemdDefaultsResponse) GetProfiles() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value
// and a boolean to check if the value has been set.
func (o *RpmModulemdDefaultsResponse) GetProfilesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Profiles, true
}

// SetProfiles sets field value
func (o *RpmModulemdDefaultsResponse) SetProfiles(v map[string]interface{}) {
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
	if !IsNil(o.PulpCreated) {
		toSerialize["pulp_created"] = o.PulpCreated
	}
	toSerialize["module"] = o.Module
	toSerialize["stream"] = o.Stream
	toSerialize["profiles"] = o.Profiles
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

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRpmModulemdDefaultsResponse)

	if err != nil {
		return err
	}

	*o = RpmModulemdDefaultsResponse(varRpmModulemdDefaultsResponse)

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


