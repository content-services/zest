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

// checks if the FileFileAlternateContentSourceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileFileAlternateContentSourceResponse{}

// FileFileAlternateContentSourceResponse Serializer for File alternate content source.
type FileFileAlternateContentSourceResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// Name of Alternate Content Source.
	Name string `json:"name"`
	// Date of last refresh of AlternateContentSource.
	LastRefreshed NullableTime `json:"last_refreshed,omitempty"`
	// List of paths that will be appended to the Remote url when searching for content.
	Paths []string `json:"paths,omitempty"`
	// The remote to provide alternate content source.
	Remote string `json:"remote"`
}

// NewFileFileAlternateContentSourceResponse instantiates a new FileFileAlternateContentSourceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileFileAlternateContentSourceResponse(name string, remote string) *FileFileAlternateContentSourceResponse {
	this := FileFileAlternateContentSourceResponse{}
	this.Name = name
	this.Remote = remote
	return &this
}

// NewFileFileAlternateContentSourceResponseWithDefaults instantiates a new FileFileAlternateContentSourceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileFileAlternateContentSourceResponseWithDefaults() *FileFileAlternateContentSourceResponse {
	this := FileFileAlternateContentSourceResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *FileFileAlternateContentSourceResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileFileAlternateContentSourceResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *FileFileAlternateContentSourceResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *FileFileAlternateContentSourceResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *FileFileAlternateContentSourceResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileFileAlternateContentSourceResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *FileFileAlternateContentSourceResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *FileFileAlternateContentSourceResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetName returns the Name field value
func (o *FileFileAlternateContentSourceResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FileFileAlternateContentSourceResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FileFileAlternateContentSourceResponse) SetName(v string) {
	o.Name = v
}

// GetLastRefreshed returns the LastRefreshed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FileFileAlternateContentSourceResponse) GetLastRefreshed() time.Time {
	if o == nil || IsNil(o.LastRefreshed.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastRefreshed.Get()
}

// GetLastRefreshedOk returns a tuple with the LastRefreshed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FileFileAlternateContentSourceResponse) GetLastRefreshedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastRefreshed.Get(), o.LastRefreshed.IsSet()
}

// HasLastRefreshed returns a boolean if a field has been set.
func (o *FileFileAlternateContentSourceResponse) HasLastRefreshed() bool {
	if o != nil && o.LastRefreshed.IsSet() {
		return true
	}

	return false
}

// SetLastRefreshed gets a reference to the given NullableTime and assigns it to the LastRefreshed field.
func (o *FileFileAlternateContentSourceResponse) SetLastRefreshed(v time.Time) {
	o.LastRefreshed.Set(&v)
}
// SetLastRefreshedNil sets the value for LastRefreshed to be an explicit nil
func (o *FileFileAlternateContentSourceResponse) SetLastRefreshedNil() {
	o.LastRefreshed.Set(nil)
}

// UnsetLastRefreshed ensures that no value is present for LastRefreshed, not even an explicit nil
func (o *FileFileAlternateContentSourceResponse) UnsetLastRefreshed() {
	o.LastRefreshed.Unset()
}

// GetPaths returns the Paths field value if set, zero value otherwise.
func (o *FileFileAlternateContentSourceResponse) GetPaths() []string {
	if o == nil || IsNil(o.Paths) {
		var ret []string
		return ret
	}
	return o.Paths
}

// GetPathsOk returns a tuple with the Paths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileFileAlternateContentSourceResponse) GetPathsOk() ([]string, bool) {
	if o == nil || IsNil(o.Paths) {
		return nil, false
	}
	return o.Paths, true
}

// HasPaths returns a boolean if a field has been set.
func (o *FileFileAlternateContentSourceResponse) HasPaths() bool {
	if o != nil && !IsNil(o.Paths) {
		return true
	}

	return false
}

// SetPaths gets a reference to the given []string and assigns it to the Paths field.
func (o *FileFileAlternateContentSourceResponse) SetPaths(v []string) {
	o.Paths = v
}

// GetRemote returns the Remote field value
func (o *FileFileAlternateContentSourceResponse) GetRemote() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Remote
}

// GetRemoteOk returns a tuple with the Remote field value
// and a boolean to check if the value has been set.
func (o *FileFileAlternateContentSourceResponse) GetRemoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Remote, true
}

// SetRemote sets field value
func (o *FileFileAlternateContentSourceResponse) SetRemote(v string) {
	o.Remote = v
}

func (o FileFileAlternateContentSourceResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileFileAlternateContentSourceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: pulp_href is readOnly
	// skip: pulp_created is readOnly
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

type NullableFileFileAlternateContentSourceResponse struct {
	value *FileFileAlternateContentSourceResponse
	isSet bool
}

func (v NullableFileFileAlternateContentSourceResponse) Get() *FileFileAlternateContentSourceResponse {
	return v.value
}

func (v *NullableFileFileAlternateContentSourceResponse) Set(val *FileFileAlternateContentSourceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFileFileAlternateContentSourceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFileFileAlternateContentSourceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileFileAlternateContentSourceResponse(val *FileFileAlternateContentSourceResponse) *NullableFileFileAlternateContentSourceResponse {
	return &NullableFileFileAlternateContentSourceResponse{value: val, isSet: true}
}

func (v NullableFileFileAlternateContentSourceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileFileAlternateContentSourceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


