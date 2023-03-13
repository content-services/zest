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

// checks if the PatchedrpmRpmAlternateContentSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedrpmRpmAlternateContentSource{}

// PatchedrpmRpmAlternateContentSource Serializer for RPM alternate content source.
type PatchedrpmRpmAlternateContentSource struct {
	// Name of Alternate Content Source.
	Name *string `json:"name,omitempty"`
	// Date of last refresh of AlternateContentSource.
	LastRefreshed NullableTime `json:"last_refreshed,omitempty"`
	// List of paths that will be appended to the Remote url when searching for content.
	Paths []string `json:"paths,omitempty"`
	// The remote to provide alternate content source.
	Remote *string `json:"remote,omitempty"`
}

// NewPatchedrpmRpmAlternateContentSource instantiates a new PatchedrpmRpmAlternateContentSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedrpmRpmAlternateContentSource() *PatchedrpmRpmAlternateContentSource {
	this := PatchedrpmRpmAlternateContentSource{}
	return &this
}

// NewPatchedrpmRpmAlternateContentSourceWithDefaults instantiates a new PatchedrpmRpmAlternateContentSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedrpmRpmAlternateContentSourceWithDefaults() *PatchedrpmRpmAlternateContentSource {
	this := PatchedrpmRpmAlternateContentSource{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedrpmRpmAlternateContentSource) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedrpmRpmAlternateContentSource) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedrpmRpmAlternateContentSource) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedrpmRpmAlternateContentSource) SetName(v string) {
	o.Name = &v
}

// GetLastRefreshed returns the LastRefreshed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedrpmRpmAlternateContentSource) GetLastRefreshed() time.Time {
	if o == nil || IsNil(o.LastRefreshed.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastRefreshed.Get()
}

// GetLastRefreshedOk returns a tuple with the LastRefreshed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedrpmRpmAlternateContentSource) GetLastRefreshedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastRefreshed.Get(), o.LastRefreshed.IsSet()
}

// HasLastRefreshed returns a boolean if a field has been set.
func (o *PatchedrpmRpmAlternateContentSource) HasLastRefreshed() bool {
	if o != nil && o.LastRefreshed.IsSet() {
		return true
	}

	return false
}

// SetLastRefreshed gets a reference to the given NullableTime and assigns it to the LastRefreshed field.
func (o *PatchedrpmRpmAlternateContentSource) SetLastRefreshed(v time.Time) {
	o.LastRefreshed.Set(&v)
}
// SetLastRefreshedNil sets the value for LastRefreshed to be an explicit nil
func (o *PatchedrpmRpmAlternateContentSource) SetLastRefreshedNil() {
	o.LastRefreshed.Set(nil)
}

// UnsetLastRefreshed ensures that no value is present for LastRefreshed, not even an explicit nil
func (o *PatchedrpmRpmAlternateContentSource) UnsetLastRefreshed() {
	o.LastRefreshed.Unset()
}

// GetPaths returns the Paths field value if set, zero value otherwise.
func (o *PatchedrpmRpmAlternateContentSource) GetPaths() []string {
	if o == nil || IsNil(o.Paths) {
		var ret []string
		return ret
	}
	return o.Paths
}

// GetPathsOk returns a tuple with the Paths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedrpmRpmAlternateContentSource) GetPathsOk() ([]string, bool) {
	if o == nil || IsNil(o.Paths) {
		return nil, false
	}
	return o.Paths, true
}

// HasPaths returns a boolean if a field has been set.
func (o *PatchedrpmRpmAlternateContentSource) HasPaths() bool {
	if o != nil && !IsNil(o.Paths) {
		return true
	}

	return false
}

// SetPaths gets a reference to the given []string and assigns it to the Paths field.
func (o *PatchedrpmRpmAlternateContentSource) SetPaths(v []string) {
	o.Paths = v
}

// GetRemote returns the Remote field value if set, zero value otherwise.
func (o *PatchedrpmRpmAlternateContentSource) GetRemote() string {
	if o == nil || IsNil(o.Remote) {
		var ret string
		return ret
	}
	return *o.Remote
}

// GetRemoteOk returns a tuple with the Remote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedrpmRpmAlternateContentSource) GetRemoteOk() (*string, bool) {
	if o == nil || IsNil(o.Remote) {
		return nil, false
	}
	return o.Remote, true
}

// HasRemote returns a boolean if a field has been set.
func (o *PatchedrpmRpmAlternateContentSource) HasRemote() bool {
	if o != nil && !IsNil(o.Remote) {
		return true
	}

	return false
}

// SetRemote gets a reference to the given string and assigns it to the Remote field.
func (o *PatchedrpmRpmAlternateContentSource) SetRemote(v string) {
	o.Remote = &v
}

func (o PatchedrpmRpmAlternateContentSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedrpmRpmAlternateContentSource) ToMap() (map[string]interface{}, error) {
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

type NullablePatchedrpmRpmAlternateContentSource struct {
	value *PatchedrpmRpmAlternateContentSource
	isSet bool
}

func (v NullablePatchedrpmRpmAlternateContentSource) Get() *PatchedrpmRpmAlternateContentSource {
	return v.value
}

func (v *NullablePatchedrpmRpmAlternateContentSource) Set(val *PatchedrpmRpmAlternateContentSource) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedrpmRpmAlternateContentSource) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedrpmRpmAlternateContentSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedrpmRpmAlternateContentSource(val *PatchedrpmRpmAlternateContentSource) *NullablePatchedrpmRpmAlternateContentSource {
	return &NullablePatchedrpmRpmAlternateContentSource{value: val, isSet: true}
}

func (v NullablePatchedrpmRpmAlternateContentSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedrpmRpmAlternateContentSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

