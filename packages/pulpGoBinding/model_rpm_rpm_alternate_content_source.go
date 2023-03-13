/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pulpGoBinding

import (
	"encoding/json"
	"time"
)

// checks if the RpmRpmAlternateContentSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RpmRpmAlternateContentSource{}

// RpmRpmAlternateContentSource Serializer for RPM alternate content source.
type RpmRpmAlternateContentSource struct {
	// Name of Alternate Content Source.
	Name string `json:"name"`
	// Date of last refresh of AlternateContentSource.
	LastRefreshed NullableTime `json:"last_refreshed,omitempty"`
	// List of paths that will be appended to the Remote url when searching for content.
	Paths []string `json:"paths,omitempty"`
	// The remote to provide alternate content source.
	Remote string `json:"remote"`
}

// NewRpmRpmAlternateContentSource instantiates a new RpmRpmAlternateContentSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRpmRpmAlternateContentSource(name string, remote string) *RpmRpmAlternateContentSource {
	this := RpmRpmAlternateContentSource{}
	this.Name = name
	this.Remote = remote
	return &this
}

// NewRpmRpmAlternateContentSourceWithDefaults instantiates a new RpmRpmAlternateContentSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRpmRpmAlternateContentSourceWithDefaults() *RpmRpmAlternateContentSource {
	this := RpmRpmAlternateContentSource{}
	return &this
}

// GetName returns the Name field value
func (o *RpmRpmAlternateContentSource) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RpmRpmAlternateContentSource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RpmRpmAlternateContentSource) SetName(v string) {
	o.Name = v
}

// GetLastRefreshed returns the LastRefreshed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RpmRpmAlternateContentSource) GetLastRefreshed() time.Time {
	if o == nil || IsNil(o.LastRefreshed.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastRefreshed.Get()
}

// GetLastRefreshedOk returns a tuple with the LastRefreshed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RpmRpmAlternateContentSource) GetLastRefreshedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastRefreshed.Get(), o.LastRefreshed.IsSet()
}

// HasLastRefreshed returns a boolean if a field has been set.
func (o *RpmRpmAlternateContentSource) HasLastRefreshed() bool {
	if o != nil && o.LastRefreshed.IsSet() {
		return true
	}

	return false
}

// SetLastRefreshed gets a reference to the given NullableTime and assigns it to the LastRefreshed field.
func (o *RpmRpmAlternateContentSource) SetLastRefreshed(v time.Time) {
	o.LastRefreshed.Set(&v)
}
// SetLastRefreshedNil sets the value for LastRefreshed to be an explicit nil
func (o *RpmRpmAlternateContentSource) SetLastRefreshedNil() {
	o.LastRefreshed.Set(nil)
}

// UnsetLastRefreshed ensures that no value is present for LastRefreshed, not even an explicit nil
func (o *RpmRpmAlternateContentSource) UnsetLastRefreshed() {
	o.LastRefreshed.Unset()
}

// GetPaths returns the Paths field value if set, zero value otherwise.
func (o *RpmRpmAlternateContentSource) GetPaths() []string {
	if o == nil || IsNil(o.Paths) {
		var ret []string
		return ret
	}
	return o.Paths
}

// GetPathsOk returns a tuple with the Paths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RpmRpmAlternateContentSource) GetPathsOk() ([]string, bool) {
	if o == nil || IsNil(o.Paths) {
		return nil, false
	}
	return o.Paths, true
}

// HasPaths returns a boolean if a field has been set.
func (o *RpmRpmAlternateContentSource) HasPaths() bool {
	if o != nil && !IsNil(o.Paths) {
		return true
	}

	return false
}

// SetPaths gets a reference to the given []string and assigns it to the Paths field.
func (o *RpmRpmAlternateContentSource) SetPaths(v []string) {
	o.Paths = v
}

// GetRemote returns the Remote field value
func (o *RpmRpmAlternateContentSource) GetRemote() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Remote
}

// GetRemoteOk returns a tuple with the Remote field value
// and a boolean to check if the value has been set.
func (o *RpmRpmAlternateContentSource) GetRemoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Remote, true
}

// SetRemote sets field value
func (o *RpmRpmAlternateContentSource) SetRemote(v string) {
	o.Remote = v
}

func (o RpmRpmAlternateContentSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RpmRpmAlternateContentSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if o.LastRefreshed.IsSet() {
		toSerialize["last_refreshed"] = o.LastRefreshed.Get()
	}
	if !IsNil(o.Paths) {
		toSerialize["paths"] = o.Paths
	}
	toSerialize["remote"] = o.Remote
	return toSerialize, nil
}

type NullableRpmRpmAlternateContentSource struct {
	value *RpmRpmAlternateContentSource
	isSet bool
}

func (v NullableRpmRpmAlternateContentSource) Get() *RpmRpmAlternateContentSource {
	return v.value
}

func (v *NullableRpmRpmAlternateContentSource) Set(val *RpmRpmAlternateContentSource) {
	v.value = val
	v.isSet = true
}

func (v NullableRpmRpmAlternateContentSource) IsSet() bool {
	return v.isSet
}

func (v *NullableRpmRpmAlternateContentSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRpmRpmAlternateContentSource(val *RpmRpmAlternateContentSource) *NullableRpmRpmAlternateContentSource {
	return &NullableRpmRpmAlternateContentSource{value: val, isSet: true}
}

func (v NullableRpmRpmAlternateContentSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRpmRpmAlternateContentSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

