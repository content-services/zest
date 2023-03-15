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
	"time"
)

// checks if the PatchedfileFileAlternateContentSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedfileFileAlternateContentSource{}

// PatchedfileFileAlternateContentSource Serializer for File alternate content source.
type PatchedfileFileAlternateContentSource struct {
	// Name of Alternate Content Source.
	Name *string `json:"name,omitempty"`
	// Date of last refresh of AlternateContentSource.
	LastRefreshed NullableTime `json:"last_refreshed,omitempty"`
	// List of paths that will be appended to the Remote url when searching for content.
	Paths []string `json:"paths,omitempty"`
	// The remote to provide alternate content source.
	Remote *string `json:"remote,omitempty"`
}

// NewPatchedfileFileAlternateContentSource instantiates a new PatchedfileFileAlternateContentSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedfileFileAlternateContentSource() *PatchedfileFileAlternateContentSource {
	this := PatchedfileFileAlternateContentSource{}
	return &this
}

// NewPatchedfileFileAlternateContentSourceWithDefaults instantiates a new PatchedfileFileAlternateContentSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedfileFileAlternateContentSourceWithDefaults() *PatchedfileFileAlternateContentSource {
	this := PatchedfileFileAlternateContentSource{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedfileFileAlternateContentSource) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedfileFileAlternateContentSource) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedfileFileAlternateContentSource) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedfileFileAlternateContentSource) SetName(v string) {
	o.Name = &v
}

// GetLastRefreshed returns the LastRefreshed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedfileFileAlternateContentSource) GetLastRefreshed() time.Time {
	if o == nil || IsNil(o.LastRefreshed.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastRefreshed.Get()
}

// GetLastRefreshedOk returns a tuple with the LastRefreshed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedfileFileAlternateContentSource) GetLastRefreshedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastRefreshed.Get(), o.LastRefreshed.IsSet()
}

// HasLastRefreshed returns a boolean if a field has been set.
func (o *PatchedfileFileAlternateContentSource) HasLastRefreshed() bool {
	if o != nil && o.LastRefreshed.IsSet() {
		return true
	}

	return false
}

// SetLastRefreshed gets a reference to the given NullableTime and assigns it to the LastRefreshed field.
func (o *PatchedfileFileAlternateContentSource) SetLastRefreshed(v time.Time) {
	o.LastRefreshed.Set(&v)
}
// SetLastRefreshedNil sets the value for LastRefreshed to be an explicit nil
func (o *PatchedfileFileAlternateContentSource) SetLastRefreshedNil() {
	o.LastRefreshed.Set(nil)
}

// UnsetLastRefreshed ensures that no value is present for LastRefreshed, not even an explicit nil
func (o *PatchedfileFileAlternateContentSource) UnsetLastRefreshed() {
	o.LastRefreshed.Unset()
}

// GetPaths returns the Paths field value if set, zero value otherwise.
func (o *PatchedfileFileAlternateContentSource) GetPaths() []string {
	if o == nil || IsNil(o.Paths) {
		var ret []string
		return ret
	}
	return o.Paths
}

// GetPathsOk returns a tuple with the Paths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedfileFileAlternateContentSource) GetPathsOk() ([]string, bool) {
	if o == nil || IsNil(o.Paths) {
		return nil, false
	}
	return o.Paths, true
}

// HasPaths returns a boolean if a field has been set.
func (o *PatchedfileFileAlternateContentSource) HasPaths() bool {
	if o != nil && !IsNil(o.Paths) {
		return true
	}

	return false
}

// SetPaths gets a reference to the given []string and assigns it to the Paths field.
func (o *PatchedfileFileAlternateContentSource) SetPaths(v []string) {
	o.Paths = v
}

// GetRemote returns the Remote field value if set, zero value otherwise.
func (o *PatchedfileFileAlternateContentSource) GetRemote() string {
	if o == nil || IsNil(o.Remote) {
		var ret string
		return ret
	}
	return *o.Remote
}

// GetRemoteOk returns a tuple with the Remote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedfileFileAlternateContentSource) GetRemoteOk() (*string, bool) {
	if o == nil || IsNil(o.Remote) {
		return nil, false
	}
	return o.Remote, true
}

// HasRemote returns a boolean if a field has been set.
func (o *PatchedfileFileAlternateContentSource) HasRemote() bool {
	if o != nil && !IsNil(o.Remote) {
		return true
	}

	return false
}

// SetRemote gets a reference to the given string and assigns it to the Remote field.
func (o *PatchedfileFileAlternateContentSource) SetRemote(v string) {
	o.Remote = &v
}

func (o PatchedfileFileAlternateContentSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedfileFileAlternateContentSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.LastRefreshed.IsSet() {
		toSerialize["last_refreshed"] = o.LastRefreshed.Get()
	}
	if !IsNil(o.Paths) {
		toSerialize["paths"] = o.Paths
	}
	if !IsNil(o.Remote) {
		toSerialize["remote"] = o.Remote
	}
	return toSerialize, nil
}

type NullablePatchedfileFileAlternateContentSource struct {
	value *PatchedfileFileAlternateContentSource
	isSet bool
}

func (v NullablePatchedfileFileAlternateContentSource) Get() *PatchedfileFileAlternateContentSource {
	return v.value
}

func (v *NullablePatchedfileFileAlternateContentSource) Set(val *PatchedfileFileAlternateContentSource) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedfileFileAlternateContentSource) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedfileFileAlternateContentSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedfileFileAlternateContentSource(val *PatchedfileFileAlternateContentSource) *NullablePatchedfileFileAlternateContentSource {
	return &NullablePatchedfileFileAlternateContentSource{value: val, isSet: true}
}

func (v NullablePatchedfileFileAlternateContentSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedfileFileAlternateContentSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


