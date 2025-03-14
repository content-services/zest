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

// checks if the ContentSettingsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentSettingsResponse{}

// ContentSettingsResponse Serializer for information about content-app-settings for the pulp instance
type ContentSettingsResponse struct {
	// The CONTENT_ORIGIN setting for this Pulp instance
	ContentOrigin NullableString `json:"content_origin,omitempty"`
	// The CONTENT_PATH_PREFIX setting for this Pulp instance
	ContentPathPrefix string `json:"content_path_prefix"`
	AdditionalProperties map[string]interface{}
}

type _ContentSettingsResponse ContentSettingsResponse

// NewContentSettingsResponse instantiates a new ContentSettingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentSettingsResponse(contentPathPrefix string) *ContentSettingsResponse {
	this := ContentSettingsResponse{}
	this.ContentPathPrefix = contentPathPrefix
	return &this
}

// NewContentSettingsResponseWithDefaults instantiates a new ContentSettingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentSettingsResponseWithDefaults() *ContentSettingsResponse {
	this := ContentSettingsResponse{}
	return &this
}

// GetContentOrigin returns the ContentOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContentSettingsResponse) GetContentOrigin() string {
	if o == nil || IsNil(o.ContentOrigin.Get()) {
		var ret string
		return ret
	}
	return *o.ContentOrigin.Get()
}

// GetContentOriginOk returns a tuple with the ContentOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContentSettingsResponse) GetContentOriginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentOrigin.Get(), o.ContentOrigin.IsSet()
}

// HasContentOrigin returns a boolean if a field has been set.
func (o *ContentSettingsResponse) HasContentOrigin() bool {
	if o != nil && o.ContentOrigin.IsSet() {
		return true
	}

	return false
}

// SetContentOrigin gets a reference to the given NullableString and assigns it to the ContentOrigin field.
func (o *ContentSettingsResponse) SetContentOrigin(v string) {
	o.ContentOrigin.Set(&v)
}
// SetContentOriginNil sets the value for ContentOrigin to be an explicit nil
func (o *ContentSettingsResponse) SetContentOriginNil() {
	o.ContentOrigin.Set(nil)
}

// UnsetContentOrigin ensures that no value is present for ContentOrigin, not even an explicit nil
func (o *ContentSettingsResponse) UnsetContentOrigin() {
	o.ContentOrigin.Unset()
}

// GetContentPathPrefix returns the ContentPathPrefix field value
func (o *ContentSettingsResponse) GetContentPathPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentPathPrefix
}

// GetContentPathPrefixOk returns a tuple with the ContentPathPrefix field value
// and a boolean to check if the value has been set.
func (o *ContentSettingsResponse) GetContentPathPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentPathPrefix, true
}

// SetContentPathPrefix sets field value
func (o *ContentSettingsResponse) SetContentPathPrefix(v string) {
	o.ContentPathPrefix = v
}

func (o ContentSettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentSettingsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ContentOrigin.IsSet() {
		toSerialize["content_origin"] = o.ContentOrigin.Get()
	}
	toSerialize["content_path_prefix"] = o.ContentPathPrefix

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ContentSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"content_path_prefix",
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

	varContentSettingsResponse := _ContentSettingsResponse{}

	err = json.Unmarshal(data, &varContentSettingsResponse)

	if err != nil {
		return err
	}

	*o = ContentSettingsResponse(varContentSettingsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "content_origin")
		delete(additionalProperties, "content_path_prefix")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableContentSettingsResponse struct {
	value *ContentSettingsResponse
	isSet bool
}

func (v NullableContentSettingsResponse) Get() *ContentSettingsResponse {
	return v.value
}

func (v *NullableContentSettingsResponse) Set(val *ContentSettingsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContentSettingsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContentSettingsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentSettingsResponse(val *ContentSettingsResponse) *NullableContentSettingsResponse {
	return &NullableContentSettingsResponse{value: val, isSet: true}
}

func (v NullableContentSettingsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentSettingsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


