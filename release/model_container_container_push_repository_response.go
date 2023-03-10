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

// checks if the ContainerContainerPushRepositoryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerContainerPushRepositoryResponse{}

// ContainerContainerPushRepositoryResponse Serializer for Container Push Repositories.
type ContainerContainerPushRepositoryResponse struct {
	// An optional description.
	Description NullableString `json:"description,omitempty"`
	VersionsHref *string `json:"versions_href,omitempty"`
	// A unique name for this repository.
	Name string `json:"name"`
	PulpLabels *map[string]string `json:"pulp_labels,omitempty"`
	LatestVersionHref *string `json:"latest_version_href,omitempty"`
	// Retain X versions of the repository. Default is null which retains all versions. This is provided as a tech preview in Pulp 3 and may change in the future.
	RetainRepoVersions NullableInt32 `json:"retain_repo_versions,omitempty"`
	// A reference to an associated signing service.
	ManifestSigningService NullableString `json:"manifest_signing_service,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	PulpHref *string `json:"pulp_href,omitempty"`
}

// NewContainerContainerPushRepositoryResponse instantiates a new ContainerContainerPushRepositoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerContainerPushRepositoryResponse(name string) *ContainerContainerPushRepositoryResponse {
	this := ContainerContainerPushRepositoryResponse{}
	this.Name = name
	return &this
}

// NewContainerContainerPushRepositoryResponseWithDefaults instantiates a new ContainerContainerPushRepositoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerContainerPushRepositoryResponseWithDefaults() *ContainerContainerPushRepositoryResponse {
	this := ContainerContainerPushRepositoryResponse{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerContainerPushRepositoryResponse) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerContainerPushRepositoryResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ContainerContainerPushRepositoryResponse) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ContainerContainerPushRepositoryResponse) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ContainerContainerPushRepositoryResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ContainerContainerPushRepositoryResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetVersionsHref returns the VersionsHref field value if set, zero value otherwise.
func (o *ContainerContainerPushRepositoryResponse) GetVersionsHref() string {
	if o == nil || IsNil(o.VersionsHref) {
		var ret string
		return ret
	}
	return *o.VersionsHref
}

// GetVersionsHrefOk returns a tuple with the VersionsHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerPushRepositoryResponse) GetVersionsHrefOk() (*string, bool) {
	if o == nil || IsNil(o.VersionsHref) {
		return nil, false
	}
	return o.VersionsHref, true
}

// HasVersionsHref returns a boolean if a field has been set.
func (o *ContainerContainerPushRepositoryResponse) HasVersionsHref() bool {
	if o != nil && !IsNil(o.VersionsHref) {
		return true
	}

	return false
}

// SetVersionsHref gets a reference to the given string and assigns it to the VersionsHref field.
func (o *ContainerContainerPushRepositoryResponse) SetVersionsHref(v string) {
	o.VersionsHref = &v
}

// GetName returns the Name field value
func (o *ContainerContainerPushRepositoryResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContainerContainerPushRepositoryResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContainerContainerPushRepositoryResponse) SetName(v string) {
	o.Name = v
}

// GetPulpLabels returns the PulpLabels field value if set, zero value otherwise.
func (o *ContainerContainerPushRepositoryResponse) GetPulpLabels() map[string]string {
	if o == nil || IsNil(o.PulpLabels) {
		var ret map[string]string
		return ret
	}
	return *o.PulpLabels
}

// GetPulpLabelsOk returns a tuple with the PulpLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerPushRepositoryResponse) GetPulpLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PulpLabels) {
		return nil, false
	}
	return o.PulpLabels, true
}

// HasPulpLabels returns a boolean if a field has been set.
func (o *ContainerContainerPushRepositoryResponse) HasPulpLabels() bool {
	if o != nil && !IsNil(o.PulpLabels) {
		return true
	}

	return false
}

// SetPulpLabels gets a reference to the given map[string]string and assigns it to the PulpLabels field.
func (o *ContainerContainerPushRepositoryResponse) SetPulpLabels(v map[string]string) {
	o.PulpLabels = &v
}

// GetLatestVersionHref returns the LatestVersionHref field value if set, zero value otherwise.
func (o *ContainerContainerPushRepositoryResponse) GetLatestVersionHref() string {
	if o == nil || IsNil(o.LatestVersionHref) {
		var ret string
		return ret
	}
	return *o.LatestVersionHref
}

// GetLatestVersionHrefOk returns a tuple with the LatestVersionHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerPushRepositoryResponse) GetLatestVersionHrefOk() (*string, bool) {
	if o == nil || IsNil(o.LatestVersionHref) {
		return nil, false
	}
	return o.LatestVersionHref, true
}

// HasLatestVersionHref returns a boolean if a field has been set.
func (o *ContainerContainerPushRepositoryResponse) HasLatestVersionHref() bool {
	if o != nil && !IsNil(o.LatestVersionHref) {
		return true
	}

	return false
}

// SetLatestVersionHref gets a reference to the given string and assigns it to the LatestVersionHref field.
func (o *ContainerContainerPushRepositoryResponse) SetLatestVersionHref(v string) {
	o.LatestVersionHref = &v
}

// GetRetainRepoVersions returns the RetainRepoVersions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerContainerPushRepositoryResponse) GetRetainRepoVersions() int32 {
	if o == nil || IsNil(o.RetainRepoVersions.Get()) {
		var ret int32
		return ret
	}
	return *o.RetainRepoVersions.Get()
}

// GetRetainRepoVersionsOk returns a tuple with the RetainRepoVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerContainerPushRepositoryResponse) GetRetainRepoVersionsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RetainRepoVersions.Get(), o.RetainRepoVersions.IsSet()
}

// HasRetainRepoVersions returns a boolean if a field has been set.
func (o *ContainerContainerPushRepositoryResponse) HasRetainRepoVersions() bool {
	if o != nil && o.RetainRepoVersions.IsSet() {
		return true
	}

	return false
}

// SetRetainRepoVersions gets a reference to the given NullableInt32 and assigns it to the RetainRepoVersions field.
func (o *ContainerContainerPushRepositoryResponse) SetRetainRepoVersions(v int32) {
	o.RetainRepoVersions.Set(&v)
}
// SetRetainRepoVersionsNil sets the value for RetainRepoVersions to be an explicit nil
func (o *ContainerContainerPushRepositoryResponse) SetRetainRepoVersionsNil() {
	o.RetainRepoVersions.Set(nil)
}

// UnsetRetainRepoVersions ensures that no value is present for RetainRepoVersions, not even an explicit nil
func (o *ContainerContainerPushRepositoryResponse) UnsetRetainRepoVersions() {
	o.RetainRepoVersions.Unset()
}

// GetManifestSigningService returns the ManifestSigningService field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerContainerPushRepositoryResponse) GetManifestSigningService() string {
	if o == nil || IsNil(o.ManifestSigningService.Get()) {
		var ret string
		return ret
	}
	return *o.ManifestSigningService.Get()
}

// GetManifestSigningServiceOk returns a tuple with the ManifestSigningService field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerContainerPushRepositoryResponse) GetManifestSigningServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ManifestSigningService.Get(), o.ManifestSigningService.IsSet()
}

// HasManifestSigningService returns a boolean if a field has been set.
func (o *ContainerContainerPushRepositoryResponse) HasManifestSigningService() bool {
	if o != nil && o.ManifestSigningService.IsSet() {
		return true
	}

	return false
}

// SetManifestSigningService gets a reference to the given NullableString and assigns it to the ManifestSigningService field.
func (o *ContainerContainerPushRepositoryResponse) SetManifestSigningService(v string) {
	o.ManifestSigningService.Set(&v)
}
// SetManifestSigningServiceNil sets the value for ManifestSigningService to be an explicit nil
func (o *ContainerContainerPushRepositoryResponse) SetManifestSigningServiceNil() {
	o.ManifestSigningService.Set(nil)
}

// UnsetManifestSigningService ensures that no value is present for ManifestSigningService, not even an explicit nil
func (o *ContainerContainerPushRepositoryResponse) UnsetManifestSigningService() {
	o.ManifestSigningService.Unset()
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *ContainerContainerPushRepositoryResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerPushRepositoryResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *ContainerContainerPushRepositoryResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *ContainerContainerPushRepositoryResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *ContainerContainerPushRepositoryResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerContainerPushRepositoryResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *ContainerContainerPushRepositoryResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *ContainerContainerPushRepositoryResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

func (o ContainerContainerPushRepositoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerContainerPushRepositoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	// skip: versions_href is readOnly
	toSerialize["name"] = o.Name
	if !IsNil(o.PulpLabels) {
		toSerialize["pulp_labels"] = o.PulpLabels
	}
	// skip: latest_version_href is readOnly
	if o.RetainRepoVersions.IsSet() {
		toSerialize["retain_repo_versions"] = o.RetainRepoVersions.Get()
	}
	if o.ManifestSigningService.IsSet() {
		toSerialize["manifest_signing_service"] = o.ManifestSigningService.Get()
	}
	// skip: pulp_created is readOnly
	// skip: pulp_href is readOnly
	return toSerialize, nil
}

type NullableContainerContainerPushRepositoryResponse struct {
	value *ContainerContainerPushRepositoryResponse
	isSet bool
}

func (v NullableContainerContainerPushRepositoryResponse) Get() *ContainerContainerPushRepositoryResponse {
	return v.value
}

func (v *NullableContainerContainerPushRepositoryResponse) Set(val *ContainerContainerPushRepositoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerContainerPushRepositoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerContainerPushRepositoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerContainerPushRepositoryResponse(val *ContainerContainerPushRepositoryResponse) *NullableContainerContainerPushRepositoryResponse {
	return &NullableContainerContainerPushRepositoryResponse{value: val, isSet: true}
}

func (v NullableContainerContainerPushRepositoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerContainerPushRepositoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


