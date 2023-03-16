/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package zest/release/v3

import (
	"encoding/json"
)

// checks if the ContentSettingsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentSettingsResponse{}

// ContentSettingsResponse Serializer for information about content-app-settings for the pulp instance
type ContentSettingsResponse struct {
	// The CONTENT_ORIGIN setting for this Pulp instance
	ContentOrigin string `json:"content_origin"`
	// The CONTENT_PATH_PREFIX setting for this Pulp instance
	ContentPathPrefix string `json:"content_path_prefix"`
}

// NewContentSettingsResponse instantiates a new ContentSettingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentSettingsResponse(contentOrigin string, contentPathPrefix string) *ContentSettingsResponse {
	this := ContentSettingsResponse{}
	this.ContentOrigin = contentOrigin
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

// GetContentOrigin returns the ContentOrigin field value
func (o *ContentSettingsResponse) GetContentOrigin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentOrigin
}

// GetContentOriginOk returns a tuple with the ContentOrigin field value
// and a boolean to check if the value has been set.
func (o *ContentSettingsResponse) GetContentOriginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentOrigin, true
}

// SetContentOrigin sets field value
func (o *ContentSettingsResponse) SetContentOrigin(v string) {
	o.ContentOrigin = v
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
	toSerialize["content_origin"] = o.ContentOrigin
	toSerialize["content_path_prefix"] = o.ContentPathPrefix
	return toSerialize, nil
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


