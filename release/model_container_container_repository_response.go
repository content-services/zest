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

// checks if the ContainerContainerRepositoryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerContainerRepositoryResponse{}

// ContainerContainerRepositoryResponse Serializer for Container Repositories.
type ContainerContainerRepositoryResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	VersionsHref *string `json:"versions_href,omitempty"`
	PulpLabels *map[string]string `json:"pulp_labels,omitempty"`
	LatestVersionHref *string `json:"latest_version_href,omitempty"`
	// A unique name for this repository.
	Name string `json:"name"`
	// An optional description.
	Description NullableString `json:"description,omitempty"`
	// Retain X versions of the repository. Default is null which retains all versions.
	RetainRepoVersions NullableInt64 `json:"retain_repo_versions,omitempty"`
	// An optional remote to use by default when syncing.
	Remote NullableString `json:"remote,omitempty"`
	// A reference to an associated signing service.
	ManifestSigningService NullableString `json:"manifest_signing_service,omitempty"`
}

// NewContainerContainerRepositoryResponse instantiates a new ContainerContainerRepositoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerContainerRepositoryResponse(name string) *ContainerContainerRepositoryResponse {
	this := ContainerContainerRepositoryResponse{}
	this.Name = name
	return &this
}

// NewContainerContainerRepositoryResponseWithDefaults instantiates a new ContainerContainerRepositoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerContainerRepositoryResponseWithDefaults() *ContainerContainerRepositoryResponse {
	this := ContainerContainerRepositoryResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *ContainerContainerRepositoryResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerRepositoryResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *ContainerContainerRepositoryResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *ContainerContainerRepositoryResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *ContainerContainerRepositoryResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerRepositoryResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *ContainerContainerRepositoryResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *ContainerContainerRepositoryResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetVersionsHref returns the VersionsHref field value if set, zero value otherwise.
func (o *ContainerContainerRepositoryResponse) GetVersionsHref() string {
	if o == nil || IsNil(o.VersionsHref) {
		var ret string
		return ret
	}
	return *o.VersionsHref
}

// GetVersionsHrefOk returns a tuple with the VersionsHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerRepositoryResponse) GetVersionsHrefOk() (*string, bool) {
	if o == nil || IsNil(o.VersionsHref) {
		return nil, false
	}
	return o.VersionsHref, true
}

// HasVersionsHref returns a boolean if a field has been set.
func (o *ContainerContainerRepositoryResponse) HasVersionsHref() bool {
	if o != nil && !IsNil(o.VersionsHref) {
		return true
	}

	return false
}

// SetVersionsHref gets a reference to the given string and assigns it to the VersionsHref field.
func (o *ContainerContainerRepositoryResponse) SetVersionsHref(v string) {
	o.VersionsHref = &v
}

// GetPulpLabels returns the PulpLabels field value if set, zero value otherwise.
func (o *ContainerContainerRepositoryResponse) GetPulpLabels() map[string]string {
	if o == nil || IsNil(o.PulpLabels) {
		var ret map[string]string
		return ret
	}
	return *o.PulpLabels
}

// GetPulpLabelsOk returns a tuple with the PulpLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerRepositoryResponse) GetPulpLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PulpLabels) {
		return nil, false
	}
	return o.PulpLabels, true
}

// HasPulpLabels returns a boolean if a field has been set.
func (o *ContainerContainerRepositoryResponse) HasPulpLabels() bool {
	if o != nil && !IsNil(o.PulpLabels) {
		return true
	}

	return false
}

// SetPulpLabels gets a reference to the given map[string]string and assigns it to the PulpLabels field.
func (o *ContainerContainerRepositoryResponse) SetPulpLabels(v map[string]string) {
	o.PulpLabels = &v
}

// GetLatestVersionHref returns the LatestVersionHref field value if set, zero value otherwise.
func (o *ContainerContainerRepositoryResponse) GetLatestVersionHref() string {
	if o == nil || IsNil(o.LatestVersionHref) {
		var ret string
		return ret
	}
	return *o.LatestVersionHref
}

// GetLatestVersionHrefOk returns a tuple with the LatestVersionHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerRepositoryResponse) GetLatestVersionHrefOk() (*string, bool) {
	if o == nil || IsNil(o.LatestVersionHref) {
		return nil, false
	}
	return o.LatestVersionHref, true
}

// HasLatestVersionHref returns a boolean if a field has been set.
func (o *ContainerContainerRepositoryResponse) HasLatestVersionHref() bool {
	if o != nil && !IsNil(o.LatestVersionHref) {
		return true
	}

	return false
}

// SetLatestVersionHref gets a reference to the given string and assigns it to the LatestVersionHref field.
func (o *ContainerContainerRepositoryResponse) SetLatestVersionHref(v string) {
	o.LatestVersionHref = &v
}

// GetName returns the Name field value
func (o *ContainerContainerRepositoryResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContainerContainerRepositoryResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContainerContainerRepositoryResponse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerContainerRepositoryResponse) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerContainerRepositoryResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ContainerContainerRepositoryResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ContainerContainerRepositoryResponse) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ContainerContainerRepositoryResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ContainerContainerRepositoryResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetRetainRepoVersions returns the RetainRepoVersions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerContainerRepositoryResponse) GetRetainRepoVersions() int64 {
	if o == nil || IsNil(o.RetainRepoVersions.Get()) {
		var ret int64
		return ret
	}
	return *o.RetainRepoVersions.Get()
}

// GetRetainRepoVersionsOk returns a tuple with the RetainRepoVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerContainerRepositoryResponse) GetRetainRepoVersionsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RetainRepoVersions.Get(), o.RetainRepoVersions.IsSet()
}

// HasRetainRepoVersions returns a boolean if a field has been set.
func (o *ContainerContainerRepositoryResponse) HasRetainRepoVersions() bool {
	if o != nil && o.RetainRepoVersions.IsSet() {
		return true
	}

	return false
}

// SetRetainRepoVersions gets a reference to the given NullableInt64 and assigns it to the RetainRepoVersions field.
func (o *ContainerContainerRepositoryResponse) SetRetainRepoVersions(v int64) {
	o.RetainRepoVersions.Set(&v)
}
// SetRetainRepoVersionsNil sets the value for RetainRepoVersions to be an explicit nil
func (o *ContainerContainerRepositoryResponse) SetRetainRepoVersionsNil() {
	o.RetainRepoVersions.Set(nil)
}

// UnsetRetainRepoVersions ensures that no value is present for RetainRepoVersions, not even an explicit nil
func (o *ContainerContainerRepositoryResponse) UnsetRetainRepoVersions() {
	o.RetainRepoVersions.Unset()
}

// GetRemote returns the Remote field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerContainerRepositoryResponse) GetRemote() string {
	if o == nil || IsNil(o.Remote.Get()) {
		var ret string
		return ret
	}
	return *o.Remote.Get()
}

// GetRemoteOk returns a tuple with the Remote field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerContainerRepositoryResponse) GetRemoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Remote.Get(), o.Remote.IsSet()
}

// HasRemote returns a boolean if a field has been set.
func (o *ContainerContainerRepositoryResponse) HasRemote() bool {
	if o != nil && o.Remote.IsSet() {
		return true
	}

	return false
}

// SetRemote gets a reference to the given NullableString and assigns it to the Remote field.
func (o *ContainerContainerRepositoryResponse) SetRemote(v string) {
	o.Remote.Set(&v)
}
// SetRemoteNil sets the value for Remote to be an explicit nil
func (o *ContainerContainerRepositoryResponse) SetRemoteNil() {
	o.Remote.Set(nil)
}

// UnsetRemote ensures that no value is present for Remote, not even an explicit nil
func (o *ContainerContainerRepositoryResponse) UnsetRemote() {
	o.Remote.Unset()
}

// GetManifestSigningService returns the ManifestSigningService field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerContainerRepositoryResponse) GetManifestSigningService() string {
	if o == nil || IsNil(o.ManifestSigningService.Get()) {
		var ret string
		return ret
	}
	return *o.ManifestSigningService.Get()
}

// GetManifestSigningServiceOk returns a tuple with the ManifestSigningService field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerContainerRepositoryResponse) GetManifestSigningServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ManifestSigningService.Get(), o.ManifestSigningService.IsSet()
}

// HasManifestSigningService returns a boolean if a field has been set.
func (o *ContainerContainerRepositoryResponse) HasManifestSigningService() bool {
	if o != nil && o.ManifestSigningService.IsSet() {
		return true
	}

	return false
}

// SetManifestSigningService gets a reference to the given NullableString and assigns it to the ManifestSigningService field.
func (o *ContainerContainerRepositoryResponse) SetManifestSigningService(v string) {
	o.ManifestSigningService.Set(&v)
}
// SetManifestSigningServiceNil sets the value for ManifestSigningService to be an explicit nil
func (o *ContainerContainerRepositoryResponse) SetManifestSigningServiceNil() {
	o.ManifestSigningService.Set(nil)
}

// UnsetManifestSigningService ensures that no value is present for ManifestSigningService, not even an explicit nil
func (o *ContainerContainerRepositoryResponse) UnsetManifestSigningService() {
	o.ManifestSigningService.Unset()
}

func (o ContainerContainerRepositoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerContainerRepositoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: pulp_href is readOnly
	// skip: pulp_created is readOnly
	// skip: versions_href is readOnly
	if !IsNil(o.PulpLabels) {
		toSerialize["pulp_labels"] = o.PulpLabels
	}
	// skip: latest_version_href is readOnly
	toSerialize["name"] = o.Name
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.RetainRepoVersions.IsSet() {
		toSerialize["retain_repo_versions"] = o.RetainRepoVersions.Get()
	}
	if o.Remote.IsSet() {
		toSerialize["remote"] = o.Remote.Get()
	}
	if o.ManifestSigningService.IsSet() {
		toSerialize["manifest_signing_service"] = o.ManifestSigningService.Get()
	}
	return toSerialize, nil
}

type NullableContainerContainerRepositoryResponse struct {
	value *ContainerContainerRepositoryResponse
	isSet bool
}

func (v NullableContainerContainerRepositoryResponse) Get() *ContainerContainerRepositoryResponse {
	return v.value
}

func (v *NullableContainerContainerRepositoryResponse) Set(val *ContainerContainerRepositoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerContainerRepositoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerContainerRepositoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerContainerRepositoryResponse(val *ContainerContainerRepositoryResponse) *NullableContainerContainerRepositoryResponse {
	return &NullableContainerContainerRepositoryResponse{value: val, isSet: true}
}

func (v NullableContainerContainerRepositoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerContainerRepositoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


