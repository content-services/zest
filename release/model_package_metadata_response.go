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

// checks if the PackageMetadataResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackageMetadataResponse{}

// PackageMetadataResponse A Serializer for a package's metadata.
type PackageMetadataResponse struct {
	// Cache value from last PyPI sync
	LastSerial int64 `json:"last_serial"`
	// Core metadata of the package
	Info interface{} `json:"info"`
	// List of all the releases of the package
	Releases interface{} `json:"releases"`
	Urls interface{} `json:"urls"`
	AdditionalProperties map[string]interface{}
}

type _PackageMetadataResponse PackageMetadataResponse

// NewPackageMetadataResponse instantiates a new PackageMetadataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageMetadataResponse(lastSerial int64, info interface{}, releases interface{}, urls interface{}) *PackageMetadataResponse {
	this := PackageMetadataResponse{}
	this.LastSerial = lastSerial
	this.Info = info
	this.Releases = releases
	this.Urls = urls
	return &this
}

// NewPackageMetadataResponseWithDefaults instantiates a new PackageMetadataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageMetadataResponseWithDefaults() *PackageMetadataResponse {
	this := PackageMetadataResponse{}
	return &this
}

// GetLastSerial returns the LastSerial field value
func (o *PackageMetadataResponse) GetLastSerial() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.LastSerial
}

// GetLastSerialOk returns a tuple with the LastSerial field value
// and a boolean to check if the value has been set.
func (o *PackageMetadataResponse) GetLastSerialOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastSerial, true
}

// SetLastSerial sets field value
func (o *PackageMetadataResponse) SetLastSerial(v int64) {
	o.LastSerial = v
}

// GetInfo returns the Info field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PackageMetadataResponse) GetInfo() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Info
}

// GetInfoOk returns a tuple with the Info field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PackageMetadataResponse) GetInfoOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return &o.Info, true
}

// SetInfo sets field value
func (o *PackageMetadataResponse) SetInfo(v interface{}) {
	o.Info = v
}

// GetReleases returns the Releases field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PackageMetadataResponse) GetReleases() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Releases
}

// GetReleasesOk returns a tuple with the Releases field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PackageMetadataResponse) GetReleasesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Releases) {
		return nil, false
	}
	return &o.Releases, true
}

// SetReleases sets field value
func (o *PackageMetadataResponse) SetReleases(v interface{}) {
	o.Releases = v
}

// GetUrls returns the Urls field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PackageMetadataResponse) GetUrls() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PackageMetadataResponse) GetUrlsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Urls) {
		return nil, false
	}
	return &o.Urls, true
}

// SetUrls sets field value
func (o *PackageMetadataResponse) SetUrls(v interface{}) {
	o.Urls = v
}

func (o PackageMetadataResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PackageMetadataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["last_serial"] = o.LastSerial
	if o.Info != nil {
		toSerialize["info"] = o.Info
	}
	if o.Releases != nil {
		toSerialize["releases"] = o.Releases
	}
	if o.Urls != nil {
		toSerialize["urls"] = o.Urls
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PackageMetadataResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"last_serial",
		"info",
		"releases",
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

	varPackageMetadataResponse := _PackageMetadataResponse{}

	err = json.Unmarshal(data, &varPackageMetadataResponse)

	if err != nil {
		return err
	}

	*o = PackageMetadataResponse(varPackageMetadataResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "last_serial")
		delete(additionalProperties, "info")
		delete(additionalProperties, "releases")
		delete(additionalProperties, "urls")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePackageMetadataResponse struct {
	value *PackageMetadataResponse
	isSet bool
}

func (v NullablePackageMetadataResponse) Get() *PackageMetadataResponse {
	return v.value
}

func (v *NullablePackageMetadataResponse) Set(val *PackageMetadataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageMetadataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageMetadataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageMetadataResponse(val *PackageMetadataResponse) *NullablePackageMetadataResponse {
	return &NullablePackageMetadataResponse{value: val, isSet: true}
}

func (v NullablePackageMetadataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageMetadataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


