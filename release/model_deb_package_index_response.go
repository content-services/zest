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
)

// checks if the DebPackageIndexResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DebPackageIndexResponse{}

// DebPackageIndexResponse A serializer for PackageIndex.
type DebPackageIndexResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// A dict mapping relative paths inside the Content to the correspondingArtifact URLs. E.g.: {'relative/path': '/artifacts/1/'
	Artifacts map[string]interface{} `json:"artifacts"`
	// Release this index file belongs to.
	Release string `json:"release"`
	// Component of the component - architecture combination.
	Component *string `json:"component,omitempty"`
	// Architecture of the component - architecture combination.
	Architecture *string `json:"architecture,omitempty"`
	// Path of file relative to url.
	RelativePath *string `json:"relative_path,omitempty"`
}

// NewDebPackageIndexResponse instantiates a new DebPackageIndexResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDebPackageIndexResponse(artifacts map[string]interface{}, release string) *DebPackageIndexResponse {
	this := DebPackageIndexResponse{}
	this.Artifacts = artifacts
	this.Release = release
	return &this
}

// NewDebPackageIndexResponseWithDefaults instantiates a new DebPackageIndexResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDebPackageIndexResponseWithDefaults() *DebPackageIndexResponse {
	this := DebPackageIndexResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *DebPackageIndexResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebPackageIndexResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *DebPackageIndexResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *DebPackageIndexResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *DebPackageIndexResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebPackageIndexResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *DebPackageIndexResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *DebPackageIndexResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetArtifacts returns the Artifacts field value
func (o *DebPackageIndexResponse) GetArtifacts() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Artifacts
}

// GetArtifactsOk returns a tuple with the Artifacts field value
// and a boolean to check if the value has been set.
func (o *DebPackageIndexResponse) GetArtifactsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Artifacts, true
}

// SetArtifacts sets field value
func (o *DebPackageIndexResponse) SetArtifacts(v map[string]interface{}) {
	o.Artifacts = v
}

// GetRelease returns the Release field value
func (o *DebPackageIndexResponse) GetRelease() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Release
}

// GetReleaseOk returns a tuple with the Release field value
// and a boolean to check if the value has been set.
func (o *DebPackageIndexResponse) GetReleaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Release, true
}

// SetRelease sets field value
func (o *DebPackageIndexResponse) SetRelease(v string) {
	o.Release = v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *DebPackageIndexResponse) GetComponent() string {
	if o == nil || IsNil(o.Component) {
		var ret string
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebPackageIndexResponse) GetComponentOk() (*string, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *DebPackageIndexResponse) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given string and assigns it to the Component field.
func (o *DebPackageIndexResponse) SetComponent(v string) {
	o.Component = &v
}

// GetArchitecture returns the Architecture field value if set, zero value otherwise.
func (o *DebPackageIndexResponse) GetArchitecture() string {
	if o == nil || IsNil(o.Architecture) {
		var ret string
		return ret
	}
	return *o.Architecture
}

// GetArchitectureOk returns a tuple with the Architecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebPackageIndexResponse) GetArchitectureOk() (*string, bool) {
	if o == nil || IsNil(o.Architecture) {
		return nil, false
	}
	return o.Architecture, true
}

// HasArchitecture returns a boolean if a field has been set.
func (o *DebPackageIndexResponse) HasArchitecture() bool {
	if o != nil && !IsNil(o.Architecture) {
		return true
	}

	return false
}

// SetArchitecture gets a reference to the given string and assigns it to the Architecture field.
func (o *DebPackageIndexResponse) SetArchitecture(v string) {
	o.Architecture = &v
}

// GetRelativePath returns the RelativePath field value if set, zero value otherwise.
func (o *DebPackageIndexResponse) GetRelativePath() string {
	if o == nil || IsNil(o.RelativePath) {
		var ret string
		return ret
	}
	return *o.RelativePath
}

// GetRelativePathOk returns a tuple with the RelativePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebPackageIndexResponse) GetRelativePathOk() (*string, bool) {
	if o == nil || IsNil(o.RelativePath) {
		return nil, false
	}
	return o.RelativePath, true
}

// HasRelativePath returns a boolean if a field has been set.
func (o *DebPackageIndexResponse) HasRelativePath() bool {
	if o != nil && !IsNil(o.RelativePath) {
		return true
	}

	return false
}

// SetRelativePath gets a reference to the given string and assigns it to the RelativePath field.
func (o *DebPackageIndexResponse) SetRelativePath(v string) {
	o.RelativePath = &v
}

func (o DebPackageIndexResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DebPackageIndexResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: pulp_href is readOnly
	// skip: pulp_created is readOnly
	toSerialize["artifacts"] = o.Artifacts
	toSerialize["release"] = o.Release
	if !IsNil(o.Component) {
		toSerialize["component"] = o.Component
	}
	if !IsNil(o.Architecture) {
		toSerialize["architecture"] = o.Architecture
	}
	if !IsNil(o.RelativePath) {
		toSerialize["relative_path"] = o.RelativePath
	}
	return toSerialize, nil
}

type NullableDebPackageIndexResponse struct {
	value *DebPackageIndexResponse
	isSet bool
}

func (v NullableDebPackageIndexResponse) Get() *DebPackageIndexResponse {
	return v.value
}

func (v *NullableDebPackageIndexResponse) Set(val *DebPackageIndexResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDebPackageIndexResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDebPackageIndexResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDebPackageIndexResponse(val *DebPackageIndexResponse) *NullableDebPackageIndexResponse {
	return &NullableDebPackageIndexResponse{value: val, isSet: true}
}

func (v NullableDebPackageIndexResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDebPackageIndexResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


