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

// checks if the ProfileArtifactResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileArtifactResponse{}

// ProfileArtifactResponse struct for ProfileArtifactResponse
type ProfileArtifactResponse struct {
	Urls map[string]string `json:"urls"`
	AdditionalProperties map[string]interface{}
}

type _ProfileArtifactResponse ProfileArtifactResponse

// NewProfileArtifactResponse instantiates a new ProfileArtifactResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileArtifactResponse(urls map[string]string) *ProfileArtifactResponse {
	this := ProfileArtifactResponse{}
	this.Urls = urls
	return &this
}

// NewProfileArtifactResponseWithDefaults instantiates a new ProfileArtifactResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileArtifactResponseWithDefaults() *ProfileArtifactResponse {
	this := ProfileArtifactResponse{}
	return &this
}

// GetUrls returns the Urls field value
func (o *ProfileArtifactResponse) GetUrls() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value
// and a boolean to check if the value has been set.
func (o *ProfileArtifactResponse) GetUrlsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Urls, true
}

// SetUrls sets field value
func (o *ProfileArtifactResponse) SetUrls(v map[string]string) {
	o.Urls = v
}

func (o ProfileArtifactResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileArtifactResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["urls"] = o.Urls

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProfileArtifactResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"urls",
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

	varProfileArtifactResponse := _ProfileArtifactResponse{}

	err = json.Unmarshal(data, &varProfileArtifactResponse)

	if err != nil {
		return err
	}

	*o = ProfileArtifactResponse(varProfileArtifactResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "urls")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProfileArtifactResponse struct {
	value *ProfileArtifactResponse
	isSet bool
}

func (v NullableProfileArtifactResponse) Get() *ProfileArtifactResponse {
	return v.value
}

func (v *NullableProfileArtifactResponse) Set(val *ProfileArtifactResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileArtifactResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileArtifactResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileArtifactResponse(val *ProfileArtifactResponse) *NullableProfileArtifactResponse {
	return &NullableProfileArtifactResponse{value: val, isSet: true}
}

func (v NullableProfileArtifactResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileArtifactResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


